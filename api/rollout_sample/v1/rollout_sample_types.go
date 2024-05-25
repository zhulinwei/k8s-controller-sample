package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These const variables are used in our custom controller.
const (
	GroupName string = "zlw.com"
	Version   string = "v1"
)

// RolloutSample is a specification for a RolloutSample resource
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=rollout
// +kubebuilder:printcolumn:name="Replicas",type=integer,JSONPath=`.spec.replicas`
type RolloutSample struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RolloutSampleSpec   `json:"spec"`
	// +optional
	Status RolloutSampleStatus `json:"status"`
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
	// +kubebuilder:validation:Maximum=100
	Replicas *int32 `json:"replicas"`
}

// RolloutSampleStatus is the status for a RolloutSample resource
type RolloutSampleStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// RolloutSampleList is a list of RolloutSample resources
// +kubebuilder:object:root=true
type RolloutSampleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RolloutSample `json:"items"`
}
