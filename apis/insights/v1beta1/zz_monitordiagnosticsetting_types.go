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

type EnabledLogInitParameters struct {

	// The name of a Diagnostic Log Category for this Resource.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The name of a Diagnostic Log Category Group for this Resource.
	CategoryGroup *string `json:"categoryGroup,omitempty" tf:"category_group,omitempty"`

	// A retention_policy block as defined below.
	RetentionPolicy []RetentionPolicyInitParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type EnabledLogObservation struct {

	// The name of a Diagnostic Log Category for this Resource.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The name of a Diagnostic Log Category Group for this Resource.
	CategoryGroup *string `json:"categoryGroup,omitempty" tf:"category_group,omitempty"`

	// A retention_policy block as defined below.
	RetentionPolicy []RetentionPolicyObservation `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type EnabledLogParameters struct {

	// The name of a Diagnostic Log Category for this Resource.
	// +kubebuilder:validation:Optional
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The name of a Diagnostic Log Category Group for this Resource.
	// +kubebuilder:validation:Optional
	CategoryGroup *string `json:"categoryGroup,omitempty" tf:"category_group,omitempty"`

	// A retention_policy block as defined below.
	// +kubebuilder:validation:Optional
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type LogInitParameters struct {

	// The name of a Diagnostic Log Category for this Resource.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The name of a Diagnostic Log Category Group for this Resource.
	CategoryGroup *string `json:"categoryGroup,omitempty" tf:"category_group,omitempty"`

	// Is this Diagnostic Log enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A retention_policy block as defined below.
	RetentionPolicy []LogRetentionPolicyInitParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type LogObservation struct {

	// The name of a Diagnostic Log Category for this Resource.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The name of a Diagnostic Log Category Group for this Resource.
	CategoryGroup *string `json:"categoryGroup,omitempty" tf:"category_group,omitempty"`

	// Is this Diagnostic Log enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A retention_policy block as defined below.
	RetentionPolicy []LogRetentionPolicyObservation `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type LogParameters struct {

	// The name of a Diagnostic Log Category for this Resource.
	// +kubebuilder:validation:Optional
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The name of a Diagnostic Log Category Group for this Resource.
	// +kubebuilder:validation:Optional
	CategoryGroup *string `json:"categoryGroup,omitempty" tf:"category_group,omitempty"`

	// Is this Diagnostic Log enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A retention_policy block as defined below.
	// +kubebuilder:validation:Optional
	RetentionPolicy []LogRetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type LogRetentionPolicyInitParameters struct {

	// The number of days for which this Retention Policy should apply.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Is this Retention Policy enabled?
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LogRetentionPolicyObservation struct {

	// The number of days for which this Retention Policy should apply.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Is this Retention Policy enabled?
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LogRetentionPolicyParameters struct {

	// The number of days for which this Retention Policy should apply.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Is this Retention Policy enabled?
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type MetricInitParameters struct {

	// The name of a Diagnostic Metric Category for this Resource.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// Is this Diagnostic Metric enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A retention_policy block as defined below.
	RetentionPolicy []MetricRetentionPolicyInitParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type MetricObservation struct {

	// The name of a Diagnostic Metric Category for this Resource.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// Is this Diagnostic Metric enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A retention_policy block as defined below.
	RetentionPolicy []MetricRetentionPolicyObservation `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type MetricParameters struct {

	// The name of a Diagnostic Metric Category for this Resource.
	// +kubebuilder:validation:Optional
	Category *string `json:"category" tf:"category,omitempty"`

	// Is this Diagnostic Metric enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A retention_policy block as defined below.
	// +kubebuilder:validation:Optional
	RetentionPolicy []MetricRetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type MetricRetentionPolicyInitParameters struct {

	// The number of days for which this Retention Policy should apply.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Is this Retention Policy enabled?
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type MetricRetentionPolicyObservation struct {

	// The number of days for which this Retention Policy should apply.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Is this Retention Policy enabled?
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type MetricRetentionPolicyParameters struct {

	// The number of days for which this Retention Policy should apply.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Is this Retention Policy enabled?
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type MonitorDiagnosticSettingInitParameters struct {

	// One or more enabled_log blocks as defined below.
	EnabledLog []EnabledLogInitParameters `json:"enabledLog,omitempty" tf:"enabled_log,omitempty"`

	// Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data.
	EventHubAuthorizationRuleID *string `json:"eventhubAuthorizationRuleId,omitempty" tf:"eventhub_authorization_rule_id,omitempty"`

	// Specifies the name of the Event Hub where Diagnostics Data should be sent.
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// One or more log blocks as defined below.
	Log []LogInitParameters `json:"log,omitempty" tf:"log,omitempty"`

	// Possible values are AzureDiagnostics and Dedicated. When set to Dedicated, logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
	LogAnalyticsDestinationType *string `json:"logAnalyticsDestinationType,omitempty" tf:"log_analytics_destination_type,omitempty"`

	// Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent.
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// One or more metric blocks as defined below.
	Metric []MetricInitParameters `json:"metric,omitempty" tf:"metric,omitempty"`

	// Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the market partner solution where Diagnostics Data should be sent. For potential partner integrations, click to learn more about partner integration.
	PartnerSolutionID *string `json:"partnerSolutionId,omitempty" tf:"partner_solution_id,omitempty"`

	// The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`
}

type MonitorDiagnosticSettingObservation struct {

	// One or more enabled_log blocks as defined below.
	EnabledLog []EnabledLogObservation `json:"enabledLog,omitempty" tf:"enabled_log,omitempty"`

	// Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data.
	EventHubAuthorizationRuleID *string `json:"eventhubAuthorizationRuleId,omitempty" tf:"eventhub_authorization_rule_id,omitempty"`

	// Specifies the name of the Event Hub where Diagnostics Data should be sent.
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// The ID of the Diagnostic Setting.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more log blocks as defined below.
	Log []LogObservation `json:"log,omitempty" tf:"log,omitempty"`

	// Possible values are AzureDiagnostics and Dedicated. When set to Dedicated, logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
	LogAnalyticsDestinationType *string `json:"logAnalyticsDestinationType,omitempty" tf:"log_analytics_destination_type,omitempty"`

	// Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent.
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// One or more metric blocks as defined below.
	Metric []MetricObservation `json:"metric,omitempty" tf:"metric,omitempty"`

	// Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the market partner solution where Diagnostics Data should be sent. For potential partner integrations, click to learn more about partner integration.
	PartnerSolutionID *string `json:"partnerSolutionId,omitempty" tf:"partner_solution_id,omitempty"`

	// The ID of the Storage Account where logs should be sent.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`
}

type MonitorDiagnosticSettingParameters struct {

	// One or more enabled_log blocks as defined below.
	// +kubebuilder:validation:Optional
	EnabledLog []EnabledLogParameters `json:"enabledLog,omitempty" tf:"enabled_log,omitempty"`

	// Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data.
	// +kubebuilder:validation:Optional
	EventHubAuthorizationRuleID *string `json:"eventhubAuthorizationRuleId,omitempty" tf:"eventhub_authorization_rule_id,omitempty"`

	// Specifies the name of the Event Hub where Diagnostics Data should be sent.
	// +kubebuilder:validation:Optional
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// One or more log blocks as defined below.
	// +kubebuilder:validation:Optional
	Log []LogParameters `json:"log,omitempty" tf:"log,omitempty"`

	// Possible values are AzureDiagnostics and Dedicated. When set to Dedicated, logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
	// +kubebuilder:validation:Optional
	LogAnalyticsDestinationType *string `json:"logAnalyticsDestinationType,omitempty" tf:"log_analytics_destination_type,omitempty"`

	// Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// One or more metric blocks as defined below.
	// +kubebuilder:validation:Optional
	Metric []MetricParameters `json:"metric,omitempty" tf:"metric,omitempty"`

	// Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the market partner solution where Diagnostics Data should be sent. For potential partner integrations, click to learn more about partner integration.
	// +kubebuilder:validation:Optional
	PartnerSolutionID *string `json:"partnerSolutionId,omitempty" tf:"partner_solution_id,omitempty"`

	// The ID of the Storage Account where logs should be sent.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`

	// The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`
}

type RetentionPolicyInitParameters struct {

	// The number of days for which this Retention Policy should apply.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Is this Retention Policy enabled?
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type RetentionPolicyObservation struct {

	// The number of days for which this Retention Policy should apply.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Is this Retention Policy enabled?
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type RetentionPolicyParameters struct {

	// The number of days for which this Retention Policy should apply.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Is this Retention Policy enabled?
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

// MonitorDiagnosticSettingSpec defines the desired state of MonitorDiagnosticSetting
type MonitorDiagnosticSettingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorDiagnosticSettingParameters `json:"forProvider"`
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
	InitProvider MonitorDiagnosticSettingInitParameters `json:"initProvider,omitempty"`
}

// MonitorDiagnosticSettingStatus defines the observed state of MonitorDiagnosticSetting.
type MonitorDiagnosticSettingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorDiagnosticSettingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorDiagnosticSetting is the Schema for the MonitorDiagnosticSettings API. Manages a Diagnostic Setting for an existing Resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorDiagnosticSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetResourceId) || (has(self.initProvider) && has(self.initProvider.targetResourceId))",message="spec.forProvider.targetResourceId is a required parameter"
	Spec   MonitorDiagnosticSettingSpec   `json:"spec"`
	Status MonitorDiagnosticSettingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorDiagnosticSettingList contains a list of MonitorDiagnosticSettings
type MonitorDiagnosticSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorDiagnosticSetting `json:"items"`
}

// Repository type metadata.
var (
	MonitorDiagnosticSetting_Kind             = "MonitorDiagnosticSetting"
	MonitorDiagnosticSetting_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorDiagnosticSetting_Kind}.String()
	MonitorDiagnosticSetting_KindAPIVersion   = MonitorDiagnosticSetting_Kind + "." + CRDGroupVersion.String()
	MonitorDiagnosticSetting_GroupVersionKind = CRDGroupVersion.WithKind(MonitorDiagnosticSetting_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorDiagnosticSetting{}, &MonitorDiagnosticSettingList{})
}
