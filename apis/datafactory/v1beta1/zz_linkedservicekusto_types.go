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

type LinkedServiceKustoInitParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The service principal id in which to authenticate against the Kusto Database.
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// The service principal tenant id or name in which to authenticate against the Kusto Database.
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`

	// Whether to use the Data Factory's managed identity to authenticate against the Kusto Database.
	UseManagedIdentity *bool `json:"useManagedIdentity,omitempty" tf:"use_managed_identity,omitempty"`
}

type LinkedServiceKustoObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// The Kusto Database Name.
	KustoDatabaseName *string `json:"kustoDatabaseName,omitempty" tf:"kusto_database_name,omitempty"`

	// The URI of the Kusto Cluster endpoint.
	KustoEndpoint *string `json:"kustoEndpoint,omitempty" tf:"kusto_endpoint,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The service principal id in which to authenticate against the Kusto Database.
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// The service principal tenant id or name in which to authenticate against the Kusto Database.
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`

	// Whether to use the Data Factory's managed identity to authenticate against the Kusto Database.
	UseManagedIdentity *bool `json:"useManagedIdentity,omitempty" tf:"use_managed_identity,omitempty"`
}

type LinkedServiceKustoParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// The Kusto Database Name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Database
	// +kubebuilder:validation:Optional
	KustoDatabaseName *string `json:"kustoDatabaseName,omitempty" tf:"kusto_database_name,omitempty"`

	// Reference to a Database in kusto to populate kustoDatabaseName.
	// +kubebuilder:validation:Optional
	KustoDatabaseNameRef *v1.Reference `json:"kustoDatabaseNameRef,omitempty" tf:"-"`

	// Selector for a Database in kusto to populate kustoDatabaseName.
	// +kubebuilder:validation:Optional
	KustoDatabaseNameSelector *v1.Selector `json:"kustoDatabaseNameSelector,omitempty" tf:"-"`

	// The URI of the Kusto Cluster endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uri",true)
	// +kubebuilder:validation:Optional
	KustoEndpoint *string `json:"kustoEndpoint,omitempty" tf:"kusto_endpoint,omitempty"`

	// Reference to a Cluster in kusto to populate kustoEndpoint.
	// +kubebuilder:validation:Optional
	KustoEndpointRef *v1.Reference `json:"kustoEndpointRef,omitempty" tf:"-"`

	// Selector for a Cluster in kusto to populate kustoEndpoint.
	// +kubebuilder:validation:Optional
	KustoEndpointSelector *v1.Selector `json:"kustoEndpointSelector,omitempty" tf:"-"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The service principal id in which to authenticate against the Kusto Database.
	// +kubebuilder:validation:Optional
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// The service principal key in which to authenticate against the Kusto Database.
	// +kubebuilder:validation:Optional
	ServicePrincipalKeySecretRef *v1.SecretKeySelector `json:"servicePrincipalKeySecretRef,omitempty" tf:"-"`

	// The service principal tenant id or name in which to authenticate against the Kusto Database.
	// +kubebuilder:validation:Optional
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`

	// Whether to use the Data Factory's managed identity to authenticate against the Kusto Database.
	// +kubebuilder:validation:Optional
	UseManagedIdentity *bool `json:"useManagedIdentity,omitempty" tf:"use_managed_identity,omitempty"`
}

// LinkedServiceKustoSpec defines the desired state of LinkedServiceKusto
type LinkedServiceKustoSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceKustoParameters `json:"forProvider"`
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
	InitProvider LinkedServiceKustoInitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceKustoStatus defines the observed state of LinkedServiceKusto.
type LinkedServiceKustoStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceKustoObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceKusto is the Schema for the LinkedServiceKustos API. Manages a Linked Service (connection) between a Kusto Cluster and Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure},path=linkedservicekustoes
type LinkedServiceKusto struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkedServiceKustoSpec   `json:"spec"`
	Status            LinkedServiceKustoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceKustoList contains a list of LinkedServiceKustos
type LinkedServiceKustoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceKusto `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceKusto_Kind             = "LinkedServiceKusto"
	LinkedServiceKusto_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceKusto_Kind}.String()
	LinkedServiceKusto_KindAPIVersion   = LinkedServiceKusto_Kind + "." + CRDGroupVersion.String()
	LinkedServiceKusto_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceKusto_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceKusto{}, &LinkedServiceKustoList{})
}
