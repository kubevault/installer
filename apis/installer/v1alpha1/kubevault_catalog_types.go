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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindKubevaultCatalog = "KubevaultCatalog"
	ResourceKubevaultCatalog     = "kubevaultcatalog"
	ResourceKubevaultCatalogs    = "kubevaultcatalogs"
)

// KubevaultCatalog defines the schama for Vault Catalog Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubevaultcatalogs,singular=kubevaultcatalog,categories={kubevault,appscode}
type KubevaultCatalog struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubevaultCatalogSpec `json:"spec,omitempty"`
}

// KubevaultCatalogSpec is the schema for Vault Catalog values file
type KubevaultCatalogSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string      `json:"fullnameOverride"`
	RegistryFQDN     string      `json:"registryFQDN"`
	Image            RegistryRef `json:"image"`
	SkipDeprecated   bool        `json:"skipDeprecated"`
}

type RegistryRef struct {
	Registry                 string `json:"registry"`
	OverrideOfficialRegistry bool   `json:"overrideOfficialRegistry"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubevaultCatalogList is a list of KubevaultCatalogs
type KubevaultCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubevaultCatalog CRD objects
	Items []KubevaultCatalog `json:"items,omitempty"`
}
