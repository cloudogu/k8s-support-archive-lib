package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatusPhase string

const (
	ConditionSupportArchiveCreated = "Created"
)

// SupportArchiveSpec defines the desired state of SupportArchive.
type SupportArchiveSpec struct {
	// ExcludedContents defines which contents should not be included in the SupportArchive.
	ExcludedContents ExcludedContents `json:"excludedContents,omitempty"`
	// LoggingConfig defines how logs should be collected.
	LoggingConfig LoggingConfig `json:"loggingConfig,omitempty"`
}

type ExcludedContents struct {
	// SystemState concerns all Kubernetes resources (excluding Secrets) with label `app: ces`.
	SystemState bool `json:"systemState,omitempty"`
	// SensitiveData concerns Secrets with label `app: ces`.
	// They will be censored even if included.
	SensitiveData bool `json:"sensitiveData,omitempty"`
	// Events concerns Kubernetes events.
	Events bool `json:"events,omitempty"`
	// LogsAndEvents concerns application logs.
	Logs bool `json:"logs,omitempty"`
	// VolumeInfo concerns metrics about volumes.
	VolumeInfo bool `json:"volumeInfo,omitempty"`
	// SystemInfo concerns information about the system like the kubernetes version and nodes.
	SystemInfo bool `json:"systemInfo,omitempty"`
}

type LoggingConfig struct {
	// StartTime is the minimal time from when logs and events should be included.
	StartTime metav1.Time `json:"startTime,omitempty"`
	// EndTime is the maximal time from when logs and events should be included.
	EndTime metav1.Time `json:"endTime,omitempty"`
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
