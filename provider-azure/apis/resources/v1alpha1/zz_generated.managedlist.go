/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import resource "github.com/crossplane/crossplane-runtime/pkg/resource"

// GetItems of this ManagementGroupSubscriptionAssociationList.
func (l *ManagementGroupSubscriptionAssociationList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this ManagementGroupTemplateDeploymentList.
func (l *ManagementGroupTemplateDeploymentList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this SubscriptionTemplateDeploymentList.
func (l *SubscriptionTemplateDeploymentList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}
