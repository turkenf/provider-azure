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

type LoadBalancerFrontendIPConfigurationInitParameters struct {

	// The Frontend IP Configuration ID of a Gateway SKU Load Balancer.
	GatewayLoadBalancerFrontendIPConfigurationID *string `json:"gatewayLoadBalancerFrontendIpConfigurationId,omitempty" tf:"gateway_load_balancer_frontend_ip_configuration_id,omitempty"`

	// Specifies the name of the frontend IP configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The allocation method for the Private IP Address used by this Load Balancer. Possible values as Dynamic and Static.
	PrivateIPAddressAllocation *string `json:"privateIpAddressAllocation,omitempty" tf:"private_ip_address_allocation,omitempty"`

	// The version of IP that the Private IP Address is. Possible values are IPv4 or IPv6.
	PrivateIPAddressVersion *string `json:"privateIpAddressVersion,omitempty" tf:"private_ip_address_version,omitempty"`

	// The ID of a Public IP Prefix which should be associated with the Load Balancer. Public IP Prefix can only be used with outbound rules.
	PublicIPPrefixID *string `json:"publicIpPrefixId,omitempty" tf:"public_ip_prefix_id,omitempty"`

	// Specifies a list of Availability Zones in which the IP Address for this Load Balancer should be located.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type LoadBalancerFrontendIPConfigurationObservation struct {

	// The Frontend IP Configuration ID of a Gateway SKU Load Balancer.
	GatewayLoadBalancerFrontendIPConfigurationID *string `json:"gatewayLoadBalancerFrontendIpConfigurationId,omitempty" tf:"gateway_load_balancer_frontend_ip_configuration_id,omitempty"`

	// The id of the Frontend IP Configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The list of IDs of inbound rules that use this frontend IP.
	InboundNATRules []*string `json:"inboundNatRules,omitempty" tf:"inbound_nat_rules,omitempty"`

	// The list of IDs of load balancing rules that use this frontend IP.
	LoadBalancerRules []*string `json:"loadBalancerRules,omitempty" tf:"load_balancer_rules,omitempty"`

	// Specifies the name of the frontend IP configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The list of IDs outbound rules that use this frontend IP.
	OutboundRules []*string `json:"outboundRules,omitempty" tf:"outbound_rules,omitempty"`

	// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The allocation method for the Private IP Address used by this Load Balancer. Possible values as Dynamic and Static.
	PrivateIPAddressAllocation *string `json:"privateIpAddressAllocation,omitempty" tf:"private_ip_address_allocation,omitempty"`

	// The version of IP that the Private IP Address is. Possible values are IPv4 or IPv6.
	PrivateIPAddressVersion *string `json:"privateIpAddressVersion,omitempty" tf:"private_ip_address_version,omitempty"`

	// The ID of a Public IP Address which should be associated with the Load Balancer.
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// The ID of a Public IP Prefix which should be associated with the Load Balancer. Public IP Prefix can only be used with outbound rules.
	PublicIPPrefixID *string `json:"publicIpPrefixId,omitempty" tf:"public_ip_prefix_id,omitempty"`

	// The ID of the Subnet which should be associated with the IP Configuration.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Specifies a list of Availability Zones in which the IP Address for this Load Balancer should be located.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type LoadBalancerFrontendIPConfigurationParameters struct {

	// The Frontend IP Configuration ID of a Gateway SKU Load Balancer.
	// +kubebuilder:validation:Optional
	GatewayLoadBalancerFrontendIPConfigurationID *string `json:"gatewayLoadBalancerFrontendIpConfigurationId,omitempty" tf:"gateway_load_balancer_frontend_ip_configuration_id,omitempty"`

	// Specifies the name of the frontend IP configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
	// +kubebuilder:validation:Optional
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The allocation method for the Private IP Address used by this Load Balancer. Possible values as Dynamic and Static.
	// +kubebuilder:validation:Optional
	PrivateIPAddressAllocation *string `json:"privateIpAddressAllocation,omitempty" tf:"private_ip_address_allocation,omitempty"`

	// The version of IP that the Private IP Address is. Possible values are IPv4 or IPv6.
	// +kubebuilder:validation:Optional
	PrivateIPAddressVersion *string `json:"privateIpAddressVersion,omitempty" tf:"private_ip_address_version,omitempty"`

	// The ID of a Public IP Address which should be associated with the Load Balancer.
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

	// The ID of a Public IP Prefix which should be associated with the Load Balancer. Public IP Prefix can only be used with outbound rules.
	// +kubebuilder:validation:Optional
	PublicIPPrefixID *string `json:"publicIpPrefixId,omitempty" tf:"public_ip_prefix_id,omitempty"`

	// The ID of the Subnet which should be associated with the IP Configuration.
	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Specifies a list of Availability Zones in which the IP Address for this Load Balancer should be located.
	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type LoadBalancerInitParameters struct {

	// Specifies the Edge Zone within the Azure Region where this Load Balancer should exist. Changing this forces a new Load Balancer to be created.
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// One or multiple frontend_ip_configuration blocks as documented below.
	FrontendIPConfiguration []LoadBalancerFrontendIPConfigurationInitParameters `json:"frontendIpConfiguration,omitempty" tf:"frontend_ip_configuration,omitempty"`

	// Specifies the supported Azure Region where the Load Balancer should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The SKU of the Azure Load Balancer. Accepted values are Basic, Standard and Gateway. Defaults to Basic. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// sku_tier -  The SKU tier of this Load Balancer. Possible values are Global and Regional. Defaults to Regional. Changing this forces a new resource to be created.
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LoadBalancerObservation struct {

	// Specifies the Edge Zone within the Azure Region where this Load Balancer should exist. Changing this forces a new Load Balancer to be created.
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// One or multiple frontend_ip_configuration blocks as documented below.
	FrontendIPConfiguration []LoadBalancerFrontendIPConfigurationObservation `json:"frontendIpConfiguration,omitempty" tf:"frontend_ip_configuration,omitempty"`

	// The Load Balancer ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure Region where the Load Balancer should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The first private IP address assigned to the load balancer in frontend_ip_configuration blocks, if any.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The list of private IP address assigned to the load balancer in frontend_ip_configuration blocks, if any.
	PrivateIPAddresses []*string `json:"privateIpAddresses,omitempty" tf:"private_ip_addresses,omitempty"`

	// The name of the Resource Group in which to create the Load Balancer. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The SKU of the Azure Load Balancer. Accepted values are Basic, Standard and Gateway. Defaults to Basic. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// sku_tier -  The SKU tier of this Load Balancer. Possible values are Global and Regional. Defaults to Regional. Changing this forces a new resource to be created.
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LoadBalancerParameters struct {

	// Specifies the Edge Zone within the Azure Region where this Load Balancer should exist. Changing this forces a new Load Balancer to be created.
	// +kubebuilder:validation:Optional
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// One or multiple frontend_ip_configuration blocks as documented below.
	// +kubebuilder:validation:Optional
	FrontendIPConfiguration []LoadBalancerFrontendIPConfigurationParameters `json:"frontendIpConfiguration,omitempty" tf:"frontend_ip_configuration,omitempty"`

	// Specifies the supported Azure Region where the Load Balancer should be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group in which to create the Load Balancer. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU of the Azure Load Balancer. Accepted values are Basic, Standard and Gateway. Defaults to Basic. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// sku_tier -  The SKU tier of this Load Balancer. Possible values are Global and Regional. Defaults to Regional. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LoadBalancerSpec defines the desired state of LoadBalancer
type LoadBalancerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerParameters `json:"forProvider"`
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
	InitProvider LoadBalancerInitParameters `json:"initProvider,omitempty"`
}

// LoadBalancerStatus defines the observed state of LoadBalancer.
type LoadBalancerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancer is the Schema for the LoadBalancers API. Manages a Load Balancer Resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   LoadBalancerSpec   `json:"spec"`
	Status LoadBalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerList contains a list of LoadBalancers
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancer_Kind             = "LoadBalancer"
	LoadBalancer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancer_Kind}.String()
	LoadBalancer_KindAPIVersion   = LoadBalancer_Kind + "." + CRDGroupVersion.String()
	LoadBalancer_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancer_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
