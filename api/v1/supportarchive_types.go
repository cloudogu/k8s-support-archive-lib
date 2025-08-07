package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatusPhase string

const (
	ConditionSupportArchiveCreated = "Created"
	ConditionVolumeInfoFetched     = "VolumeInfoFetched"
	ConditionNodeInfoFetched       = "NodeInfoFetched"
)

// SupportArchiveSpec defines the desired state of SupportArchive.
type SupportArchiveSpec struct {
	// ExcludedContents defines which contents should not be included in the SupportArchive.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ExcludedContents is immutable"
	ExcludedContents ExcludedContents `json:"excludedContents"`
	// ContentTimeframe defines the timeframe of the contents in the supportArchive.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ContentTimeframe is immutable"
	ContentTimeframe ContentTimeframe `json:"contentTimeframe"`
}

type ExcludedContents struct {
	// SystemState concerns all Kubernetes resources (excluding Secrets) with label `app: ces`.
	// +required
	SystemState bool `json:"systemState"`
	// SensitiveData concerns Secrets with label `app: ces`.
	// They will be censored even if included.
	// +required
	SensitiveData bool `json:"sensitiveData"`
	// Events concerns Kubernetes events.
	// +required
	Events bool `json:"events"`
	// Logs concerns application logs.
	// +required
	Logs bool `json:"logs"`
	// VolumeInfo concerns metrics about volumes.
	// +required
	VolumeInfo bool `json:"volumeInfo"`
	// SystemInfo concerns information about the system like the kubernetes version and nodes.
	// +required
	SystemInfo bool `json:"systemInfo"`
}

type ContentTimeframe struct {
	// StartTime is the minimal time from when logs and events should be included.
	// +required
	StartTime metav1.Time `json:"startTime"`
	// EndTime is the maximal time from when logs and events should be included.
	// +required
	EndTime metav1.Time `json:"endTime"`
}

// SupportArchiveStatus defines the observed state of SupportArchive.
type SupportArchiveStatus struct {
	// Errors contains error messages that accumulated during execution.
	Errors []string `json:"errors,omitempty"`
	// DownloadPath exposes where the created archive can be obtained.
	DownloadPath string `json:"downloadPath,omitempty"`
	// Conditions exposes the actual progress of the support archive creation.
	// +listType=map
	// +listMapKey=type
	// +patchStrategy=merge
	// +patchMergeKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels=app=ces;app.kubernetes.io/name=k8s-support-archive-operator;k8s.cloudogu.com/component.name=k8s-support-archive-operator-crd
// +kubebuilder:resource:shortName="sar"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="The age of the resource"

// SupportArchive is the Schema for the supportarchives API.
type SupportArchive struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SupportArchiveSpec   `json:"spec"`
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
