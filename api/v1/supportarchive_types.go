package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SupportArchiveSpec defines the desired state of SupportArchive.
type SupportArchiveSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of SupportArchive. Edit supportarchive_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// SupportArchiveStatus defines the observed state of SupportArchive.
type SupportArchiveStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels=app=ces;app.kubernetes.io/name=k8s-support-archive-operator;k8s.cloudogu.com/component.name=k8s-support-archive-operator-crd

// SupportArchive is the Schema for the supportarchives API.
type SupportArchive struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SupportArchiveSpec   `json:"spec,omitempty"`
	Status SupportArchiveStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SupportArchiveList contains a list of SupportArchive.
type SupportArchiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SupportArchive `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SupportArchive{}, &SupportArchiveList{})
}
