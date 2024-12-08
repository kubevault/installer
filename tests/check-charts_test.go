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
	"errors"
	"path/filepath"
	"runtime"
	"testing"

	"kmodules.xyz/image-packer/pkg/lib"
)

func Test_CheckImageArchitectures(t *testing.T) {
	dir, err := rootDir()
	if err != nil {
		t.Error(err)
	}

	if err := lib.CheckImageArchitectures([]string{
		filepath.Join(dir, "catalog", "imagelist.yaml"),
	}, []string{
		"vault:1.10.3",
		"vault:1.11.5",
		"vault:1.12.1",
		"vault:1.13.3",
		"vault:1.2.0",
		"vault:1.2.2",
		"vault:1.2.3",
		"vault:1.5.9",
		"vault:1.6.5",
		"vault:1.7.2",
		"vault:1.7.3",
		"vault:1.8.2",
		"vault:1.9.2",
	}); err != nil {
		t.Errorf("CheckImageArchitectures() error = %v", err)
	}
}

func rootDir() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("failed to locate root dir")
	}

	return filepath.Clean(filepath.Join(filepath.Dir(file), "..")), nil
}
