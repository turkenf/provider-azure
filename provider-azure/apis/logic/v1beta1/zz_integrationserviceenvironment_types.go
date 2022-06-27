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

type IntegrationServiceEnvironmentObservation struct {
	ConnectorEndpointIPAddresses []*string `json:"connectorEndpointIpAddresses,omitempty" tf:"connector_endpoint_ip_addresses,omitempty"`

	ConnectorOutboundIPAddresses []*string `json:"connectorOutboundIpAddresses,omitempty" tf:"connector_outbound_ip_addresses,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	WorkflowEndpointIPAddresses []*string `json:"workflowEndpointIpAddresses,omitempty" tf:"workflow_endpoint_ip_addresses,omitempty"`

	WorkflowOutboundIPAddresses []*string `json:"workflowOutboundIpAddresses,omitempty" tf:"workflow_outbound_ip_addresses,omitempty"`
}

type IntegrationServiceEnvironmentParameters struct {

	// +kubebuilder:validation:Required
	AccessEndpointType *string `json:"accessEndpointType" tf:"access_endpoint_type,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VirtualNetworkSubnetIds []*string `json:"virtualNetworkSubnetIds" tf:"virtual_network_subnet_ids,omitempty"`
}

// IntegrationServiceEnvironmentSpec defines the desired state of IntegrationServiceEnvironment
type IntegrationServiceEnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationServiceEnvironmentParameters `json:"forProvider"`
}

// IntegrationServiceEnvironmentStatus defines the observed state of IntegrationServiceEnvironment.
type IntegrationServiceEnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationServiceEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationServiceEnvironment is the Schema for the IntegrationServiceEnvironments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IntegrationServiceEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationServiceEnvironmentSpec   `json:"spec"`
	Status            IntegrationServiceEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationServiceEnvironmentList contains a list of IntegrationServiceEnvironments
type IntegrationServiceEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IntegrationServiceEnvironment `json:"items"`
}

// Repository type metadata.
var (
	IntegrationServiceEnvironment_Kind             = "IntegrationServiceEnvironment"
	IntegrationServiceEnvironment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IntegrationServiceEnvironment_Kind}.String()
	IntegrationServiceEnvironment_KindAPIVersion   = IntegrationServiceEnvironment_Kind + "." + CRDGroupVersion.String()
	IntegrationServiceEnvironment_GroupVersionKind = CRDGroupVersion.WithKind(IntegrationServiceEnvironment_Kind)
)

func init() {
	SchemeBuilder.Register(&IntegrationServiceEnvironment{}, &IntegrationServiceEnvironmentList{})
}
