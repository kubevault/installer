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
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindCsiVault = "CsiVault"
	ResourceCsiVault     = "csivault"
	ResourceCsiVaults    = "csivaults"
)

// CsiVault defines the schama for Vault CSI Driver Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=csivaults,singular=csivault,categories={kubevault,appscode}
type CsiVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CsiVaultSpec `json:"spec,omitempty"`
}

type DriverComponent struct {
	Name       string `json:"name"`
	ImageRef   `json:",inline"`
	PullPolicy string `json:"pullPolicy"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

type ComponentList struct {
	Controller string `json:"controller"`
	Node       string `json:"node"`
}

// CsiVaultSpec is the schema for CSI Vault Driver values file
type CsiVaultSpec struct {
	NameOverride     string          `json:"nameOverride"`
	FullnameOverride string          `json:"fullnameOverride"`
	Attacher         DriverComponent `json:"attacher"`
	Plugin           DriverComponent `json:"plugin"`
	Provisioner      DriverComponent `json:"provisioner"`
	NodeRegistrar    DriverComponent `json:"nodeRegistrar"`
	Components       ComponentList   `json:"components"`
	DriverName       string          `json:"driverName"`
	PluginAddress    string          `json:"pluginAddress"`
	PluginDir        string          `json:"pluginDir"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	RBAC     CSIDriverRBAC  `json:"rbac"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	Apiserver          CSIDriverAPIServer       `json:"apiserver"`
	//+optional
	EnableAnalytics bool       `json:"enableAnalytics"`
	Monitoring      Monitoring `json:"monitoring"`
}

type CSIDriverRBAC struct {
	Create bool `json:"create"`
}

type CSIDriverAPIServer struct {
	UseKubeapiserverFqdnForAks bool            `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                HealthcheckSpec `json:"healthcheck"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CsiVaultList is a list of CsiVaults
type CsiVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CsiVault CRD objects
	Items []CsiVault `json:"items,omitempty"`
}
