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
	ResourceKindKubevault = "Kubevault"
	ResourceKubevault     = "kubevault"
	ResourceKubevaults    = "kubevaults"
)

// Kubevault defines the schama for KubeVault Operator Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubevaults,singular=kubevault,categories={kubevault,appscode}
type Kubevault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubevaultSpec `json:"spec,omitempty"`
}

type ImageRef struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type Container struct {
	ImageRef `json:",inline"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

// KubevaultSpec is the schema for KubeVault operator values file
type KubevaultSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string    `json:"fullnameOverride"`
	ReplicaCount     int32     `json:"replicaCount"`
	Operator         Container `json:"operator"`
	Cleaner          ImageRef  `json:"cleaner"`
	ImagePullPolicy  string    `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	Apiserver          OperatorWebhookServer    `json:"apiserver"`
	//+optional
	EnableAnalytics bool               `json:"enableAnalytics"`
	Monitoring      OperatorMonitoring `json:"monitoring"`
	//+optional
	ClusterName *string `json:"clusterName"`
}

type ServiceAccountSpec struct {
	Create bool `json:"create"`
	//+optional
	Name *string `json:"name"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}

type OperatorWebhookServer struct {
	GroupPriorityMinimum    int32  `json:"groupPriorityMinimum"`
	VersionPriority         int32  `json:"versionPriority"`
	EnableMutatingWebhook   bool   `json:"enableMutatingWebhook"`
	EnableValidatingWebhook bool   `json:"enableValidatingWebhook"`
	CA                      string `json:"ca"`
	//+optional
	BypassValidatingWebhookXray bool            `json:"bypassValidatingWebhookXray"`
	UseKubeapiserverFqdnForAks  bool            `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                 HealthcheckSpec `json:"healthcheck"`
	ServingCerts                ServingCerts    `json:"servingCerts"`
}

type HealthcheckSpec struct {
	//+optional
	Enabled bool `json:"enabled"`
}

type ServingCerts struct {
	Generate bool `json:"generate"`
	//+optional
	CaCrt string `json:"caCrt"`
	//+optional
	ServerCrt string `json:"serverCrt"`
	//+optional
	ServerKey string `json:"serverKey"`
}

type OperatorMonitoring struct {
	Agent string `json:"agent"`
	//+optional
	Operator       bool                  `json:"operator"`
	Prometheus     *PrometheusSpec       `json:"prometheus"`
	ServiceMonitor *ServiceMonitorLabels `json:"serviceMonitor"`
}

type PrometheusSpec struct {
	//+optional
	Namespace string `json:"namespace"`
}

type ServiceMonitorLabels struct {
	//+optional
	Labels map[string]string `json:"labels"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubevaultList is a list of Kubevaults
type KubevaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Kubevault CRD objects
	Items []Kubevault `json:"items,omitempty"`
}
