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
	ResourceKindKubeVaultOperator = "KubeVaultOperator"
	ResourceKubeVaultOperator     = "kubevaultoperator"
	ResourceKubeVaultOperators    = "kubevaultoperators"
)

// KubeVaultOperator defines the schama for KubeVault Operator Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubevaultoperators,singular=kubevaultoperator,categories={kubevault,appscode}
type KubeVaultOperator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              KubeVaultOperatorSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type ImageRef struct {
	Registry   string `json:"registry" protobuf:"bytes,1,opt,name=registry"`
	Repository string `json:"repository" protobuf:"bytes,2,opt,name=repository"`
	Tag        string `json:"tag" protobuf:"bytes,3,opt,name=tag"`
}

type Container struct {
	ImageRef `json:",inline" protobuf:"bytes,1,opt,name=imageRef"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources" protobuf:"bytes,2,opt,name=resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext" protobuf:"bytes,3,opt,name=securityContext"`
}

// KubeVaultOperatorSpec is the schema for KubeVault operator values file
type KubeVaultOperatorSpec struct {
	//+optional
	NameOverride string `json:"nameOverride" protobuf:"bytes,1,opt,name=nameOverride"`
	//+optional
	FullnameOverride string    `json:"fullnameOverride" protobuf:"bytes,2,opt,name=fullnameOverride"`
	ReplicaCount     int32     `json:"replicaCount" protobuf:"varint,3,opt,name=replicaCount"`
	Operator         Container `json:"operator" protobuf:"bytes,4,opt,name=operator"`
	Cleaner          ImageRef  `json:"cleaner" protobuf:"bytes,5,opt,name=cleaner"`
	ImagePullPolicy  string    `json:"imagePullPolicy" protobuf:"bytes,6,opt,name=imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets" protobuf:"bytes,7,rep,name=imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon" protobuf:"varint,8,opt,name=criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel" protobuf:"varint,9,opt,name=logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations" protobuf:"bytes,10,rep,name=annotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector" protobuf:"bytes,11,rep,name=nodeSelector"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations" protobuf:"bytes,12,rep,name=podAnnotations"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations" protobuf:"bytes,13,rep,name=tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity" protobuf:"bytes,14,opt,name=affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext" protobuf:"bytes,15,opt,name=podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount" protobuf:"bytes,16,opt,name=serviceAccount"`
	Apiserver          OperatorWebhookServer    `json:"apiserver" protobuf:"bytes,17,opt,name=apiserver"`
	//+optional
	EnableAnalytics bool               `json:"enableAnalytics" protobuf:"varint,18,opt,name=enableAnalytics"`
	Monitoring      OperatorMonitoring `json:"monitoring" protobuf:"bytes,19,opt,name=monitoring"`
	//+optional
	ClusterName *string `json:"clusterName" protobuf:"bytes,20,opt,name=clusterName"`
}

type ServiceAccountSpec struct {
	Create bool `json:"create" protobuf:"varint,1,opt,name=create"`
	//+optional
	Name *string `json:"name" protobuf:"bytes,2,opt,name=name"`
	//+optional
	Annotations map[string]string `json:"annotations" protobuf:"bytes,3,rep,name=annotations"`
}

type OperatorWebhookServer struct {
	GroupPriorityMinimum    int32  `json:"groupPriorityMinimum" protobuf:"varint,1,opt,name=groupPriorityMinimum"`
	VersionPriority         int32  `json:"versionPriority" protobuf:"varint,2,opt,name=versionPriority"`
	EnableMutatingWebhook   bool   `json:"enableMutatingWebhook" protobuf:"varint,3,opt,name=enableMutatingWebhook"`
	EnableValidatingWebhook bool   `json:"enableValidatingWebhook" protobuf:"varint,4,opt,name=enableValidatingWebhook"`
	CA                      string `json:"ca" protobuf:"bytes,5,opt,name=ca"`
	//+optional
	BypassValidatingWebhookXray bool            `json:"bypassValidatingWebhookXray" protobuf:"varint,6,opt,name=bypassValidatingWebhookXray"`
	UseKubeapiserverFqdnForAks  bool            `json:"useKubeapiserverFqdnForAks" protobuf:"varint,7,opt,name=useKubeapiserverFqdnForAks"`
	Healthcheck                 HealthcheckSpec `json:"healthcheck" protobuf:"bytes,8,opt,name=healthcheck"`
	ServingCerts                ServingCerts    `json:"servingCerts" protobuf:"bytes,9,opt,name=servingCerts"`
}

type HealthcheckSpec struct {
	//+optional
	Enabled bool `json:"enabled" protobuf:"varint,1,opt,name=enabled"`
}

type ServingCerts struct {
	Generate bool `json:"generate" protobuf:"varint,1,opt,name=generate"`
	//+optional
	CaCrt string `json:"caCrt" protobuf:"bytes,2,opt,name=caCrt"`
	//+optional
	ServerCrt string `json:"serverCrt" protobuf:"bytes,3,opt,name=serverCrt"`
	//+optional
	ServerKey string `json:"serverKey" protobuf:"bytes,4,opt,name=serverKey"`
}

type OperatorMonitoring struct {
	Agent string `json:"agent" protobuf:"bytes,1,opt,name=agent"`
	//+optional
	Operator       bool                  `json:"operator" protobuf:"varint,2,opt,name=operator"`
	Prometheus     *PrometheusSpec       `json:"prometheus" protobuf:"bytes,3,opt,name=prometheus"`
	ServiceMonitor *ServiceMonitorLabels `json:"serviceMonitor" protobuf:"bytes,4,opt,name=serviceMonitor"`
}

type PrometheusSpec struct {
	//+optional
	Namespace string `json:"namespace" protobuf:"bytes,1,opt,name=namespace"`
}

type ServiceMonitorLabels struct {
	//+optional
	Labels map[string]string `json:"labels" protobuf:"bytes,1,rep,name=labels"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeVaultOperatorList is a list of KubeVaultOperators
type KubeVaultOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of KubeVaultOperator CRD objects
	Items []KubeVaultOperator `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
