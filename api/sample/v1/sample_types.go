package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These const variables are used in our custom controller.
const (
	GroupName string = "sample.zlw.com"
	Version   string = "v1"
	//Plural    string = "samples"
	//Singluar  string = "sample"
	//ShortName string = "sample"
	//Name      string = Plural + "." + GroupName
)

// Sample is a specification for a Sample resource
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=demo
// +kubebuilder:printcolumn:name="Replicas",type=int,JSONPath=`.spec.replicas`
type Sample struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SampleSpec   `json:"spec"`
	Status SampleStatus `json:"status"`
}

// SampleSpec is the spec for a Foo resource
type SampleSpec struct {
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Maximum=15
	Replicas *int32 `json:"replicas"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=15
	DeploymentName string `json:"deploymentName"`
}

// SampleStatus is the status for a Sample resource
type SampleStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// SampleList is a list of Sample resources
// +kubebuilder:object:root=true
type SampleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Sample `json:"items"`
}
