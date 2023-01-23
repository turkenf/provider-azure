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

type CertificateObservation_2 struct {

	// The Expiration Date of this Certificate, formatted as an RFC3339 string.
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The ID of the API Management Certificate.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Subject of this Certificate.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// The Thumbprint of this Certificate.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type CertificateParameters_2 struct {

	// The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The base-64 encoded certificate data, which must be a PFX file.
	// +kubebuilder:validation:Optional
	DataSecretRef *v1.SecretKeySelector `json:"dataSecretRef,omitempty" tf:"-"`

	// The Client ID of the User Assigned Managed Identity to use for retrieving certificate.
	// +kubebuilder:validation:Optional
	KeyVaultIdentityClientID *string `json:"keyVaultIdentityClientId,omitempty" tf:"key_vault_identity_client_id,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be of the type application/x-pkcs12.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("secret_id",true)
	// +kubebuilder:validation:Optional
	KeyVaultSecretID *string `json:"keyVaultSecretId,omitempty" tf:"key_vault_secret_id,omitempty"`

	// Reference to a Certificate in keyvault to populate keyVaultSecretId.
	// +kubebuilder:validation:Optional
	KeyVaultSecretIDRef *v1.Reference `json:"keyVaultSecretIdRef,omitempty" tf:"-"`

	// Selector for a Certificate in keyvault to populate keyVaultSecretId.
	// +kubebuilder:validation:Optional
	KeyVaultSecretIDSelector *v1.Selector `json:"keyVaultSecretIdSelector,omitempty" tf:"-"`

	// The password used for this certificate.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
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

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters_2 `json:"forProvider"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Certificate is the Schema for the Certificates API. Manages an Certificate within an API Management Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec"`
	Status            CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}