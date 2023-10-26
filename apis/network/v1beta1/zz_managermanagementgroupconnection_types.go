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

type ManagerManagementGroupConnectionInitParameters struct {

	// A description of the Network Manager Management Group Connection.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type ManagerManagementGroupConnectionObservation struct {

	// The Connection state of the Network Manager Management Group Connection.
	ConnectionState *string `json:"connectionState,omitempty" tf:"connection_state,omitempty"`

	// A description of the Network Manager Management Group Connection.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Network Manager Management Group Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the target Management Group. Changing this forces a new resource to be created.
	ManagementGroupID *string `json:"managementGroupId,omitempty" tf:"management_group_id,omitempty"`

	// Specifies the ID of the Network Manager which the Management Group is connected to. Changing this forces a new resource to be created.
	NetworkManagerID *string `json:"networkManagerId,omitempty" tf:"network_manager_id,omitempty"`
}

type ManagerManagementGroupConnectionParameters struct {

	// A description of the Network Manager Management Group Connection.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the ID of the target Management Group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/management/v1beta1.ManagementGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ManagementGroupID *string `json:"managementGroupId,omitempty" tf:"management_group_id,omitempty"`

	// Reference to a ManagementGroup in management to populate managementGroupId.
	// +kubebuilder:validation:Optional
	ManagementGroupIDRef *v1.Reference `json:"managementGroupIdRef,omitempty" tf:"-"`

	// Selector for a ManagementGroup in management to populate managementGroupId.
	// +kubebuilder:validation:Optional
	ManagementGroupIDSelector *v1.Selector `json:"managementGroupIdSelector,omitempty" tf:"-"`

	// Specifies the ID of the Network Manager which the Management Group is connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Manager
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkManagerID *string `json:"networkManagerId,omitempty" tf:"network_manager_id,omitempty"`

	// Reference to a Manager in network to populate networkManagerId.
	// +kubebuilder:validation:Optional
	NetworkManagerIDRef *v1.Reference `json:"networkManagerIdRef,omitempty" tf:"-"`

	// Selector for a Manager in network to populate networkManagerId.
	// +kubebuilder:validation:Optional
	NetworkManagerIDSelector *v1.Selector `json:"networkManagerIdSelector,omitempty" tf:"-"`
}

// ManagerManagementGroupConnectionSpec defines the desired state of ManagerManagementGroupConnection
type ManagerManagementGroupConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagerManagementGroupConnectionParameters `json:"forProvider"`
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
	InitProvider ManagerManagementGroupConnectionInitParameters `json:"initProvider,omitempty"`
}

// ManagerManagementGroupConnectionStatus defines the observed state of ManagerManagementGroupConnection.
type ManagerManagementGroupConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagerManagementGroupConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerManagementGroupConnection is the Schema for the ManagerManagementGroupConnections API. Manages a Network Manager Management Group Connection.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagerManagementGroupConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagerManagementGroupConnectionSpec   `json:"spec"`
	Status            ManagerManagementGroupConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerManagementGroupConnectionList contains a list of ManagerManagementGroupConnections
type ManagerManagementGroupConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagerManagementGroupConnection `json:"items"`
}

// Repository type metadata.
var (
	ManagerManagementGroupConnection_Kind             = "ManagerManagementGroupConnection"
	ManagerManagementGroupConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagerManagementGroupConnection_Kind}.String()
	ManagerManagementGroupConnection_KindAPIVersion   = ManagerManagementGroupConnection_Kind + "." + CRDGroupVersion.String()
	ManagerManagementGroupConnection_GroupVersionKind = CRDGroupVersion.WithKind(ManagerManagementGroupConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagerManagementGroupConnection{}, &ManagerManagementGroupConnectionList{})
}
