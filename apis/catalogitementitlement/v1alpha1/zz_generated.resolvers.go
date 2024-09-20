/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/avarei/provider-vra/apis/project/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProjectIDRef,
		Selector:     mg.Spec.ForProvider.ProjectIDSelector,
		To: reference.To{
			List:    &v1alpha1.ProjectList{},
			Managed: &v1alpha1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProjectID")
	}
	mg.Spec.ForProvider.ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ProjectIDRef,
		Selector:     mg.Spec.InitProvider.ProjectIDSelector,
		To: reference.To{
			List:    &v1alpha1.ProjectList{},
			Managed: &v1alpha1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ProjectID")
	}
	mg.Spec.InitProvider.ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectIDRef = rsp.ResolvedReference

	return nil
}
