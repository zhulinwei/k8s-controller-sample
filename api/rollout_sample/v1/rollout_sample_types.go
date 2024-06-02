package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=rollout
// +kubebuilder:printcolumn:name="Replicas",type=integer,JSONPath=`.spec.replicas`
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RolloutSample is a specification for a RolloutSample resource
type RolloutSample struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec RolloutSampleSpec `json:"spec"`
	// +optional
	Status RolloutSampleStatus `json:"status"`
}

// +kubebuilder:object:root=true

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

// +kubebuilder:object:root=true

// RolloutSampleStatus is the status for a RolloutSample resource
type RolloutSampleStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RolloutSampleList is a list of RolloutSample resources
type RolloutSampleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RolloutSample `json:"items"`
}
