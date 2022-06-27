/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Cluster
func (mg *Cluster) GetTerraformResourceType() string {
	return "azurerm_eventhub_cluster"
}

// GetConnectionDetailsMapping for this Cluster
func (tr *Cluster) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Cluster
func (tr *Cluster) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Cluster
func (tr *Cluster) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Cluster
func (tr *Cluster) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Cluster
func (tr *Cluster) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Cluster
func (tr *Cluster) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Cluster using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Cluster) LateInitialize(attrs []byte) (bool, error) {
	params := &ClusterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Cluster) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this NamespaceAuthorizationRule
func (mg *NamespaceAuthorizationRule) GetTerraformResourceType() string {
	return "azurerm_eventhub_namespace_authorization_rule"
}

// GetConnectionDetailsMapping for this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"primary_connection_string": "status.atProvider.primaryConnectionString", "primary_connection_string_alias": "status.atProvider.primaryConnectionStringAlias", "primary_key": "status.atProvider.primaryKey", "secondary_connection_string": "status.atProvider.secondaryConnectionString", "secondary_connection_string_alias": "status.atProvider.secondaryConnectionStringAlias", "secondary_key": "status.atProvider.secondaryKey"}
}

// GetObservation of this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this NamespaceAuthorizationRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NamespaceAuthorizationRule) LateInitialize(attrs []byte) (bool, error) {
	params := &NamespaceAuthorizationRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NamespaceAuthorizationRule) GetTerraformSchemaVersion() int {
	return 2
}

// GetTerraformResourceType returns Terraform resource type for this NamespaceCustomerManagedKey
func (mg *NamespaceCustomerManagedKey) GetTerraformResourceType() string {
	return "azurerm_eventhub_namespace_customer_managed_key"
}

// GetConnectionDetailsMapping for this NamespaceCustomerManagedKey
func (tr *NamespaceCustomerManagedKey) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this NamespaceCustomerManagedKey
func (tr *NamespaceCustomerManagedKey) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NamespaceCustomerManagedKey
func (tr *NamespaceCustomerManagedKey) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NamespaceCustomerManagedKey
func (tr *NamespaceCustomerManagedKey) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NamespaceCustomerManagedKey
func (tr *NamespaceCustomerManagedKey) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NamespaceCustomerManagedKey
func (tr *NamespaceCustomerManagedKey) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this NamespaceCustomerManagedKey using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NamespaceCustomerManagedKey) LateInitialize(attrs []byte) (bool, error) {
	params := &NamespaceCustomerManagedKeyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NamespaceCustomerManagedKey) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this NamespaceDisasterRecoveryConfig
func (mg *NamespaceDisasterRecoveryConfig) GetTerraformResourceType() string {
	return "azurerm_eventhub_namespace_disaster_recovery_config"
}

// GetConnectionDetailsMapping for this NamespaceDisasterRecoveryConfig
func (tr *NamespaceDisasterRecoveryConfig) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this NamespaceDisasterRecoveryConfig
func (tr *NamespaceDisasterRecoveryConfig) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NamespaceDisasterRecoveryConfig
func (tr *NamespaceDisasterRecoveryConfig) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NamespaceDisasterRecoveryConfig
func (tr *NamespaceDisasterRecoveryConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NamespaceDisasterRecoveryConfig
func (tr *NamespaceDisasterRecoveryConfig) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NamespaceDisasterRecoveryConfig
func (tr *NamespaceDisasterRecoveryConfig) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this NamespaceDisasterRecoveryConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NamespaceDisasterRecoveryConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &NamespaceDisasterRecoveryConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NamespaceDisasterRecoveryConfig) GetTerraformSchemaVersion() int {
	return 0
}
