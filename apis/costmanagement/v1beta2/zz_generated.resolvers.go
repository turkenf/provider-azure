// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *ResourceGroupCostManagementExport) ResolveReferences( // ResolveReferences of this ResourceGroupCostManagementExport.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.ExportDataStorageLocation != nil {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Container", "ContainerList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExportDataStorageLocation.ContainerID),
				Extract:      resource.ExtractParamPath("resource_manager_id", true),
				Reference:    mg.Spec.ForProvider.ExportDataStorageLocation.ContainerIDRef,
				Selector:     mg.Spec.ForProvider.ExportDataStorageLocation.ContainerIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ExportDataStorageLocation.ContainerID")
		}
		mg.Spec.ForProvider.ExportDataStorageLocation.ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ExportDataStorageLocation.ContainerIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ResourceGroupIDRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupID")
	}
	mg.Spec.ForProvider.ResourceGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupIDRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.ExportDataStorageLocation != nil {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Container", "ContainerList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ExportDataStorageLocation.ContainerID),
				Extract:      resource.ExtractParamPath("resource_manager_id", true),
				Reference:    mg.Spec.InitProvider.ExportDataStorageLocation.ContainerIDRef,
				Selector:     mg.Spec.InitProvider.ExportDataStorageLocation.ContainerIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ExportDataStorageLocation.ContainerID")
		}
		mg.Spec.InitProvider.ExportDataStorageLocation.ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ExportDataStorageLocation.ContainerIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceGroupID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ResourceGroupIDRef,
			Selector:     mg.Spec.InitProvider.ResourceGroupIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceGroupID")
	}
	mg.Spec.InitProvider.ResourceGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubscriptionCostManagementExport.
func (mg *SubscriptionCostManagementExport) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.ExportDataStorageLocation != nil {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Container", "ContainerList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExportDataStorageLocation.ContainerID),
				Extract:      resource.ExtractParamPath("resource_manager_id", true),
				Reference:    mg.Spec.ForProvider.ExportDataStorageLocation.ContainerIDRef,
				Selector:     mg.Spec.ForProvider.ExportDataStorageLocation.ContainerIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ExportDataStorageLocation.ContainerID")
		}
		mg.Spec.ForProvider.ExportDataStorageLocation.ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ExportDataStorageLocation.ContainerIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "Subscription", "SubscriptionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubscriptionID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SubscriptionIDRef,
			Selector:     mg.Spec.ForProvider.SubscriptionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubscriptionID")
	}
	mg.Spec.ForProvider.SubscriptionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubscriptionIDRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.ExportDataStorageLocation != nil {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Container", "ContainerList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ExportDataStorageLocation.ContainerID),
				Extract:      resource.ExtractParamPath("resource_manager_id", true),
				Reference:    mg.Spec.InitProvider.ExportDataStorageLocation.ContainerIDRef,
				Selector:     mg.Spec.InitProvider.ExportDataStorageLocation.ContainerIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ExportDataStorageLocation.ContainerID")
		}
		mg.Spec.InitProvider.ExportDataStorageLocation.ContainerID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ExportDataStorageLocation.ContainerIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "Subscription", "SubscriptionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubscriptionID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SubscriptionIDRef,
			Selector:     mg.Spec.InitProvider.SubscriptionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubscriptionID")
	}
	mg.Spec.InitProvider.SubscriptionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubscriptionIDRef = rsp.ResolvedReference

	return nil
}