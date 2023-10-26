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

type SubscriptionPolicyRemediationInitParameters struct {

	// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
	FailurePercentage *float64 `json:"failurePercentage,omitempty" tf:"failure_percentage,omitempty"`

	// A list of the resource locations that will be remediated.
	LocationFilters []*string `json:"locationFilters,omitempty" tf:"location_filters,omitempty"`

	// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
	ParallelDeployments *float64 `json:"parallelDeployments,omitempty" tf:"parallel_deployments,omitempty"`

	// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty" tf:"policy_definition_id,omitempty"`

	// The unique ID for the policy definition reference within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`

	// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
	ResourceCount *float64 `json:"resourceCount,omitempty" tf:"resource_count,omitempty"`

	// The way that resources to remediate are discovered. Possible values are ExistingNonCompliant, ReEvaluateCompliance. Defaults to ExistingNonCompliant.
	ResourceDiscoveryMode *string `json:"resourceDiscoveryMode,omitempty" tf:"resource_discovery_mode,omitempty"`

	// The Subscription ID at which the Policy Remediation should be applied. Changing this forces a new resource to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type SubscriptionPolicyRemediationObservation struct {

	// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
	FailurePercentage *float64 `json:"failurePercentage,omitempty" tf:"failure_percentage,omitempty"`

	// The ID of the Policy Remediation.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of the resource locations that will be remediated.
	LocationFilters []*string `json:"locationFilters,omitempty" tf:"location_filters,omitempty"`

	// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
	ParallelDeployments *float64 `json:"parallelDeployments,omitempty" tf:"parallel_deployments,omitempty"`

	// The ID of the Policy Assignment that should be remediated.
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty" tf:"policy_assignment_id,omitempty"`

	// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty" tf:"policy_definition_id,omitempty"`

	// The unique ID for the policy definition reference within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`

	// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
	ResourceCount *float64 `json:"resourceCount,omitempty" tf:"resource_count,omitempty"`

	// The way that resources to remediate are discovered. Possible values are ExistingNonCompliant, ReEvaluateCompliance. Defaults to ExistingNonCompliant.
	ResourceDiscoveryMode *string `json:"resourceDiscoveryMode,omitempty" tf:"resource_discovery_mode,omitempty"`

	// The Subscription ID at which the Policy Remediation should be applied. Changing this forces a new resource to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type SubscriptionPolicyRemediationParameters struct {

	// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
	// +kubebuilder:validation:Optional
	FailurePercentage *float64 `json:"failurePercentage,omitempty" tf:"failure_percentage,omitempty"`

	// A list of the resource locations that will be remediated.
	// +kubebuilder:validation:Optional
	LocationFilters []*string `json:"locationFilters,omitempty" tf:"location_filters,omitempty"`

	// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
	// +kubebuilder:validation:Optional
	ParallelDeployments *float64 `json:"parallelDeployments,omitempty" tf:"parallel_deployments,omitempty"`

	// The ID of the Policy Assignment that should be remediated.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/authorization/v1beta1.SubscriptionPolicyAssignment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty" tf:"policy_assignment_id,omitempty"`

	// Reference to a SubscriptionPolicyAssignment in authorization to populate policyAssignmentId.
	// +kubebuilder:validation:Optional
	PolicyAssignmentIDRef *v1.Reference `json:"policyAssignmentIdRef,omitempty" tf:"-"`

	// Selector for a SubscriptionPolicyAssignment in authorization to populate policyAssignmentId.
	// +kubebuilder:validation:Optional
	PolicyAssignmentIDSelector *v1.Selector `json:"policyAssignmentIdSelector,omitempty" tf:"-"`

	// The unique ID for the policy definition within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	// +kubebuilder:validation:Optional
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty" tf:"policy_definition_id,omitempty"`

	// The unique ID for the policy definition reference within the policy set definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	// +kubebuilder:validation:Optional
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`

	// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
	// +kubebuilder:validation:Optional
	ResourceCount *float64 `json:"resourceCount,omitempty" tf:"resource_count,omitempty"`

	// The way that resources to remediate are discovered. Possible values are ExistingNonCompliant, ReEvaluateCompliance. Defaults to ExistingNonCompliant.
	// +kubebuilder:validation:Optional
	ResourceDiscoveryMode *string `json:"resourceDiscoveryMode,omitempty" tf:"resource_discovery_mode,omitempty"`

	// The Subscription ID at which the Policy Remediation should be applied. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

// SubscriptionPolicyRemediationSpec defines the desired state of SubscriptionPolicyRemediation
type SubscriptionPolicyRemediationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionPolicyRemediationParameters `json:"forProvider"`
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
	InitProvider SubscriptionPolicyRemediationInitParameters `json:"initProvider,omitempty"`
}

// SubscriptionPolicyRemediationStatus defines the observed state of SubscriptionPolicyRemediation.
type SubscriptionPolicyRemediationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionPolicyRemediationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionPolicyRemediation is the Schema for the SubscriptionPolicyRemediations API. Manages an Azure Subscription Policy Remediation.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SubscriptionPolicyRemediation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subscriptionId) || (has(self.initProvider) && has(self.initProvider.subscriptionId))",message="spec.forProvider.subscriptionId is a required parameter"
	Spec   SubscriptionPolicyRemediationSpec   `json:"spec"`
	Status SubscriptionPolicyRemediationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionPolicyRemediationList contains a list of SubscriptionPolicyRemediations
type SubscriptionPolicyRemediationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubscriptionPolicyRemediation `json:"items"`
}

// Repository type metadata.
var (
	SubscriptionPolicyRemediation_Kind             = "SubscriptionPolicyRemediation"
	SubscriptionPolicyRemediation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubscriptionPolicyRemediation_Kind}.String()
	SubscriptionPolicyRemediation_KindAPIVersion   = SubscriptionPolicyRemediation_Kind + "." + CRDGroupVersion.String()
	SubscriptionPolicyRemediation_GroupVersionKind = CRDGroupVersion.WithKind(SubscriptionPolicyRemediation_Kind)
)

func init() {
	SchemeBuilder.Register(&SubscriptionPolicyRemediation{}, &SubscriptionPolicyRemediationList{})
}
