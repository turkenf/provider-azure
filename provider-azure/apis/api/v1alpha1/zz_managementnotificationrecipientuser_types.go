/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ManagementNotificationRecipientUserObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagementNotificationRecipientUserParameters struct {

	// +kubebuilder:validation:Required
	APIManagementID *string `json:"apiManagementId" tf:"api_management_id,omitempty"`

	// +kubebuilder:validation:Required
	NotificationType *string `json:"notificationType" tf:"notification_type,omitempty"`

	// +kubebuilder:validation:Required
	UserID *string `json:"userId" tf:"user_id,omitempty"`
}

// ManagementNotificationRecipientUserSpec defines the desired state of ManagementNotificationRecipientUser
type ManagementNotificationRecipientUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementNotificationRecipientUserParameters `json:"forProvider"`
}

// ManagementNotificationRecipientUserStatus defines the observed state of ManagementNotificationRecipientUser.
type ManagementNotificationRecipientUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementNotificationRecipientUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementNotificationRecipientUser is the Schema for the ManagementNotificationRecipientUsers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagementNotificationRecipientUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementNotificationRecipientUserSpec   `json:"spec"`
	Status            ManagementNotificationRecipientUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementNotificationRecipientUserList contains a list of ManagementNotificationRecipientUsers
type ManagementNotificationRecipientUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementNotificationRecipientUser `json:"items"`
}

// Repository type metadata.
var (
	ManagementNotificationRecipientUser_Kind             = "ManagementNotificationRecipientUser"
	ManagementNotificationRecipientUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagementNotificationRecipientUser_Kind}.String()
	ManagementNotificationRecipientUser_KindAPIVersion   = ManagementNotificationRecipientUser_Kind + "." + CRDGroupVersion.String()
	ManagementNotificationRecipientUser_GroupVersionKind = CRDGroupVersion.WithKind(ManagementNotificationRecipientUser_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagementNotificationRecipientUser{}, &ManagementNotificationRecipientUserList{})
}
