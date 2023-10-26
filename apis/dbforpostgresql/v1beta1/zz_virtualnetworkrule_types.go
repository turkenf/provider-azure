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

type VirtualNetworkRuleInitParameters struct {

	// Should the Virtual Network Rule be created before the Subnet has the Virtual Network Service Endpoint enabled?
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`
}

type VirtualNetworkRuleObservation struct {

	// The ID of the PostgreSQL Virtual Network Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Should the Virtual Network Rule be created before the Subnet has the Virtual Network Service Endpoint enabled?
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// The name of the resource group where the PostgreSQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the SQL Server to which this PostgreSQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// The ID of the subnet that the PostgreSQL server will be connected to.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type VirtualNetworkRuleParameters struct {

	// Should the Virtual Network Rule be created before the Subnet has the Virtual Network Service Endpoint enabled?
	// +kubebuilder:validation:Optional
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// The name of the resource group where the PostgreSQL server resides. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the SQL Server to which this PostgreSQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Server
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Reference to a Server to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// Selector for a Server to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`

	// The ID of the subnet that the PostgreSQL server will be connected to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// VirtualNetworkRuleSpec defines the desired state of VirtualNetworkRule
type VirtualNetworkRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNetworkRuleParameters `json:"forProvider"`
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
	InitProvider VirtualNetworkRuleInitParameters `json:"initProvider,omitempty"`
}

// VirtualNetworkRuleStatus defines the observed state of VirtualNetworkRule.
type VirtualNetworkRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNetworkRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkRule is the Schema for the VirtualNetworkRules API. Manages a PostgreSQL Virtual Network Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkRuleSpec   `json:"spec"`
	Status            VirtualNetworkRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkRuleList contains a list of VirtualNetworkRules
type VirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkRule `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetworkRule_Kind             = "VirtualNetworkRule"
	VirtualNetworkRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualNetworkRule_Kind}.String()
	VirtualNetworkRule_KindAPIVersion   = VirtualNetworkRule_Kind + "." + CRDGroupVersion.String()
	VirtualNetworkRule_GroupVersionKind = CRDGroupVersion.WithKind(VirtualNetworkRule_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetworkRule{}, &VirtualNetworkRuleList{})
}
