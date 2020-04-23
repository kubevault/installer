/*
Copyright The KubeVault Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

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
	ResourceKindVaultCatalog = "VaultCatalog"
	ResourceVaultCatalog     = "vaultcatalog"
	ResourceVaultCatalogs    = "vaultcatalogs"
)

// VaultCatalog defines the schama for Vault Catalog Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=vaultcatalogs,singular=vaultcatalog,categories={kubevault,appscode}
type VaultCatalog struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              VaultCatalogSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// VaultCatalogSpec is the schema for Vault Catalog values file
type VaultCatalogSpec struct {
	DockerRegistry string `json:"dockerRegistry" protobuf:"bytes,1,opt,name=dockerRegistry"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VaultCatalogList is a list of VaultCatalogs
type VaultCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of VaultCatalog CRD objects
	Items []VaultCatalog `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
