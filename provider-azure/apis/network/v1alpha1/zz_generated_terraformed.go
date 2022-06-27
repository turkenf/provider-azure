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

// GetTerraformResourceType returns Terraform resource type for this VirtualNetworkGateway
func (mg *VirtualNetworkGateway) GetTerraformResourceType() string {
	return "azurerm_virtual_network_gateway"
}

// GetConnectionDetailsMapping for this VirtualNetworkGateway
func (tr *VirtualNetworkGateway) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this VirtualNetworkGateway
func (tr *VirtualNetworkGateway) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this VirtualNetworkGateway
func (tr *VirtualNetworkGateway) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this VirtualNetworkGateway
func (tr *VirtualNetworkGateway) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this VirtualNetworkGateway
func (tr *VirtualNetworkGateway) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this VirtualNetworkGateway
func (tr *VirtualNetworkGateway) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this VirtualNetworkGateway using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *VirtualNetworkGateway) LateInitialize(attrs []byte) (bool, error) {
	params := &VirtualNetworkGatewayParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *VirtualNetworkGateway) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this VirtualNetworkGatewayConnection
func (mg *VirtualNetworkGatewayConnection) GetTerraformResourceType() string {
	return "azurerm_virtual_network_gateway_connection"
}

// GetConnectionDetailsMapping for this VirtualNetworkGatewayConnection
func (tr *VirtualNetworkGatewayConnection) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"authorization_key": "spec.forProvider.authorizationKeySecretRef", "shared_key": "spec.forProvider.sharedKeySecretRef"}
}

// GetObservation of this VirtualNetworkGatewayConnection
func (tr *VirtualNetworkGatewayConnection) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this VirtualNetworkGatewayConnection
func (tr *VirtualNetworkGatewayConnection) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this VirtualNetworkGatewayConnection
func (tr *VirtualNetworkGatewayConnection) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this VirtualNetworkGatewayConnection
func (tr *VirtualNetworkGatewayConnection) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this VirtualNetworkGatewayConnection
func (tr *VirtualNetworkGatewayConnection) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this VirtualNetworkGatewayConnection using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *VirtualNetworkGatewayConnection) LateInitialize(attrs []byte) (bool, error) {
	params := &VirtualNetworkGatewayConnectionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *VirtualNetworkGatewayConnection) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this VirtualNetworkPeering
func (mg *VirtualNetworkPeering) GetTerraformResourceType() string {
	return "azurerm_virtual_network_peering"
}

// GetConnectionDetailsMapping for this VirtualNetworkPeering
func (tr *VirtualNetworkPeering) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this VirtualNetworkPeering
func (tr *VirtualNetworkPeering) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this VirtualNetworkPeering
func (tr *VirtualNetworkPeering) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this VirtualNetworkPeering
func (tr *VirtualNetworkPeering) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this VirtualNetworkPeering
func (tr *VirtualNetworkPeering) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this VirtualNetworkPeering
func (tr *VirtualNetworkPeering) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this VirtualNetworkPeering using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *VirtualNetworkPeering) LateInitialize(attrs []byte) (bool, error) {
	params := &VirtualNetworkPeeringParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *VirtualNetworkPeering) GetTerraformSchemaVersion() int {
	return 0
}
