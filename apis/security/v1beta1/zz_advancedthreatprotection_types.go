/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AdvancedThreatProtectionObservation struct {

	// The ID of the Advanced Threat Protection resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AdvancedThreatProtectionParameters struct {

	// Should Advanced Threat Protection be enabled on this resource?
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	TargetResourceID *string `json:"targetResourceId" tf:"target_resource_id,omitempty"`
}

// AdvancedThreatProtectionSpec defines the desired state of AdvancedThreatProtection
type AdvancedThreatProtectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AdvancedThreatProtectionParameters `json:"forProvider"`
}

// AdvancedThreatProtectionStatus defines the observed state of AdvancedThreatProtection.
type AdvancedThreatProtectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AdvancedThreatProtectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AdvancedThreatProtection is the Schema for the AdvancedThreatProtections API. Manages a resources Advanced Threat Protection setting.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AdvancedThreatProtection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AdvancedThreatProtectionSpec   `json:"spec"`
	Status            AdvancedThreatProtectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AdvancedThreatProtectionList contains a list of AdvancedThreatProtections
type AdvancedThreatProtectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdvancedThreatProtection `json:"items"`
}

// Repository type metadata.
var (
	AdvancedThreatProtection_Kind             = "AdvancedThreatProtection"
	AdvancedThreatProtection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AdvancedThreatProtection_Kind}.String()
	AdvancedThreatProtection_KindAPIVersion   = AdvancedThreatProtection_Kind + "." + CRDGroupVersion.String()
	AdvancedThreatProtection_GroupVersionKind = CRDGroupVersion.WithKind(AdvancedThreatProtection_Kind)
)

func init() {
	SchemeBuilder.Register(&AdvancedThreatProtection{}, &AdvancedThreatProtectionList{})
}