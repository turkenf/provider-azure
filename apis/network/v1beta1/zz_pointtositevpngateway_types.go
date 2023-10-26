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

type ConnectionConfigurationInitParameters struct {

	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to false.
	InternetSecurityEnabled *bool `json:"internetSecurityEnabled,omitempty" tf:"internet_security_enabled,omitempty"`

	// The Name which should be used for this Connection Configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A route block as defined below.
	Route []RouteInitParameters `json:"route,omitempty" tf:"route,omitempty"`

	// A vpn_client_address_pool block as defined below.
	VPNClientAddressPool []VPNClientAddressPoolInitParameters `json:"vpnClientAddressPool,omitempty" tf:"vpn_client_address_pool,omitempty"`
}

type ConnectionConfigurationObservation struct {

	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to false.
	InternetSecurityEnabled *bool `json:"internetSecurityEnabled,omitempty" tf:"internet_security_enabled,omitempty"`

	// The Name which should be used for this Connection Configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A route block as defined below.
	Route []RouteObservation `json:"route,omitempty" tf:"route,omitempty"`

	// A vpn_client_address_pool block as defined below.
	VPNClientAddressPool []VPNClientAddressPoolObservation `json:"vpnClientAddressPool,omitempty" tf:"vpn_client_address_pool,omitempty"`
}

type ConnectionConfigurationParameters struct {

	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to false.
	// +kubebuilder:validation:Optional
	InternetSecurityEnabled *bool `json:"internetSecurityEnabled,omitempty" tf:"internet_security_enabled,omitempty"`

	// The Name which should be used for this Connection Configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A route block as defined below.
	// +kubebuilder:validation:Optional
	Route []RouteParameters `json:"route,omitempty" tf:"route,omitempty"`

	// A vpn_client_address_pool block as defined below.
	// +kubebuilder:validation:Optional
	VPNClientAddressPool []VPNClientAddressPoolParameters `json:"vpnClientAddressPool" tf:"vpn_client_address_pool,omitempty"`
}

type PointToSiteVPNGatewayInitParameters struct {

	// A connection_configuration block as defined below.
	ConnectionConfiguration []ConnectionConfigurationInitParameters `json:"connectionConfiguration,omitempty" tf:"connection_configuration,omitempty"`

	// A list of IP Addresses of DNS Servers for the Point-to-Site VPN Gateway.
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Is the Routing Preference for the Public IP Interface of the VPN Gateway enabled? Defaults to false. Changing this forces a new resource to be created.
	RoutingPreferenceInternetEnabled *bool `json:"routingPreferenceInternetEnabled,omitempty" tf:"routing_preference_internet_enabled,omitempty"`

	// The Scale Unit for this Point-to-Site VPN Gateway.
	ScaleUnit *float64 `json:"scaleUnit,omitempty" tf:"scale_unit,omitempty"`

	// A mapping of tags to assign to the Point-to-Site VPN Gateway.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PointToSiteVPNGatewayObservation struct {

	// A connection_configuration block as defined below.
	ConnectionConfiguration []ConnectionConfigurationObservation `json:"connectionConfiguration,omitempty" tf:"connection_configuration,omitempty"`

	// A list of IP Addresses of DNS Servers for the Point-to-Site VPN Gateway.
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// The ID of the Point-to-Site VPN Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Is the Routing Preference for the Public IP Interface of the VPN Gateway enabled? Defaults to false. Changing this forces a new resource to be created.
	RoutingPreferenceInternetEnabled *bool `json:"routingPreferenceInternetEnabled,omitempty" tf:"routing_preference_internet_enabled,omitempty"`

	// The Scale Unit for this Point-to-Site VPN Gateway.
	ScaleUnit *float64 `json:"scaleUnit,omitempty" tf:"scale_unit,omitempty"`

	// A mapping of tags to assign to the Point-to-Site VPN Gateway.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
	VPNServerConfigurationID *string `json:"vpnServerConfigurationId,omitempty" tf:"vpn_server_configuration_id,omitempty"`

	// The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
	VirtualHubID *string `json:"virtualHubId,omitempty" tf:"virtual_hub_id,omitempty"`
}

type PointToSiteVPNGatewayParameters struct {

	// A connection_configuration block as defined below.
	// +kubebuilder:validation:Optional
	ConnectionConfiguration []ConnectionConfigurationParameters `json:"connectionConfiguration,omitempty" tf:"connection_configuration,omitempty"`

	// A list of IP Addresses of DNS Servers for the Point-to-Site VPN Gateway.
	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Is the Routing Preference for the Public IP Interface of the VPN Gateway enabled? Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	RoutingPreferenceInternetEnabled *bool `json:"routingPreferenceInternetEnabled,omitempty" tf:"routing_preference_internet_enabled,omitempty"`

	// The Scale Unit for this Point-to-Site VPN Gateway.
	// +kubebuilder:validation:Optional
	ScaleUnit *float64 `json:"scaleUnit,omitempty" tf:"scale_unit,omitempty"`

	// A mapping of tags to assign to the Point-to-Site VPN Gateway.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=VPNServerConfiguration
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPNServerConfigurationID *string `json:"vpnServerConfigurationId,omitempty" tf:"vpn_server_configuration_id,omitempty"`

	// Reference to a VPNServerConfiguration to populate vpnServerConfigurationId.
	// +kubebuilder:validation:Optional
	VPNServerConfigurationIDRef *v1.Reference `json:"vpnServerConfigurationIdRef,omitempty" tf:"-"`

	// Selector for a VPNServerConfiguration to populate vpnServerConfigurationId.
	// +kubebuilder:validation:Optional
	VPNServerConfigurationIDSelector *v1.Selector `json:"vpnServerConfigurationIdSelector,omitempty" tf:"-"`

	// The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=VirtualHub
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualHubID *string `json:"virtualHubId,omitempty" tf:"virtual_hub_id,omitempty"`

	// Reference to a VirtualHub to populate virtualHubId.
	// +kubebuilder:validation:Optional
	VirtualHubIDRef *v1.Reference `json:"virtualHubIdRef,omitempty" tf:"-"`

	// Selector for a VirtualHub to populate virtualHubId.
	// +kubebuilder:validation:Optional
	VirtualHubIDSelector *v1.Selector `json:"virtualHubIdSelector,omitempty" tf:"-"`
}

type RouteInitParameters struct {

	// The Virtual Hub Route Table resource id associated with this Routing Configuration.
	AssociatedRouteTableID *string `json:"associatedRouteTableId,omitempty" tf:"associated_route_table_id,omitempty"`

	// The resource ID of the Route Map associated with this Routing Configuration for inbound learned routes.
	InboundRouteMapID *string `json:"inboundRouteMapId,omitempty" tf:"inbound_route_map_id,omitempty"`

	// The resource ID of the Route Map associated with this Routing Configuration for outbound advertised routes.
	OutboundRouteMapID *string `json:"outboundRouteMapId,omitempty" tf:"outbound_route_map_id,omitempty"`

	// A propagated_route_table block as defined below.
	PropagatedRouteTable []RoutePropagatedRouteTableInitParameters `json:"propagatedRouteTable,omitempty" tf:"propagated_route_table,omitempty"`
}

type RouteObservation struct {

	// The Virtual Hub Route Table resource id associated with this Routing Configuration.
	AssociatedRouteTableID *string `json:"associatedRouteTableId,omitempty" tf:"associated_route_table_id,omitempty"`

	// The resource ID of the Route Map associated with this Routing Configuration for inbound learned routes.
	InboundRouteMapID *string `json:"inboundRouteMapId,omitempty" tf:"inbound_route_map_id,omitempty"`

	// The resource ID of the Route Map associated with this Routing Configuration for outbound advertised routes.
	OutboundRouteMapID *string `json:"outboundRouteMapId,omitempty" tf:"outbound_route_map_id,omitempty"`

	// A propagated_route_table block as defined below.
	PropagatedRouteTable []RoutePropagatedRouteTableObservation `json:"propagatedRouteTable,omitempty" tf:"propagated_route_table,omitempty"`
}

type RouteParameters struct {

	// The Virtual Hub Route Table resource id associated with this Routing Configuration.
	// +kubebuilder:validation:Optional
	AssociatedRouteTableID *string `json:"associatedRouteTableId" tf:"associated_route_table_id,omitempty"`

	// The resource ID of the Route Map associated with this Routing Configuration for inbound learned routes.
	// +kubebuilder:validation:Optional
	InboundRouteMapID *string `json:"inboundRouteMapId,omitempty" tf:"inbound_route_map_id,omitempty"`

	// The resource ID of the Route Map associated with this Routing Configuration for outbound advertised routes.
	// +kubebuilder:validation:Optional
	OutboundRouteMapID *string `json:"outboundRouteMapId,omitempty" tf:"outbound_route_map_id,omitempty"`

	// A propagated_route_table block as defined below.
	// +kubebuilder:validation:Optional
	PropagatedRouteTable []RoutePropagatedRouteTableParameters `json:"propagatedRouteTable,omitempty" tf:"propagated_route_table,omitempty"`
}

type RoutePropagatedRouteTableInitParameters struct {

	// The list of Virtual Hub Route Table resource id which the routes will be propagated to.
	Ids []*string `json:"ids,omitempty" tf:"ids,omitempty"`

	// The list of labels to logically group Virtual Hub Route Tables which the routes will be propagated to.
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type RoutePropagatedRouteTableObservation struct {

	// The list of Virtual Hub Route Table resource id which the routes will be propagated to.
	Ids []*string `json:"ids,omitempty" tf:"ids,omitempty"`

	// The list of labels to logically group Virtual Hub Route Tables which the routes will be propagated to.
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type RoutePropagatedRouteTableParameters struct {

	// The list of Virtual Hub Route Table resource id which the routes will be propagated to.
	// +kubebuilder:validation:Optional
	Ids []*string `json:"ids" tf:"ids,omitempty"`

	// The list of labels to logically group Virtual Hub Route Tables which the routes will be propagated to.
	// +kubebuilder:validation:Optional
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type VPNClientAddressPoolInitParameters struct {

	// A list of CIDR Ranges which should be used as Address Prefixes.
	AddressPrefixes []*string `json:"addressPrefixes,omitempty" tf:"address_prefixes,omitempty"`
}

type VPNClientAddressPoolObservation struct {

	// A list of CIDR Ranges which should be used as Address Prefixes.
	AddressPrefixes []*string `json:"addressPrefixes,omitempty" tf:"address_prefixes,omitempty"`
}

type VPNClientAddressPoolParameters struct {

	// A list of CIDR Ranges which should be used as Address Prefixes.
	// +kubebuilder:validation:Optional
	AddressPrefixes []*string `json:"addressPrefixes" tf:"address_prefixes,omitempty"`
}

// PointToSiteVPNGatewaySpec defines the desired state of PointToSiteVPNGateway
type PointToSiteVPNGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PointToSiteVPNGatewayParameters `json:"forProvider"`
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
	InitProvider PointToSiteVPNGatewayInitParameters `json:"initProvider,omitempty"`
}

// PointToSiteVPNGatewayStatus defines the observed state of PointToSiteVPNGateway.
type PointToSiteVPNGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PointToSiteVPNGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PointToSiteVPNGateway is the Schema for the PointToSiteVPNGateways API. Manages a Point-to-Site VPN Gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PointToSiteVPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionConfiguration) || (has(self.initProvider) && has(self.initProvider.connectionConfiguration))",message="spec.forProvider.connectionConfiguration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scaleUnit) || (has(self.initProvider) && has(self.initProvider.scaleUnit))",message="spec.forProvider.scaleUnit is a required parameter"
	Spec   PointToSiteVPNGatewaySpec   `json:"spec"`
	Status PointToSiteVPNGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PointToSiteVPNGatewayList contains a list of PointToSiteVPNGateways
type PointToSiteVPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PointToSiteVPNGateway `json:"items"`
}

// Repository type metadata.
var (
	PointToSiteVPNGateway_Kind             = "PointToSiteVPNGateway"
	PointToSiteVPNGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PointToSiteVPNGateway_Kind}.String()
	PointToSiteVPNGateway_KindAPIVersion   = PointToSiteVPNGateway_Kind + "." + CRDGroupVersion.String()
	PointToSiteVPNGateway_GroupVersionKind = CRDGroupVersion.WithKind(PointToSiteVPNGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&PointToSiteVPNGateway{}, &PointToSiteVPNGatewayList{})
}
