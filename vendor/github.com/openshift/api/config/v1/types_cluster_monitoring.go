/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterMonitoring is the Custom Resource object which holds the current status of Cluster Monitoring Operator. CMO is a central component of the monitoring stack.
//
// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:internal
// +openshift:api-approved.openshift.io=https://github.com/openshift/api/pull/1929
// +openshift:file-pattern=cvoRunLevel=0000_10,operatorName=config-operator,operatorOrdering=01
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=clustermonitoring,scope=Cluster
// +kubebuilder:subresource:status
// +kubebuilder:metadata:annotations="description=Cluster Monitoring Operators configuration API"
// +openshift:enable:FeatureGate=ClusterMonitoringConfig
// ClusterMonitoring is the Schema for the Cluster Monitoring Operators API
type ClusterMonitoring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user configuration for the Cluster Monitoring Operator
	// +kubebuilder:validation:Required
	Spec ClusterMonitoringSpec `json:"spec,omitempty"`
	// status holds observed values from the cluster. They may not be overridden.
	// +optional
	Status ClusterMonitoringStatus `json:"status,omitempty"`
}

// MonitoringOperatorStatus defines the observed state of MonitoringOperator
type ClusterMonitoringStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:internal
type ClusterMonitoringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterMonitoring `json:"items"`
}

// ClusterMonitoringSpec defines the desired state of Cluster Monitoring Operator
type ClusterMonitoringSpec struct {
	// userDefined set the deployment mode for user-defined monitoring in addition to the default platform monitoring.
	UserDefined UserDefinedMonitoring `json:"userDefined,omitempty"`
}

// UserDefinedMonitoring config for user-defined projects.
type UserDefinedMonitoring struct {
	// mode defines the different configurations of UserDefinedMonitoring
	// +kubebuilder:validation:Enum:="Disabled";"NamespaceIsolation"
	Mode UserDefinedMode `json:"mode,omitempty"`
}

type UserDefinedMode string

const (
	// UserDefinedDisabled disables monitoring for user-defined projects. This restrics the default monitoring stack, installed in the openshift-monitoring project, to monitor only platform namespaces, which prevents any custom monitoring configurations or resources from being applied to user-defined namespaces.
	UserDefinedDisabled UserDefinedMode = "Disabled"
	// UserDefinedNamespaceIsolation enables monitoring for user-defined projects with namespace-scoped tenancy. This ensures that metrics, alerts, and monitoring data are isolated at the namespace level.
	UserDefinedNamespaceIsolation UserDefinedMode = "NamespaceIsolation"
)
