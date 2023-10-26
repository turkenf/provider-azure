/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azure/apis/appplatform/v1beta1"
	v1beta11 "github.com/upbound/provider-azure/apis/cosmosdb/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this SpringCloudConnection.
func (mg *SpringCloudConnection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SpringCloudID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SpringCloudIDRef,
		Selector:     mg.Spec.ForProvider.SpringCloudIDSelector,
		To: reference.To{
			List:    &v1beta1.SpringCloudJavaDeploymentList{},
			Managed: &v1beta1.SpringCloudJavaDeployment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SpringCloudID")
	}
	mg.Spec.ForProvider.SpringCloudID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SpringCloudIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetResourceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TargetResourceIDRef,
		Selector:     mg.Spec.ForProvider.TargetResourceIDSelector,
		To: reference.To{
			List:    &v1beta11.SQLDatabaseList{},
			Managed: &v1beta11.SQLDatabase{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetResourceID")
	}
	mg.Spec.ForProvider.TargetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetResourceIDRef = rsp.ResolvedReference

	return nil
}
