/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"encoding/json"
	goflag "flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"kubevault.dev/installer/hack/fmt/templates"

	"github.com/Masterminds/semver/v3"
	"github.com/Masterminds/sprig/v3"
	flag "github.com/spf13/pflag"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	shell "gomodules.xyz/go-sh"
	"gomodules.xyz/semvers"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kmodules.xyz/client-go/tools/parser"
	"sigs.k8s.io/yaml"
)

const registryKubeVault = "kubevault"

type AppVersion struct {
	Group   string
	Kind    string
	Version string
	Distro  string
}

type FullVersion struct {
	Version     string
	CatalogName string
}

// FullverCollection is a collection of Version instances and implements the sort
// interface. See the sort package for more details.
// https://golang.org/pkg/sort/
type FullverCollection []FullVersion

// Len returns the length of a collection. The number of Version instances
// on the slice.
func (c FullverCollection) Len() int {
	return len(c)
}

// Less is needed for the sort interface to compare two Version objects on the
// slice. If checks if one is less than the other.
func (c FullverCollection) Less(i, j int) bool {
	return CompareFullVersions(c[i], c[j])
}

// Swap is needed for the sort interface to replace the Version objects
// at two different positions in the slice.
func (c FullverCollection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func CompareFullVersions(vi FullVersion, vj FullVersion) bool {
	vvi, _ := semver.NewVersion(vi.Version)
	vvj, _ := semver.NewVersion(vj.Version)
	if result := vvi.Compare(vvj); result != 0 {
		return result < 0
	}
	return Compare(vi.CatalogName, vj.CatalogName)
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	/*
		Key/Value map used to update pg-coordinator and replication mode detector image
		// MySQL, MongoDB
		--update-spec=spec.replicationModeDetector.image=_new_image
		//Postgres
		--update-spec=spec.coordinator.image=_new_image
	*/
	specUpdates := map[string]string{}

	flag.StringVar(&dir, "dir", dir, "Path to directory where the kubevault/installer project resides (default is set o current directory)")
	flag.StringToStringVar(&specUpdates, "update-spec", specUpdates, "Key/Value map used to update unsealer image")

	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	resources, err := ListResources(filepath.Join(dir, "catalog", "raw"))
	if err != nil {
		panic(err)
	}

	appStore := map[AppVersion][]*unstructured.Unstructured{}

	// active versions
	activeAppVersions := map[string][]FullVersion{}

	for _, obj := range resources {
		// remove labels
		obj.SetNamespace("")
		obj.SetLabels(nil)
		obj.SetAnnotations(nil)

		for jp, val := range specUpdates {
			if _, ok, _ := unstructured.NestedFieldNoCopy(obj.Object, strings.Split(jp, ".")...); ok {
				err = unstructured.SetNestedField(obj.Object, val, strings.Split(jp, ".")...)
				if err != nil {
					panic(fmt.Sprintf("failed to set %s to %s in group=%s,kind=%s,name=%s", jp, val, obj.GetAPIVersion(), obj.GetKind(), obj.GetName()))
				}
			}
		}

		gv, err := schema.ParseGroupVersion(obj.GetAPIVersion())
		if err != nil {
			panic(err)
		}
		if gv.Group == "catalog.kubevault.com" {
			appKind := strings.TrimSuffix(obj.GetKind(), "Version")
			deprecated, _, _ := unstructured.NestedBool(obj.Object, "spec", "deprecated")

			distro, _, _ := unstructured.NestedString(obj.Object, "spec", "distribution")

			appVersion, _, err := unstructured.NestedString(obj.Object, "spec", "version")
			if err != nil {
				panic(err)
			}
			appverKey := AppVersion{
				Group:   gv.Group,
				Kind:    obj.GetKind(),
				Version: appVersion,
				Distro:  distro,
			}
			appStore[appverKey] = append(appStore[appverKey], obj)

			if !deprecated {
				activeAppVersions[appKind] = append(activeAppVersions[appKind], FullVersion{
					Version:     appVersion,
					CatalogName: obj.GetName(),
				})
			}
		}
	}

	{
		activeVers := map[string][]string{}

		for k, v := range activeAppVersions {
			sort.Sort(sort.Reverse(FullverCollection(v)))
			activeAppVersions[k] = v

			for _, e := range v {
				activeVers[k] = append(activeVers[k], e.CatalogName)
			}
		}

		data, err := json.MarshalIndent(activeVers, "", "  ")
		if err != nil {
			panic(err)
		}

		filename := filepath.Join(dir, "catalog", "active_versions.json")
		err = os.MkdirAll(filepath.Dir(filename), 0o755)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(filename, data, 0o644)
		if err != nil {
			panic(err)
		}
	}

	for k, v := range appStore {
		sort.Slice(v, func(i, j int) bool {
			di, _, _ := unstructured.NestedBool(v[i].Object, "spec", "deprecated")
			dj, _, _ := unstructured.NestedBool(v[j].Object, "spec", "deprecated")

			if (di && dj) || (!di && !dj) {
				// company version
				return Compare(v[i].GetName(), v[j].GetName())
			}
			return dj // or !di
		})
		appStore[k] = v

		var buf bytes.Buffer
		for i, obj := range v {
			if i > 0 {
				buf.WriteString("\n---\n")
			}
			data, err := yaml.Marshal(obj)
			if err != nil {
				panic(err)
			}
			buf.Write(data)
		}

		appKind := strings.TrimSuffix(k.Kind, "Version")

		var filenameparts []string
		if allDeprecated(v) {
			filenameparts = append(filenameparts, "deprecated")
		}
		filenameparts = append(filenameparts, strings.ToLower(appKind), k.Version)
		if k.Distro != "" {
			filenameparts = append(filenameparts, strings.ToLower(k.Distro))
		}
		filename := filepath.Join(dir, "catalog", "new_raw", strings.ToLower(appKind), fmt.Sprintf("%s.yaml", strings.Join(filenameparts, "-")))
		err = os.MkdirAll(filepath.Dir(filename), 0o755)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(filename, buf.Bytes(), 0o644)
		if err != nil {
			panic(err)
		}
	}

	{
		// move new_raw to raw
		err = os.RemoveAll(filepath.Join(dir, "catalog", "raw"))
		if err != nil {
			panic(err)
		}
		err = os.Rename(filepath.Join(dir, "catalog", "new_raw"), filepath.Join(dir, "catalog", "raw"))
		if err != nil {
			panic(err)
		}
	}

	// GENERATE CHART
	{
		for k, v := range appStore {
			appKind := strings.TrimSuffix(k.Kind, "Version")
			var buf bytes.Buffer

			for i, obj := range v {
				obj := obj.DeepCopy()

				spec, _, err := unstructured.NestedMap(obj.Object, "spec")
				if err != nil {
					panic(err)
				}
				for prop := range spec {
					templatizeRegistry := func(field string) {
						img, ok, _ := unstructured.NestedString(obj.Object, "spec", prop, field)
						if ok {
							reg, bin, tag := ParseImage(img)
							var newimg string
							switch {
							case tag == "" && reg == "":
								newimg = fmt.Sprintf(`{{ include "official.registry" (set (.Values | deepCopy) "officialRegistry" (list %q)) }}`, bin)
							case tag == "" && reg == registryKubeVault:
								newimg = fmt.Sprintf(`{{ include "catalog.registry" . }}/%s`, bin)
							case tag == "" && reg != registryKubeVault:
								newimg = fmt.Sprintf(`{{ include "official.registry" (set (.Values | deepCopy) "officialRegistry" (list %q %q)) }}`, reg, bin)
							case tag != "" && reg == "":
								newimg = fmt.Sprintf(`{{ include "official.registry" (set (.Values | deepCopy) "officialRegistry" (list %q)) }}:%s`, bin, tag)
							case reg == registryKubeVault:
								newimg = fmt.Sprintf(`{{ include "catalog.registry" . }}/%s:%s`, bin, tag)
							default:
								newimg = fmt.Sprintf(`{{ include "official.registry" (set (.Values | deepCopy) "officialRegistry" (list %q %q)) }}:%s`, reg, bin, tag)
							}
							err = unstructured.SetNestedField(obj.Object, newimg, "spec", prop, field)
							if err != nil {
								panic(err)
							}
						}
					}
					templatizeRegistry("image")
					templatizeRegistry("yqImage")
				}

				if i > 0 {
					buf.WriteString("\n---\n")
				}

				data := map[string]interface{}{
					"key":    strings.ToLower(appKind),
					"object": obj.Object,
				}
				funcMap := sprig.TxtFuncMap()
				funcMap["toYaml"] = toYAML
				funcMap["toJson"] = toJSON
				tpl := template.Must(template.New("").Funcs(funcMap).Parse(templates.AppVersion))
				err = tpl.Execute(&buf, &data)
				if err != nil {
					panic(err)
				}
			}

			var filenameparts []string
			if allDeprecated(v) {
				filenameparts = append(filenameparts, "deprecated")
			}
			filenameparts = append(filenameparts, strings.ToLower(appKind), k.Version)
			if k.Distro != "" {
				filenameparts = append(filenameparts, strings.ToLower(k.Distro))
			}
			filename := filepath.Join(dir, "charts", "kubevault-catalog", "new_templates", strings.ToLower(appKind), fmt.Sprintf("%s.yaml", strings.Join(filenameparts, "-")))
			err = os.MkdirAll(filepath.Dir(filename), 0o755)
			if err != nil {
				panic(err)
			}

			err = ioutil.WriteFile(filename, buf.Bytes(), 0o644)
			if err != nil {
				panic(err)
			}
		}

		{
			// move new_templates to templates
			err = os.Rename(filepath.Join(dir, "charts", "kubevault-catalog", "templates", "_helpers.tpl"), filepath.Join(dir, "charts", "kubevault-catalog", "new_templates", "_helpers.tpl"))
			if err != nil {
				panic(err)
			}
			err = os.RemoveAll(filepath.Join(dir, "charts", "kubevault-catalog", "templates"))
			if err != nil {
				panic(err)
			}
			err = os.Rename(filepath.Join(dir, "charts", "kubevault-catalog", "new_templates"), filepath.Join(dir, "charts", "kubevault-catalog", "templates"))
			if err != nil {
				panic(err)
			}
		}
	}

	{
		// Verify
		type ObjectKey struct {
			APIVersion string
			Kind       string
			Name       string
		}
		type DiffData struct {
			A *unstructured.Unstructured
			B *unstructured.Unstructured
		}

		dm := map[ObjectKey]*DiffData{}
		for _, obj := range resources {
			dm[ObjectKey{
				APIVersion: obj.GetAPIVersion(),
				Kind:       obj.GetKind(),
				Name:       obj.GetName(),
			}] = &DiffData{
				A: obj,
			}
		}

		failed := false
		differ := diff.New()

		sh := shell.NewSession()
		sh.SetDir(dir)
		sh.ShowCMD = true

		out, err := sh.Command("helm", "template", "charts/kubevault-catalog", "--set", "skipDeprecated=false").Output()
		if err != nil {
			panic(err)
		}

		helmout, err := parser.ListResources(out)
		if err != nil {
			panic(err)
		}

		for _, ri := range helmout {
			ri.Object.SetNamespace("")
			ri.Object.SetLabels(nil)
			ri.Object.SetAnnotations(nil)

			key := ObjectKey{
				APIVersion: ri.Object.GetAPIVersion(),
				Kind:       ri.Object.GetKind(),
				Name:       ri.Object.GetName(),
			}
			if _, ok := dm[key]; !ok {
				failed = true
				_, _ = fmt.Fprintf(os.Stderr, "missing object is raw apiVersion=%s kind=%s name=%s", key.APIVersion, key.Kind, key.Name)
			} else {
				dm[key].B = ri.Object
			}
		}

		for key, data := range dm {
			if data.B == nil {
				failed = true
				_, _ = fmt.Fprintf(os.Stderr, "missing object is catalog chart apiVersion=%s kind=%s name=%s", key.APIVersion, key.Kind, key.Name)
				continue
			}

			a, err := json.Marshal(data.A)
			if err != nil {
				panic(err)
			}
			b, err := json.Marshal(data.B)
			if err != nil {
				panic(err)
			}

			// Then, Check them
			d, err := differ.Compare(a, b)
			if err != nil {
				fmt.Printf("Failed to unmarshal file: %s\n", err.Error())
				os.Exit(3)
			}

			if d.Modified() {
				failed = true

				config := formatter.AsciiFormatterConfig{
					ShowArrayIndex: true,
					Coloring:       true,
				}

				f := formatter.NewAsciiFormatter(data.A.Object, config)
				result, err := f.Format(d)
				if err != nil {
					panic(err)
				}
				_, _ = fmt.Fprintf(os.Stderr, "mismatched apiVersion=%s kind=%s name=%s \ndiff=%s", key.APIVersion, key.Kind, key.Name, result)
				continue
			}
		}

		if failed {
			os.Exit(1)
		}
	}
}

func ParseImage(s string) (reg, bin, tag string) {
	idx := strings.IndexRune(s, ':')
	if idx != -1 {
		tag = s[idx+1:]
		s = s[:idx]
	}
	idx = strings.IndexRune(s, '/')
	if idx == -1 {
		bin = s
	} else {
		reg, bin = s[:idx], s[idx+1:]
	}
	return
}

func ListResources(dir string) ([]*unstructured.Unstructured, error) {
	var resources []*unstructured.Unstructured

	err := parser.ProcessPath(dir, func(ri parser.ResourceInfo) error {
		if ri.Object.GetNamespace() == "" {
			ri.Object.SetNamespace(core.NamespaceDefault)
		}
		resources = append(resources, ri.Object)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func allDeprecated(objs []*unstructured.Unstructured) bool {
	for _, obj := range objs {
		d, _, _ := unstructured.NestedBool(obj.Object, "spec", "deprecated")
		if !d {
			return false
		}
	}
	return true
}

// toYAML takes an interface, marshals it to yaml, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
func toYAML(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return strings.TrimSuffix(string(data), "\n")
}

// toJSON takes an interface, marshals it to json, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
func toJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return string(data)
}

func Compare(i, j string) bool {
	vi, ei := semver.NewVersion(i)
	vj, ej := semver.NewVersion(j)
	if ei == nil && ej == nil {
		return semvers.CompareVersions(vi, vj)
	}

	idx_i := strings.IndexRune(i, '-')
	var distro_i, ver_i string
	if idx_i != -1 {
		distro_i, ver_i = i[:idx_i], i[idx_i+1:]
	}
	idx_j := strings.IndexRune(j, '-')
	var distro_j, ver_j string
	if idx_j != -1 {
		distro_j, ver_j = j[:idx_j], j[idx_j+1:]
	}
	if distro_i == distro_j && distro_i != "" {
		return semvers.Compare(ver_i, ver_j)
	}
	return strings.Compare(i, j) < 0
}
