// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MongoRoleDefinitionInitParameters struct {

	// The resource ID of the Mongo DB. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta2.MongoDatabase
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	CosmosMongoDatabaseID *string `json:"cosmosMongoDatabaseId,omitempty" tf:"cosmos_mongo_database_id,omitempty"`

	// Reference to a MongoDatabase in cosmosdb to populate cosmosMongoDatabaseId.
	// +kubebuilder:validation:Optional
	CosmosMongoDatabaseIDRef *v1.Reference `json:"cosmosMongoDatabaseIdRef,omitempty" tf:"-"`

	// Selector for a MongoDatabase in cosmosdb to populate cosmosMongoDatabaseId.
	// +kubebuilder:validation:Optional
	CosmosMongoDatabaseIDSelector *v1.Selector `json:"cosmosMongoDatabaseIdSelector,omitempty" tf:"-"`

	// A list of Mongo Roles which are inherited to the Mongo Role Definition.
	InheritedRoleNames []*string `json:"inheritedRoleNames,omitempty" tf:"inherited_role_names,omitempty"`

	// A privilege block as defined below.
	Privilege []PrivilegeInitParameters `json:"privilege,omitempty" tf:"privilege,omitempty"`
}

type MongoRoleDefinitionObservation struct {

	// The resource ID of the Mongo DB. Changing this forces a new resource to be created.
	CosmosMongoDatabaseID *string `json:"cosmosMongoDatabaseId,omitempty" tf:"cosmos_mongo_database_id,omitempty"`

	// The ID of the Cosmos DB Mongo Role Definition.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of Mongo Roles which are inherited to the Mongo Role Definition.
	InheritedRoleNames []*string `json:"inheritedRoleNames,omitempty" tf:"inherited_role_names,omitempty"`

	// A privilege block as defined below.
	Privilege []PrivilegeObservation `json:"privilege,omitempty" tf:"privilege,omitempty"`
}

type MongoRoleDefinitionParameters struct {

	// The resource ID of the Mongo DB. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta2.MongoDatabase
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CosmosMongoDatabaseID *string `json:"cosmosMongoDatabaseId,omitempty" tf:"cosmos_mongo_database_id,omitempty"`

	// Reference to a MongoDatabase in cosmosdb to populate cosmosMongoDatabaseId.
	// +kubebuilder:validation:Optional
	CosmosMongoDatabaseIDRef *v1.Reference `json:"cosmosMongoDatabaseIdRef,omitempty" tf:"-"`

	// Selector for a MongoDatabase in cosmosdb to populate cosmosMongoDatabaseId.
	// +kubebuilder:validation:Optional
	CosmosMongoDatabaseIDSelector *v1.Selector `json:"cosmosMongoDatabaseIdSelector,omitempty" tf:"-"`

	// A list of Mongo Roles which are inherited to the Mongo Role Definition.
	// +kubebuilder:validation:Optional
	InheritedRoleNames []*string `json:"inheritedRoleNames,omitempty" tf:"inherited_role_names,omitempty"`

	// A privilege block as defined below.
	// +kubebuilder:validation:Optional
	Privilege []PrivilegeParameters `json:"privilege,omitempty" tf:"privilege,omitempty"`
}

type PrivilegeInitParameters struct {

	// A list of actions that are allowed.
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// A resource block as defined below.
	Resource *ResourceInitParameters `json:"resource,omitempty" tf:"resource,omitempty"`
}

type PrivilegeObservation struct {

	// A list of actions that are allowed.
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// A resource block as defined below.
	Resource *ResourceObservation `json:"resource,omitempty" tf:"resource,omitempty"`
}

type PrivilegeParameters struct {

	// A list of actions that are allowed.
	// +kubebuilder:validation:Optional
	Actions []*string `json:"actions" tf:"actions,omitempty"`

	// A resource block as defined below.
	// +kubebuilder:validation:Optional
	Resource *ResourceParameters `json:"resource" tf:"resource,omitempty"`
}

type ResourceInitParameters struct {

	// The name of the Mongo DB Collection that the Role Definition is applied.
	CollectionName *string `json:"collectionName,omitempty" tf:"collection_name,omitempty"`

	// The name of the Mongo DB that the Role Definition is applied.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`
}

type ResourceObservation struct {

	// The name of the Mongo DB Collection that the Role Definition is applied.
	CollectionName *string `json:"collectionName,omitempty" tf:"collection_name,omitempty"`

	// The name of the Mongo DB that the Role Definition is applied.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`
}

type ResourceParameters struct {

	// The name of the Mongo DB Collection that the Role Definition is applied.
	// +kubebuilder:validation:Optional
	CollectionName *string `json:"collectionName,omitempty" tf:"collection_name,omitempty"`

	// The name of the Mongo DB that the Role Definition is applied.
	// +kubebuilder:validation:Optional
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`
}

// MongoRoleDefinitionSpec defines the desired state of MongoRoleDefinition
type MongoRoleDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MongoRoleDefinitionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MongoRoleDefinitionInitParameters `json:"initProvider,omitempty"`
}

// MongoRoleDefinitionStatus defines the observed state of MongoRoleDefinition.
type MongoRoleDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MongoRoleDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MongoRoleDefinition is the Schema for the MongoRoleDefinitions API. Manages a Cosmos DB Mongo Role Definition.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MongoRoleDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MongoRoleDefinitionSpec   `json:"spec"`
	Status            MongoRoleDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MongoRoleDefinitionList contains a list of MongoRoleDefinitions
type MongoRoleDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoRoleDefinition `json:"items"`
}

// Repository type metadata.
var (
	MongoRoleDefinition_Kind             = "MongoRoleDefinition"
	MongoRoleDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MongoRoleDefinition_Kind}.String()
	MongoRoleDefinition_KindAPIVersion   = MongoRoleDefinition_Kind + "." + CRDGroupVersion.String()
	MongoRoleDefinition_GroupVersionKind = CRDGroupVersion.WithKind(MongoRoleDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&MongoRoleDefinition{}, &MongoRoleDefinitionList{})
}
