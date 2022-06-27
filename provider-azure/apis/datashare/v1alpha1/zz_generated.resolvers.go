/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DataSetBlobStorage.
func (mg *DataSetBlobStorage) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageAccount); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccount[i3].ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.StorageAccount[i3].ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.StorageAccount[i3].ResourceGroupNameSelector,
			To: reference.To{
				List:    &v1beta1.ResourceGroupList{},
				Managed: &v1beta1.ResourceGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccount[i3].ResourceGroupName")
		}
		mg.Spec.ForProvider.StorageAccount[i3].ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageAccount[i3].ResourceGroupNameRef = rsp.ResolvedReference

	}

	return nil
}
