/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BlobInventoryPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BlobInventoryPolicyParameters struct {

	// +kubebuilder:validation:Required
	Rules []RulesParameters `json:"rules" tf:"rules,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// +kubebuilder:validation:Required
	BlobTypes []*string `json:"blobTypes" tf:"blob_types,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeBlobVersions *bool `json:"includeBlobVersions,omitempty" tf:"include_blob_versions,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeSnapshots *bool `json:"includeSnapshots,omitempty" tf:"include_snapshots,omitempty"`

	// +kubebuilder:validation:Optional
	PrefixMatch []*string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`
}

type RulesObservation struct {
}

type RulesParameters struct {

	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Required
	Format *string `json:"format" tf:"format,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`

	// +kubebuilder:validation:Required
	SchemaFields []*string `json:"schemaFields" tf:"schema_fields,omitempty"`

	// +kubebuilder:validation:Required
	Scope *string `json:"scope" tf:"scope,omitempty"`

	// +kubebuilder:validation:Required
	StorageContainerName *string `json:"storageContainerName" tf:"storage_container_name,omitempty"`
}

// BlobInventoryPolicySpec defines the desired state of BlobInventoryPolicy
type BlobInventoryPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BlobInventoryPolicyParameters `json:"forProvider"`
}

// BlobInventoryPolicyStatus defines the observed state of BlobInventoryPolicy.
type BlobInventoryPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BlobInventoryPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BlobInventoryPolicy is the Schema for the BlobInventoryPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BlobInventoryPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BlobInventoryPolicySpec   `json:"spec"`
	Status            BlobInventoryPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlobInventoryPolicyList contains a list of BlobInventoryPolicys
type BlobInventoryPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BlobInventoryPolicy `json:"items"`
}

// Repository type metadata.
var (
	BlobInventoryPolicy_Kind             = "BlobInventoryPolicy"
	BlobInventoryPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BlobInventoryPolicy_Kind}.String()
	BlobInventoryPolicy_KindAPIVersion   = BlobInventoryPolicy_Kind + "." + CRDGroupVersion.String()
	BlobInventoryPolicy_GroupVersionKind = CRDGroupVersion.WithKind(BlobInventoryPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&BlobInventoryPolicy{}, &BlobInventoryPolicyList{})
}
