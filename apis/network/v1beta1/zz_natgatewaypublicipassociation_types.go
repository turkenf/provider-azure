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

type NATGatewayPublicIPAssociationInitParameters struct {
}

type NATGatewayPublicIPAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the NAT Gateway. Changing this forces a new resource to be created.
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// The ID of the Public IP which this NAT Gateway which should be connected to. Changing this forces a new resource to be created.
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`
}

type NATGatewayPublicIPAssociationParameters struct {

	// The ID of the NAT Gateway. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=NATGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// Reference to a NATGateway to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDRef *v1.Reference `json:"natGatewayIdRef,omitempty" tf:"-"`

	// Selector for a NATGateway to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDSelector *v1.Selector `json:"natGatewayIdSelector,omitempty" tf:"-"`

	// The ID of the Public IP which this NAT Gateway which should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=PublicIP
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`
}

// NATGatewayPublicIPAssociationSpec defines the desired state of NATGatewayPublicIPAssociation
type NATGatewayPublicIPAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NATGatewayPublicIPAssociationParameters `json:"forProvider"`
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
	InitProvider NATGatewayPublicIPAssociationInitParameters `json:"initProvider,omitempty"`
}

// NATGatewayPublicIPAssociationStatus defines the observed state of NATGatewayPublicIPAssociation.
type NATGatewayPublicIPAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NATGatewayPublicIPAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NATGatewayPublicIPAssociation is the Schema for the NATGatewayPublicIPAssociations API. Manages the association between a NAT Gateway and a Public IP.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NATGatewayPublicIPAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NATGatewayPublicIPAssociationSpec   `json:"spec"`
	Status            NATGatewayPublicIPAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NATGatewayPublicIPAssociationList contains a list of NATGatewayPublicIPAssociations
type NATGatewayPublicIPAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NATGatewayPublicIPAssociation `json:"items"`
}

// Repository type metadata.
var (
	NATGatewayPublicIPAssociation_Kind             = "NATGatewayPublicIPAssociation"
	NATGatewayPublicIPAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NATGatewayPublicIPAssociation_Kind}.String()
	NATGatewayPublicIPAssociation_KindAPIVersion   = NATGatewayPublicIPAssociation_Kind + "." + CRDGroupVersion.String()
	NATGatewayPublicIPAssociation_GroupVersionKind = CRDGroupVersion.WithKind(NATGatewayPublicIPAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&NATGatewayPublicIPAssociation{}, &NATGatewayPublicIPAssociationList{})
}
