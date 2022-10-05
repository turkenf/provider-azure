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

type ExpressRouteGatewayObservation struct {

	// The ID of the ExpressRoute gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ExpressRouteGatewayParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the resource group in which to create the ExpressRoute gateway. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The number of scale units with which to provision the ExpressRoute gateway. Each scale unit is equal to 2Gbps, with support for up to 10 scale units (20Gbps).
	// +kubebuilder:validation:Required
	ScaleUnits *float64 `json:"scaleUnits" tf:"scale_units,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of a Virtual HUB within which the ExpressRoute gateway should be created.
	// +crossplane:generate:reference:type=VirtualHub
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualHubID *string `json:"virtualHubId,omitempty" tf:"virtual_hub_id,omitempty"`

	// Reference to a VirtualHub to populate virtualHubId.
	// +kubebuilder:validation:Optional
	VirtualHubIDRef *v1.Reference `json:"virtualHubIdRef,omitempty" tf:"-"`

	// Selector for a VirtualHub to populate virtualHubId.
	// +kubebuilder:validation:Optional
	VirtualHubIDSelector *v1.Selector `json:"virtualHubIdSelector,omitempty" tf:"-"`
}

// ExpressRouteGatewaySpec defines the desired state of ExpressRouteGateway
type ExpressRouteGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRouteGatewayParameters `json:"forProvider"`
}

// ExpressRouteGatewayStatus defines the observed state of ExpressRouteGateway.
type ExpressRouteGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRouteGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteGateway is the Schema for the ExpressRouteGateways API. Manages an ExpressRoute gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRouteGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteGatewaySpec   `json:"spec"`
	Status            ExpressRouteGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteGatewayList contains a list of ExpressRouteGateways
type ExpressRouteGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRouteGateway `json:"items"`
}

// Repository type metadata.
var (
	ExpressRouteGateway_Kind             = "ExpressRouteGateway"
	ExpressRouteGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExpressRouteGateway_Kind}.String()
	ExpressRouteGateway_KindAPIVersion   = ExpressRouteGateway_Kind + "." + CRDGroupVersion.String()
	ExpressRouteGateway_GroupVersionKind = CRDGroupVersion.WithKind(ExpressRouteGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&ExpressRouteGateway{}, &ExpressRouteGatewayList{})
}