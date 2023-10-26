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

type HashInitParameters struct {

	// Specifies the algorithm used for the hash content.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The hash value of the content.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type HashObservation struct {

	// Specifies the algorithm used for the hash content.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The hash value of the content.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type HashParameters struct {

	// Specifies the algorithm used for the hash content.
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

	// The hash value of the content.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type ModuleInitParameters struct {

	// A module_link block as defined below.
	ModuleLink []ModuleLinkInitParameters `json:"moduleLink,omitempty" tf:"module_link,omitempty"`
}

type ModuleLinkInitParameters struct {

	// A hash block as defined below.
	Hash []HashInitParameters `json:"hash,omitempty" tf:"hash,omitempty"`

	// The URI of the module content (zip or nupkg).
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ModuleLinkObservation struct {

	// A hash block as defined below.
	Hash []HashObservation `json:"hash,omitempty" tf:"hash,omitempty"`

	// The URI of the module content (zip or nupkg).
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ModuleLinkParameters struct {

	// A hash block as defined below.
	// +kubebuilder:validation:Optional
	Hash []HashParameters `json:"hash,omitempty" tf:"hash,omitempty"`

	// The URI of the module content (zip or nupkg).
	// +kubebuilder:validation:Optional
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type ModuleObservation struct {

	// The name of the automation account in which the Module is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// The Automation Module ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A module_link block as defined below.
	ModuleLink []ModuleLinkObservation `json:"moduleLink,omitempty" tf:"module_link,omitempty"`

	// The name of the resource group in which the Module is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type ModuleParameters struct {

	// The name of the automation account in which the Module is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/automation/v1beta1.Account
	// +kubebuilder:validation:Optional
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// Reference to a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameRef *v1.Reference `json:"automationAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameSelector *v1.Selector `json:"automationAccountNameSelector,omitempty" tf:"-"`

	// A module_link block as defined below.
	// +kubebuilder:validation:Optional
	ModuleLink []ModuleLinkParameters `json:"moduleLink,omitempty" tf:"module_link,omitempty"`

	// The name of the resource group in which the Module is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// ModuleSpec defines the desired state of Module
type ModuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ModuleParameters `json:"forProvider"`
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
	InitProvider ModuleInitParameters `json:"initProvider,omitempty"`
}

// ModuleStatus defines the observed state of Module.
type ModuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ModuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Module is the Schema for the Modules API. Manages a Automation Module.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Module struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.moduleLink) || (has(self.initProvider) && has(self.initProvider.moduleLink))",message="spec.forProvider.moduleLink is a required parameter"
	Spec   ModuleSpec   `json:"spec"`
	Status ModuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModuleList contains a list of Modules
type ModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Module `json:"items"`
}

// Repository type metadata.
var (
	Module_Kind             = "Module"
	Module_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Module_Kind}.String()
	Module_KindAPIVersion   = Module_Kind + "." + CRDGroupVersion.String()
	Module_GroupVersionKind = CRDGroupVersion.WithKind(Module_Kind)
)

func init() {
	SchemeBuilder.Register(&Module{}, &ModuleList{})
}
