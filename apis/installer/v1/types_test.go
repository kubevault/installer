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

package v1_test

import (
	"os"
	"testing"

	v1 "kubevault.dev/installer/apis/installer/v1"

	schemachecker "kmodules.xyz/schema-checker"
)

func TestDefaultValues(t *testing.T) {
	checker := schemachecker.New(
		os.DirFS("../../.."),
		schemachecker.TestCase{Obj: v1.KubevaultCatalogSpec{}},
		schemachecker.TestCase{Obj: v1.KubevaultOperatorSpec{}},
		schemachecker.TestCase{Obj: v1.KubevaultWebhookServerSpec{}},
		schemachecker.TestCase{Obj: v1.KubevaultSpec{}},
	)
	checker.TestAll(t)
}
