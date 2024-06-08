/*
Copyright 2017 The Kubernetes Authors.

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

// RolloutSample is a specification for a RolloutSample resource
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=rollout
// +kubebuilder:printcolumn:name="Replicas",type=integer,JSONPath=`.spec.replicas`

type RolloutSample struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec RolloutSampleSpec `json:"spec"`
	// +optional
	Status SampleStatus `json:"status"`
}

// RolloutSampleSpec is the spec for a Foo resource
type RolloutSampleSpec struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=30
	Name string `json:"name"`

	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=100
	Image string `json:"image"`

	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Maximum=30
	Replicas *int32 `json:"replicas"`
}

// SampleStatus is the status for a RolloutSample resource
type SampleStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// RolloutSampleList is a list of RolloutSample resources
// +kubebuilder:object:root=true
type RolloutSampleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RolloutSample `json:"items"`
}