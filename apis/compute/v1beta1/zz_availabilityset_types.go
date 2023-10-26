// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type AvailabilitySetInitParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies whether the availability set is managed or not. Possible values are true (to specify aligned) or false (to specify classic). Default is true. Changing this forces a new resource to be created.
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// Specifies the number of fault domains that are used. Defaults to 3. Changing this forces a new resource to be created.
	PlatformFaultDomainCount *float64 `json:"platformFaultDomainCount,omitempty" tf:"platform_fault_domain_count,omitempty"`

	// Specifies the number of update domains that are used. Defaults to 5. Changing this forces a new resource to be created.
	PlatformUpdateDomainCount *float64 `json:"platformUpdateDomainCount,omitempty" tf:"platform_update_domain_count,omitempty"`

	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AvailabilitySetObservation struct {

	// The ID of the Availability Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies whether the availability set is managed or not. Possible values are true (to specify aligned) or false (to specify classic). Default is true. Changing this forces a new resource to be created.
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// Specifies the number of fault domains that are used. Defaults to 3. Changing this forces a new resource to be created.
	PlatformFaultDomainCount *float64 `json:"platformFaultDomainCount,omitempty" tf:"platform_fault_domain_count,omitempty"`

	// Specifies the number of update domains that are used. Defaults to 5. Changing this forces a new resource to be created.
	PlatformUpdateDomainCount *float64 `json:"platformUpdateDomainCount,omitempty" tf:"platform_update_domain_count,omitempty"`

	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id,omitempty"`

	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AvailabilitySetParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies whether the availability set is managed or not. Possible values are true (to specify aligned) or false (to specify classic). Default is true. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// Specifies the number of fault domains that are used. Defaults to 3. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PlatformFaultDomainCount *float64 `json:"platformFaultDomainCount,omitempty" tf:"platform_fault_domain_count,omitempty"`

	// Specifies the number of update domains that are used. Defaults to 5. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PlatformUpdateDomainCount *float64 `json:"platformUpdateDomainCount,omitempty" tf:"platform_update_domain_count,omitempty"`

	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id,omitempty"`

	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AvailabilitySetSpec defines the desired state of AvailabilitySet
type AvailabilitySetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AvailabilitySetParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AvailabilitySetInitParameters `json:"initProvider,omitempty"`
}

// AvailabilitySetStatus defines the observed state of AvailabilitySet.
type AvailabilitySetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AvailabilitySetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AvailabilitySet is the Schema for the AvailabilitySets API. Manages an Availability Set for Virtual Machines.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AvailabilitySet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   AvailabilitySetSpec   `json:"spec"`
	Status AvailabilitySetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AvailabilitySetList contains a list of AvailabilitySets
type AvailabilitySetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AvailabilitySet `json:"items"`
}

// Repository type metadata.
var (
	AvailabilitySet_Kind             = "AvailabilitySet"
	AvailabilitySet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AvailabilitySet_Kind}.String()
	AvailabilitySet_KindAPIVersion   = AvailabilitySet_Kind + "." + CRDGroupVersion.String()
	AvailabilitySet_GroupVersionKind = CRDGroupVersion.WithKind(AvailabilitySet_Kind)
)

func init() {
	SchemeBuilder.Register(&AvailabilitySet{}, &AvailabilitySetList{})
}
