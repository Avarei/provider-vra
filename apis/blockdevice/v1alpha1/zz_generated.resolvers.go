/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/avarei/provider-vra/cmd/provider/apis/project/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BlockDevice.
func (mg *BlockDevice) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	return nil
}

// ResolveReferences of this BlockDeviceSnapshot.
func (mg *BlockDeviceSnapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BlockDeviceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BlockDeviceIDRef,
		Selector:     mg.Spec.ForProvider.BlockDeviceIDSelector,
		To: reference.To{
			List:    &BlockDeviceList{},
			Managed: &BlockDevice{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BlockDeviceID")
	}
	mg.Spec.ForProvider.BlockDeviceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BlockDeviceIDRef = rsp.ResolvedReference

	return nil
}
