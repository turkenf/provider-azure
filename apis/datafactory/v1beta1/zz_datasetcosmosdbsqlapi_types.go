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

type DataSetCosmosDBSQLAPIInitParameters struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The collection name of the Data Factory Dataset Azure Cosmos DB SQL API.
	CollectionName *string `json:"collectionName,omitempty" tf:"collection_name,omitempty"`

	// The description for the Data Factory Dataset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A schema_column block as defined below.
	SchemaColumn []DataSetCosmosDBSQLAPISchemaColumnInitParameters `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetCosmosDBSQLAPIObservation struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The collection name of the Data Factory Dataset Azure Cosmos DB SQL API.
	CollectionName *string `json:"collectionName,omitempty" tf:"collection_name,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Dataset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The ID of the Data Factory Dataset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A schema_column block as defined below.
	SchemaColumn []DataSetCosmosDBSQLAPISchemaColumnObservation `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetCosmosDBSQLAPIParameters struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The collection name of the Data Factory Dataset Azure Cosmos DB SQL API.
	// +kubebuilder:validation:Optional
	CollectionName *string `json:"collectionName,omitempty" tf:"collection_name,omitempty"`

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

	// The description for the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The Data Factory Linked Service name in which to associate the Dataset with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceCosmosDB
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceCosmosDB in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceCosmosDB in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// A map of parameters to associate with the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A schema_column block as defined below.
	// +kubebuilder:validation:Optional
	SchemaColumn []DataSetCosmosDBSQLAPISchemaColumnParameters `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetCosmosDBSQLAPISchemaColumnInitParameters struct {

	// The description of the column.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataSetCosmosDBSQLAPISchemaColumnObservation struct {

	// The description of the column.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataSetCosmosDBSQLAPISchemaColumnParameters struct {

	// The description of the column.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DataSetCosmosDBSQLAPISpec defines the desired state of DataSetCosmosDBSQLAPI
type DataSetCosmosDBSQLAPISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetCosmosDBSQLAPIParameters `json:"forProvider"`
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
	InitProvider DataSetCosmosDBSQLAPIInitParameters `json:"initProvider,omitempty"`
}

// DataSetCosmosDBSQLAPIStatus defines the observed state of DataSetCosmosDBSQLAPI.
type DataSetCosmosDBSQLAPIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetCosmosDBSQLAPIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetCosmosDBSQLAPI is the Schema for the DataSetCosmosDBSQLAPIs API. Manages an Azure Cosmos DB SQL API Dataset inside an Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataSetCosmosDBSQLAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetCosmosDBSQLAPISpec   `json:"spec"`
	Status            DataSetCosmosDBSQLAPIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetCosmosDBSQLAPIList contains a list of DataSetCosmosDBSQLAPIs
type DataSetCosmosDBSQLAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetCosmosDBSQLAPI `json:"items"`
}

// Repository type metadata.
var (
	DataSetCosmosDBSQLAPI_Kind             = "DataSetCosmosDBSQLAPI"
	DataSetCosmosDBSQLAPI_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetCosmosDBSQLAPI_Kind}.String()
	DataSetCosmosDBSQLAPI_KindAPIVersion   = DataSetCosmosDBSQLAPI_Kind + "." + CRDGroupVersion.String()
	DataSetCosmosDBSQLAPI_GroupVersionKind = CRDGroupVersion.WithKind(DataSetCosmosDBSQLAPI_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetCosmosDBSQLAPI{}, &DataSetCosmosDBSQLAPIList{})
}
