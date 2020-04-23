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
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindCSIVault = "CSIVault"
	ResourceCSIVault     = "csivault"
	ResourceCSIVaults    = "csivaults"
)

// CSIVault defines the schama for Vault CSI Driver Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=csivaults,singular=csivault,categories={kubevault,appscode}
type CSIVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              CSIVaultSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type DriverComponent struct {
	Name       string `json:"name" protobuf:"bytes,1,opt,name=name"`
	ImageRef   `json:",inline" protobuf:"bytes,2,opt,name=imageRef"`
	PullPolicy string `json:"pullPolicy" protobuf:"bytes,3,opt,name=pullPolicy"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources" protobuf:"bytes,4,opt,name=resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext" protobuf:"bytes,5,opt,name=securityContext"`
}

type ComponentList struct {
	Controller string `json:"controller" protobuf:"bytes,1,opt,name=controller"`
	Node       string `json:"node" protobuf:"bytes,2,opt,name=node"`
}

// CSIVaultSpec is the schema for CSI Vault Driver values file
type CSIVaultSpec struct {
	NameOverride     string          `json:"nameOverride" protobuf:"bytes,1,opt,name=nameOverride"`
	FullnameOverride string          `json:"fullnameOverride" protobuf:"bytes,2,opt,name=fullnameOverride"`
	Attacher         DriverComponent `json:"attacher" protobuf:"bytes,3,opt,name=attacher"`
	Plugin           DriverComponent `json:"plugin" protobuf:"bytes,4,opt,name=plugin"`
	Provisioner      DriverComponent `json:"provisioner" protobuf:"bytes,5,opt,name=provisioner"`
	NodeRegistrar    DriverComponent `json:"nodeRegistrar" protobuf:"bytes,6,opt,name=nodeRegistrar"`
	Components       ComponentList   `json:"components" protobuf:"bytes,7,opt,name=components"`
	DriverName       string          `json:"driverName" protobuf:"bytes,8,opt,name=driverName"`
	PluginAddress    string          `json:"pluginAddress" protobuf:"bytes,9,opt,name=pluginAddress"`
	PluginDir        string          `json:"pluginDir" protobuf:"bytes,10,opt,name=pluginDir"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets" protobuf:"bytes,11,rep,name=imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon" protobuf:"varint,12,opt,name=criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel" protobuf:"varint,13,opt,name=logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations" protobuf:"bytes,14,rep,name=annotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector" protobuf:"bytes,15,rep,name=nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations" protobuf:"bytes,16,rep,name=tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity" protobuf:"bytes,17,opt,name=affinity"`
	RBAC     CSIDriverRBAC  `json:"rbac" protobuf:"bytes,18,opt,name=rbac"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext" protobuf:"bytes,19,opt,name=podSecurityContext"`
	Apiserver          CSIDriverAPIServer       `json:"apiserver" protobuf:"bytes,20,opt,name=apiserver"`
	//+optional
	EnableAnalytics bool                `json:"enableAnalytics" protobuf:"varint,21,opt,name=enableAnalytics"`
	Monitoring      CSIDriverMonitoring `json:"monitoring" protobuf:"bytes,22,opt,name=monitoring"`
}

type CSIDriverRBAC struct {
	Create bool `json:"create" protobuf:"varint,1,opt,name=create"`
}

type CSIDriverAPIServer struct {
	UseKubeapiserverFqdnForAks bool            `json:"useKubeapiserverFqdnForAks" protobuf:"varint,1,opt,name=useKubeapiserverFqdnForAks"`
	Healthcheck                HealthcheckSpec `json:"healthcheck" protobuf:"bytes,2,opt,name=healthcheck"`
}

type CSIDriverMonitoring struct {
	Agent          string                `json:"agent" protobuf:"bytes,1,opt,name=agent"`
	Node           bool                  `json:"node" protobuf:"varint,2,opt,name=node"`
	Controller     bool                  `json:"controller" protobuf:"varint,3,opt,name=controller"`
	Prometheus     *PrometheusSpec       `json:"prometheus" protobuf:"bytes,4,opt,name=prometheus"`
	ServiceMonitor *ServiceMonitorLabels `json:"serviceMonitor" protobuf:"bytes,5,opt,name=serviceMonitor"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CSIVaultList is a list of CSIVaults
type CSIVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of CSIVault CRD objects
	Items []CSIVault `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
