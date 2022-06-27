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

type IdentityProviderFacebookObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IdentityProviderFacebookParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Required
	AppID *string `json:"appId" tf:"app_id,omitempty"`

	// +kubebuilder:validation:Required
	AppSecretSecretRef v1.SecretKeySelector `json:"appSecretSecretRef" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// IdentityProviderFacebookSpec defines the desired state of IdentityProviderFacebook
type IdentityProviderFacebookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityProviderFacebookParameters `json:"forProvider"`
}

// IdentityProviderFacebookStatus defines the observed state of IdentityProviderFacebook.
type IdentityProviderFacebookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityProviderFacebookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderFacebook is the Schema for the IdentityProviderFacebooks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IdentityProviderFacebook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IdentityProviderFacebookSpec   `json:"spec"`
	Status            IdentityProviderFacebookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderFacebookList contains a list of IdentityProviderFacebooks
type IdentityProviderFacebookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProviderFacebook `json:"items"`
}

// Repository type metadata.
var (
	IdentityProviderFacebook_Kind             = "IdentityProviderFacebook"
	IdentityProviderFacebook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityProviderFacebook_Kind}.String()
	IdentityProviderFacebook_KindAPIVersion   = IdentityProviderFacebook_Kind + "." + CRDGroupVersion.String()
	IdentityProviderFacebook_GroupVersionKind = CRDGroupVersion.WithKind(IdentityProviderFacebook_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityProviderFacebook{}, &IdentityProviderFacebookList{})
}
