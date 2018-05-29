package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KubeVirtList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KubeVirt `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KubeVirt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              KubeVirtSpec   `json:"spec"`
	Status            KubeVirtStatus `json:"status,omitempty"`
}

type KubeVirtSpec struct {
	// Fill me
}
type KubeVirtStatus struct {
	// Fill me
}
