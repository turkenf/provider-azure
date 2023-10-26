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

type VirtualWANInitParameters struct {

	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to true.
	AllowBranchToBranchTraffic *bool `json:"allowBranchToBranchTraffic,omitempty" tf:"allow_branch_to_branch_traffic,omitempty"`

	// Boolean flag to specify whether VPN encryption is disabled. Defaults to false.
	DisableVPNEncryption *bool `json:"disableVpnEncryption,omitempty" tf:"disable_vpn_encryption,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the Office365 local breakout category. Possible values include: Optimize, OptimizeAndAllow, All, None. Defaults to None.
	Office365LocalBreakoutCategory *string `json:"office365LocalBreakoutCategory,omitempty" tf:"office365_local_breakout_category,omitempty"`

	// A mapping of tags to assign to the Virtual WAN.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Virtual WAN type. Possible Values include: Basic and Standard. Defaults to Standard.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type VirtualWANObservation struct {

	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to true.
	AllowBranchToBranchTraffic *bool `json:"allowBranchToBranchTraffic,omitempty" tf:"allow_branch_to_branch_traffic,omitempty"`

	// Boolean flag to specify whether VPN encryption is disabled. Defaults to false.
	DisableVPNEncryption *bool `json:"disableVpnEncryption,omitempty" tf:"disable_vpn_encryption,omitempty"`

	// The ID of the Virtual WAN.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the Office365 local breakout category. Possible values include: Optimize, OptimizeAndAllow, All, None. Defaults to None.
	Office365LocalBreakoutCategory *string `json:"office365LocalBreakoutCategory,omitempty" tf:"office365_local_breakout_category,omitempty"`

	// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the Virtual WAN.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Virtual WAN type. Possible Values include: Basic and Standard. Defaults to Standard.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type VirtualWANParameters struct {

	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to true.
	// +kubebuilder:validation:Optional
	AllowBranchToBranchTraffic *bool `json:"allowBranchToBranchTraffic,omitempty" tf:"allow_branch_to_branch_traffic,omitempty"`

	// Boolean flag to specify whether VPN encryption is disabled. Defaults to false.
	// +kubebuilder:validation:Optional
	DisableVPNEncryption *bool `json:"disableVpnEncryption,omitempty" tf:"disable_vpn_encryption,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the Office365 local breakout category. Possible values include: Optimize, OptimizeAndAllow, All, None. Defaults to None.
	// +kubebuilder:validation:Optional
	Office365LocalBreakoutCategory *string `json:"office365LocalBreakoutCategory,omitempty" tf:"office365_local_breakout_category,omitempty"`

	// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the Virtual WAN.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Virtual WAN type. Possible Values include: Basic and Standard. Defaults to Standard.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// VirtualWANSpec defines the desired state of VirtualWAN
type VirtualWANSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualWANParameters `json:"forProvider"`
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
	InitProvider VirtualWANInitParameters `json:"initProvider,omitempty"`
}

// VirtualWANStatus defines the observed state of VirtualWAN.
type VirtualWANStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualWANObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualWAN is the Schema for the VirtualWANs API. Manages a Virtual WAN.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualWAN struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   VirtualWANSpec   `json:"spec"`
	Status VirtualWANStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualWANList contains a list of VirtualWANs
type VirtualWANList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualWAN `json:"items"`
}

// Repository type metadata.
var (
	VirtualWAN_Kind             = "VirtualWAN"
	VirtualWAN_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualWAN_Kind}.String()
	VirtualWAN_KindAPIVersion   = VirtualWAN_Kind + "." + CRDGroupVersion.String()
	VirtualWAN_GroupVersionKind = CRDGroupVersion.WithKind(VirtualWAN_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualWAN{}, &VirtualWANList{})
}
