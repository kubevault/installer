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
	ResourceKindKubevaultWebhookServer = "KubevaultWebhookServer"
	ResourceKubevaultWebhookServer     = "kubevaultwebhookserver"
	ResourceKubevaultWebhookServers    = "kubevaultwebhookservers"
)

// KubevaultWebhookServer defines the schama for ui server installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubevaultwebhookservers,singular=kubevaultwebhookserver,categories={kubevault,appscode}
type KubevaultWebhookServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubevaultWebhookServerSpec `json:"spec,omitempty"`
}

// KubevaultWebhookServerSpec is the schema for Identity Server values file
type KubevaultWebhookServerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string    `json:"fullnameOverride"`
	ReplicaCount     int32     `json:"replicaCount"`
	RegistryFQDN     string    `json:"registryFQDN"`
	Server           Container `json:"server"`
	ImagePullPolicy  string    `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	// +optional
	CriticalAddon bool `json:"criticalAddon"`
	// +optional
	PriorityClassName string `json:"priorityClassName"`
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
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	Apiserver          WebhookAPIServerSpec     `json:"apiserver"`
	Monitoring         UIServerMonitoring       `json:"monitoring"`
}

type WebhookAPIServerSpec struct {
	GroupPriorityMinimum       int32                   `json:"groupPriorityMinimum"`
	VersionPriority            int32                   `json:"versionPriority"`
	EnableMutatingWebhook      bool                    `json:"enableMutatingWebhook"`
	EnableValidatingWebhook    bool                    `json:"enableValidatingWebhook"`
	CA                         string                  `json:"ca"`
	UseKubeapiserverFqdnForAks bool                    `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                UIServerHealthcheckSpec `json:"healthcheck"`
	Port                       int32                   `json:"port"`
	ServingCerts               ServingCerts            `json:"servingCerts"`
	Webhook                    WebhookSpec             `json:"webhook"`
}

type WebhookSpec struct {
	FailurePolicy string `json:"failurePolicy"`
}

type UIServerHealthcheckSpec struct {
	// +optional
	Enabled bool `json:"enabled"`
}

type UIServerMonitoring struct {
	Agent          MonitoringAgent       `json:"agent"`
	ServiceMonitor *ServiceMonitorLabels `json:"serviceMonitor"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubevaultWebhookServerList is a list of KubevaultWebhookServers
type KubevaultWebhookServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubevaultWebhookServer CRD objects
	Items []KubevaultWebhookServer `json:"items,omitempty"`
}
