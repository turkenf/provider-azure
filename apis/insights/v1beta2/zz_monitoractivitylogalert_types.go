// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActionInitParameters struct {

	// The ID of the Action Group can be sourced from the .
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta2.MonitorActionGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ActionGroupID *string `json:"actionGroupId,omitempty" tf:"action_group_id,omitempty"`

	// Reference to a MonitorActionGroup in insights to populate actionGroupId.
	// +kubebuilder:validation:Optional
	ActionGroupIDRef *v1.Reference `json:"actionGroupIdRef,omitempty" tf:"-"`

	// Selector for a MonitorActionGroup in insights to populate actionGroupId.
	// +kubebuilder:validation:Optional
	ActionGroupIDSelector *v1.Selector `json:"actionGroupIdSelector,omitempty" tf:"-"`

	// The map of custom string properties to include with the post operation. These data are appended to the webhook payload.
	// +mapType=granular
	WebhookProperties map[string]*string `json:"webhookProperties,omitempty" tf:"webhook_properties,omitempty"`
}

type ActionObservation struct {

	// The ID of the Action Group can be sourced from the .
	ActionGroupID *string `json:"actionGroupId,omitempty" tf:"action_group_id,omitempty"`

	// The map of custom string properties to include with the post operation. These data are appended to the webhook payload.
	// +mapType=granular
	WebhookProperties map[string]*string `json:"webhookProperties,omitempty" tf:"webhook_properties,omitempty"`
}

type ActionParameters struct {

	// The ID of the Action Group can be sourced from the .
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta2.MonitorActionGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ActionGroupID *string `json:"actionGroupId,omitempty" tf:"action_group_id,omitempty"`

	// Reference to a MonitorActionGroup in insights to populate actionGroupId.
	// +kubebuilder:validation:Optional
	ActionGroupIDRef *v1.Reference `json:"actionGroupIdRef,omitempty" tf:"-"`

	// Selector for a MonitorActionGroup in insights to populate actionGroupId.
	// +kubebuilder:validation:Optional
	ActionGroupIDSelector *v1.Selector `json:"actionGroupIdSelector,omitempty" tf:"-"`

	// The map of custom string properties to include with the post operation. These data are appended to the webhook payload.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	WebhookProperties map[string]*string `json:"webhookProperties,omitempty" tf:"webhook_properties,omitempty"`
}

type CriteriaInitParameters struct {

	// The email address or Azure Active Directory identifier of the user who performed the operation.
	Caller *string `json:"caller,omitempty" tf:"caller,omitempty"`

	// The category of the operation. Possible values are Administrative, Autoscale, Policy, Recommendation, ResourceHealth, Security and ServiceHealth.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The severity level of the event. Possible values are Verbose, Informational, Warning, Error, and Critical.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// A list of severity level of the event. Possible values are Verbose, Informational, Warning, Error, and Critical.
	Levels []*string `json:"levels,omitempty" tf:"levels,omitempty"`

	// The Resource Manager Role-Based Access Control operation name. Supported operation should be of the form: <resourceProvider>/<resourceType>/<operation>.
	OperationName *string `json:"operationName,omitempty" tf:"operation_name,omitempty"`

	// The recommendation category of the event. Possible values are Cost, Reliability, OperationalExcellence, HighAvailability and Performance. It is only allowed when category is Recommendation.
	RecommendationCategory *string `json:"recommendationCategory,omitempty" tf:"recommendation_category,omitempty"`

	// The recommendation impact of the event. Possible values are High, Medium and Low. It is only allowed when category is Recommendation.
	RecommendationImpact *string `json:"recommendationImpact,omitempty" tf:"recommendation_impact,omitempty"`

	// The recommendation type of the event. It is only allowed when category is Recommendation.
	RecommendationType *string `json:"recommendationType,omitempty" tf:"recommendation_type,omitempty"`

	// The name of resource group monitored by the activity log alert.
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`

	// A list of names of resource groups monitored by the activity log alert.
	ResourceGroups []*string `json:"resourceGroups,omitempty" tf:"resource_groups,omitempty"`

	// A block to define fine grain resource health settings.
	ResourceHealth *ResourceHealthInitParameters `json:"resourceHealth,omitempty" tf:"resource_health,omitempty"`

	// The specific resource monitored by the activity log alert. It should be within one of the scopes.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Account in storage to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// A list of specific resources monitored by the activity log alert. It should be within one of the scopes.
	ResourceIds []*string `json:"resourceIds,omitempty" tf:"resource_ids,omitempty"`

	// The name of the resource provider monitored by the activity log alert.
	ResourceProvider *string `json:"resourceProvider,omitempty" tf:"resource_provider,omitempty"`

	// A list of names of resource providers monitored by the activity log alert.
	ResourceProviders []*string `json:"resourceProviders,omitempty" tf:"resource_providers,omitempty"`

	// The resource type monitored by the activity log alert.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// A list of resource types monitored by the activity log alert.
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`

	// A block to define fine grain service health settings.
	ServiceHealth *ServiceHealthInitParameters `json:"serviceHealth,omitempty" tf:"service_health,omitempty"`

	// The status of the event. For example, Started, Failed, or Succeeded.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A list of status of the event. For example, Started, Failed, or Succeeded.
	Statuses []*string `json:"statuses,omitempty" tf:"statuses,omitempty"`

	// The sub status of the event.
	SubStatus *string `json:"subStatus,omitempty" tf:"sub_status,omitempty"`

	// A list of sub status of the event.
	SubStatuses []*string `json:"subStatuses,omitempty" tf:"sub_statuses,omitempty"`
}

type CriteriaObservation struct {

	// The email address or Azure Active Directory identifier of the user who performed the operation.
	Caller *string `json:"caller,omitempty" tf:"caller,omitempty"`

	// The category of the operation. Possible values are Administrative, Autoscale, Policy, Recommendation, ResourceHealth, Security and ServiceHealth.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The severity level of the event. Possible values are Verbose, Informational, Warning, Error, and Critical.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// A list of severity level of the event. Possible values are Verbose, Informational, Warning, Error, and Critical.
	Levels []*string `json:"levels,omitempty" tf:"levels,omitempty"`

	// The Resource Manager Role-Based Access Control operation name. Supported operation should be of the form: <resourceProvider>/<resourceType>/<operation>.
	OperationName *string `json:"operationName,omitempty" tf:"operation_name,omitempty"`

	// The recommendation category of the event. Possible values are Cost, Reliability, OperationalExcellence, HighAvailability and Performance. It is only allowed when category is Recommendation.
	RecommendationCategory *string `json:"recommendationCategory,omitempty" tf:"recommendation_category,omitempty"`

	// The recommendation impact of the event. Possible values are High, Medium and Low. It is only allowed when category is Recommendation.
	RecommendationImpact *string `json:"recommendationImpact,omitempty" tf:"recommendation_impact,omitempty"`

	// The recommendation type of the event. It is only allowed when category is Recommendation.
	RecommendationType *string `json:"recommendationType,omitempty" tf:"recommendation_type,omitempty"`

	// The name of resource group monitored by the activity log alert.
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`

	// A list of names of resource groups monitored by the activity log alert.
	ResourceGroups []*string `json:"resourceGroups,omitempty" tf:"resource_groups,omitempty"`

	// A block to define fine grain resource health settings.
	ResourceHealth *ResourceHealthObservation `json:"resourceHealth,omitempty" tf:"resource_health,omitempty"`

	// The specific resource monitored by the activity log alert. It should be within one of the scopes.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// A list of specific resources monitored by the activity log alert. It should be within one of the scopes.
	ResourceIds []*string `json:"resourceIds,omitempty" tf:"resource_ids,omitempty"`

	// The name of the resource provider monitored by the activity log alert.
	ResourceProvider *string `json:"resourceProvider,omitempty" tf:"resource_provider,omitempty"`

	// A list of names of resource providers monitored by the activity log alert.
	ResourceProviders []*string `json:"resourceProviders,omitempty" tf:"resource_providers,omitempty"`

	// The resource type monitored by the activity log alert.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// A list of resource types monitored by the activity log alert.
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`

	// A block to define fine grain service health settings.
	ServiceHealth *ServiceHealthObservation `json:"serviceHealth,omitempty" tf:"service_health,omitempty"`

	// The status of the event. For example, Started, Failed, or Succeeded.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A list of status of the event. For example, Started, Failed, or Succeeded.
	Statuses []*string `json:"statuses,omitempty" tf:"statuses,omitempty"`

	// The sub status of the event.
	SubStatus *string `json:"subStatus,omitempty" tf:"sub_status,omitempty"`

	// A list of sub status of the event.
	SubStatuses []*string `json:"subStatuses,omitempty" tf:"sub_statuses,omitempty"`
}

type CriteriaParameters struct {

	// The email address or Azure Active Directory identifier of the user who performed the operation.
	// +kubebuilder:validation:Optional
	Caller *string `json:"caller,omitempty" tf:"caller,omitempty"`

	// The category of the operation. Possible values are Administrative, Autoscale, Policy, Recommendation, ResourceHealth, Security and ServiceHealth.
	// +kubebuilder:validation:Optional
	Category *string `json:"category" tf:"category,omitempty"`

	// The severity level of the event. Possible values are Verbose, Informational, Warning, Error, and Critical.
	// +kubebuilder:validation:Optional
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// A list of severity level of the event. Possible values are Verbose, Informational, Warning, Error, and Critical.
	// +kubebuilder:validation:Optional
	Levels []*string `json:"levels,omitempty" tf:"levels,omitempty"`

	// The Resource Manager Role-Based Access Control operation name. Supported operation should be of the form: <resourceProvider>/<resourceType>/<operation>.
	// +kubebuilder:validation:Optional
	OperationName *string `json:"operationName,omitempty" tf:"operation_name,omitempty"`

	// The recommendation category of the event. Possible values are Cost, Reliability, OperationalExcellence, HighAvailability and Performance. It is only allowed when category is Recommendation.
	// +kubebuilder:validation:Optional
	RecommendationCategory *string `json:"recommendationCategory,omitempty" tf:"recommendation_category,omitempty"`

	// The recommendation impact of the event. Possible values are High, Medium and Low. It is only allowed when category is Recommendation.
	// +kubebuilder:validation:Optional
	RecommendationImpact *string `json:"recommendationImpact,omitempty" tf:"recommendation_impact,omitempty"`

	// The recommendation type of the event. It is only allowed when category is Recommendation.
	// +kubebuilder:validation:Optional
	RecommendationType *string `json:"recommendationType,omitempty" tf:"recommendation_type,omitempty"`

	// The name of resource group monitored by the activity log alert.
	// +kubebuilder:validation:Optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`

	// A list of names of resource groups monitored by the activity log alert.
	// +kubebuilder:validation:Optional
	ResourceGroups []*string `json:"resourceGroups,omitempty" tf:"resource_groups,omitempty"`

	// A block to define fine grain resource health settings.
	// +kubebuilder:validation:Optional
	ResourceHealth *ResourceHealthParameters `json:"resourceHealth,omitempty" tf:"resource_health,omitempty"`

	// The specific resource monitored by the activity log alert. It should be within one of the scopes.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Account in storage to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// A list of specific resources monitored by the activity log alert. It should be within one of the scopes.
	// +kubebuilder:validation:Optional
	ResourceIds []*string `json:"resourceIds,omitempty" tf:"resource_ids,omitempty"`

	// The name of the resource provider monitored by the activity log alert.
	// +kubebuilder:validation:Optional
	ResourceProvider *string `json:"resourceProvider,omitempty" tf:"resource_provider,omitempty"`

	// A list of names of resource providers monitored by the activity log alert.
	// +kubebuilder:validation:Optional
	ResourceProviders []*string `json:"resourceProviders,omitempty" tf:"resource_providers,omitempty"`

	// The resource type monitored by the activity log alert.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// A list of resource types monitored by the activity log alert.
	// +kubebuilder:validation:Optional
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`

	// A block to define fine grain service health settings.
	// +kubebuilder:validation:Optional
	ServiceHealth *ServiceHealthParameters `json:"serviceHealth,omitempty" tf:"service_health,omitempty"`

	// The status of the event. For example, Started, Failed, or Succeeded.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A list of status of the event. For example, Started, Failed, or Succeeded.
	// +kubebuilder:validation:Optional
	Statuses []*string `json:"statuses,omitempty" tf:"statuses,omitempty"`

	// The sub status of the event.
	// +kubebuilder:validation:Optional
	SubStatus *string `json:"subStatus,omitempty" tf:"sub_status,omitempty"`

	// A list of sub status of the event.
	// +kubebuilder:validation:Optional
	SubStatuses []*string `json:"subStatuses,omitempty" tf:"sub_statuses,omitempty"`
}

type MonitorActivityLogAlertInitParameters struct {

	// One or more action blocks as defined below.
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// A criteria block as defined below.
	Criteria *CriteriaInitParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// The description of this activity log alert.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should this Activity Log Alert be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the activity log alert. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which to create the activity log alert instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Scope at which the Activity Log should be applied. A list of strings which could be a resource group , or a subscription, or a resource ID (such as a Storage Account).
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +listType=set
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// References to ResourceGroup in azure to populate scopes.
	// +kubebuilder:validation:Optional
	ScopesRefs []v1.Reference `json:"scopesRefs,omitempty" tf:"-"`

	// Selector for a list of ResourceGroup in azure to populate scopes.
	// +kubebuilder:validation:Optional
	ScopesSelector *v1.Selector `json:"scopesSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorActivityLogAlertObservation struct {

	// One or more action blocks as defined below.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// A criteria block as defined below.
	Criteria *CriteriaObservation `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// The description of this activity log alert.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should this Activity Log Alert be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the activity log alert.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the activity log alert. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which to create the activity log alert instance. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Scope at which the Activity Log should be applied. A list of strings which could be a resource group , or a subscription, or a resource ID (such as a Storage Account).
	// +listType=set
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorActivityLogAlertParameters struct {

	// One or more action blocks as defined below.
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// A criteria block as defined below.
	// +kubebuilder:validation:Optional
	Criteria *CriteriaParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// The description of this activity log alert.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should this Activity Log Alert be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the activity log alert. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which to create the activity log alert instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Scope at which the Activity Log should be applied. A list of strings which could be a resource group , or a subscription, or a resource ID (such as a Storage Account).
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	// +listType=set
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// References to ResourceGroup in azure to populate scopes.
	// +kubebuilder:validation:Optional
	ScopesRefs []v1.Reference `json:"scopesRefs,omitempty" tf:"-"`

	// Selector for a list of ResourceGroup in azure to populate scopes.
	// +kubebuilder:validation:Optional
	ScopesSelector *v1.Selector `json:"scopesSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ResourceHealthInitParameters struct {

	// The current resource health statuses that will log an alert. Possible values are Available, Degraded, Unavailable and Unknown.
	// +listType=set
	Current []*string `json:"current,omitempty" tf:"current,omitempty"`

	// The previous resource health statuses that will log an alert. Possible values are Available, Degraded, Unavailable and Unknown.
	// +listType=set
	Previous []*string `json:"previous,omitempty" tf:"previous,omitempty"`

	// The reason that will log an alert. Possible values are PlatformInitiated (such as a problem with the resource in an affected region of an Azure incident), UserInitiated (such as a shutdown request of a VM) and Unknown.
	// +listType=set
	Reason []*string `json:"reason,omitempty" tf:"reason,omitempty"`
}

type ResourceHealthObservation struct {

	// The current resource health statuses that will log an alert. Possible values are Available, Degraded, Unavailable and Unknown.
	// +listType=set
	Current []*string `json:"current,omitempty" tf:"current,omitempty"`

	// The previous resource health statuses that will log an alert. Possible values are Available, Degraded, Unavailable and Unknown.
	// +listType=set
	Previous []*string `json:"previous,omitempty" tf:"previous,omitempty"`

	// The reason that will log an alert. Possible values are PlatformInitiated (such as a problem with the resource in an affected region of an Azure incident), UserInitiated (such as a shutdown request of a VM) and Unknown.
	// +listType=set
	Reason []*string `json:"reason,omitempty" tf:"reason,omitempty"`
}

type ResourceHealthParameters struct {

	// The current resource health statuses that will log an alert. Possible values are Available, Degraded, Unavailable and Unknown.
	// +kubebuilder:validation:Optional
	// +listType=set
	Current []*string `json:"current,omitempty" tf:"current,omitempty"`

	// The previous resource health statuses that will log an alert. Possible values are Available, Degraded, Unavailable and Unknown.
	// +kubebuilder:validation:Optional
	// +listType=set
	Previous []*string `json:"previous,omitempty" tf:"previous,omitempty"`

	// The reason that will log an alert. Possible values are PlatformInitiated (such as a problem with the resource in an affected region of an Azure incident), UserInitiated (such as a shutdown request of a VM) and Unknown.
	// +kubebuilder:validation:Optional
	// +listType=set
	Reason []*string `json:"reason,omitempty" tf:"reason,omitempty"`
}

type ServiceHealthInitParameters struct {

	// Events this alert will monitor Possible values are Incident, Maintenance, Informational, ActionRequired and Security.
	// +listType=set
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// Locations this alert will monitor. For example, West Europe.
	// +listType=set
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`

	// Services this alert will monitor. For example, Activity Logs & Alerts, Action Groups. Defaults to all Services.
	// +listType=set
	Services []*string `json:"services,omitempty" tf:"services,omitempty"`
}

type ServiceHealthObservation struct {

	// Events this alert will monitor Possible values are Incident, Maintenance, Informational, ActionRequired and Security.
	// +listType=set
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// Locations this alert will monitor. For example, West Europe.
	// +listType=set
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`

	// Services this alert will monitor. For example, Activity Logs & Alerts, Action Groups. Defaults to all Services.
	// +listType=set
	Services []*string `json:"services,omitempty" tf:"services,omitempty"`
}

type ServiceHealthParameters struct {

	// Events this alert will monitor Possible values are Incident, Maintenance, Informational, ActionRequired and Security.
	// +kubebuilder:validation:Optional
	// +listType=set
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// Locations this alert will monitor. For example, West Europe.
	// +kubebuilder:validation:Optional
	// +listType=set
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`

	// Services this alert will monitor. For example, Activity Logs & Alerts, Action Groups. Defaults to all Services.
	// +kubebuilder:validation:Optional
	// +listType=set
	Services []*string `json:"services,omitempty" tf:"services,omitempty"`
}

// MonitorActivityLogAlertSpec defines the desired state of MonitorActivityLogAlert
type MonitorActivityLogAlertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorActivityLogAlertParameters `json:"forProvider"`
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
	InitProvider MonitorActivityLogAlertInitParameters `json:"initProvider,omitempty"`
}

// MonitorActivityLogAlertStatus defines the observed state of MonitorActivityLogAlert.
type MonitorActivityLogAlertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorActivityLogAlertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MonitorActivityLogAlert is the Schema for the MonitorActivityLogAlerts API. Manages an Activity Log Alert within Azure Monitor
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorActivityLogAlert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.criteria) || (has(self.initProvider) && has(self.initProvider.criteria))",message="spec.forProvider.criteria is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   MonitorActivityLogAlertSpec   `json:"spec"`
	Status MonitorActivityLogAlertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorActivityLogAlertList contains a list of MonitorActivityLogAlerts
type MonitorActivityLogAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorActivityLogAlert `json:"items"`
}

// Repository type metadata.
var (
	MonitorActivityLogAlert_Kind             = "MonitorActivityLogAlert"
	MonitorActivityLogAlert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorActivityLogAlert_Kind}.String()
	MonitorActivityLogAlert_KindAPIVersion   = MonitorActivityLogAlert_Kind + "." + CRDGroupVersion.String()
	MonitorActivityLogAlert_GroupVersionKind = CRDGroupVersion.WithKind(MonitorActivityLogAlert_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorActivityLogAlert{}, &MonitorActivityLogAlertList{})
}