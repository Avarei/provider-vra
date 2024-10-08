//go:build !ignore_autogenerated

/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceEntitlement) DeepCopyInto(out *CatalogSourceEntitlement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceEntitlement.
func (in *CatalogSourceEntitlement) DeepCopy() *CatalogSourceEntitlement {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceEntitlement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogSourceEntitlement) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceEntitlementInitParameters) DeepCopyInto(out *CatalogSourceEntitlementInitParameters) {
	*out = *in
	if in.CatalogSourceID != nil {
		in, out := &in.CatalogSourceID, &out.CatalogSourceID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ProjectIDRef != nil {
		in, out := &in.ProjectIDRef, &out.ProjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectIDSelector != nil {
		in, out := &in.ProjectIDSelector, &out.ProjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceEntitlementInitParameters.
func (in *CatalogSourceEntitlementInitParameters) DeepCopy() *CatalogSourceEntitlementInitParameters {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceEntitlementInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceEntitlementList) DeepCopyInto(out *CatalogSourceEntitlementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CatalogSourceEntitlement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceEntitlementList.
func (in *CatalogSourceEntitlementList) DeepCopy() *CatalogSourceEntitlementList {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceEntitlementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogSourceEntitlementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceEntitlementObservation) DeepCopyInto(out *CatalogSourceEntitlementObservation) {
	*out = *in
	if in.CatalogSourceID != nil {
		in, out := &in.CatalogSourceID, &out.CatalogSourceID
		*out = new(string)
		**out = **in
	}
	if in.Definition != nil {
		in, out := &in.Definition, &out.Definition
		*out = make([]DefinitionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceEntitlementObservation.
func (in *CatalogSourceEntitlementObservation) DeepCopy() *CatalogSourceEntitlementObservation {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceEntitlementObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceEntitlementParameters) DeepCopyInto(out *CatalogSourceEntitlementParameters) {
	*out = *in
	if in.CatalogSourceID != nil {
		in, out := &in.CatalogSourceID, &out.CatalogSourceID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ProjectIDRef != nil {
		in, out := &in.ProjectIDRef, &out.ProjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectIDSelector != nil {
		in, out := &in.ProjectIDSelector, &out.ProjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceEntitlementParameters.
func (in *CatalogSourceEntitlementParameters) DeepCopy() *CatalogSourceEntitlementParameters {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceEntitlementParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceEntitlementSpec) DeepCopyInto(out *CatalogSourceEntitlementSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceEntitlementSpec.
func (in *CatalogSourceEntitlementSpec) DeepCopy() *CatalogSourceEntitlementSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceEntitlementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceEntitlementStatus) DeepCopyInto(out *CatalogSourceEntitlementStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceEntitlementStatus.
func (in *CatalogSourceEntitlementStatus) DeepCopy() *CatalogSourceEntitlementStatus {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceEntitlementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionInitParameters) DeepCopyInto(out *DefinitionInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionInitParameters.
func (in *DefinitionInitParameters) DeepCopy() *DefinitionInitParameters {
	if in == nil {
		return nil
	}
	out := new(DefinitionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionObservation) DeepCopyInto(out *DefinitionObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IconID != nil {
		in, out := &in.IconID, &out.IconID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NumberOfItems != nil {
		in, out := &in.NumberOfItems, &out.NumberOfItems
		*out = new(float64)
		**out = **in
	}
	if in.SourceName != nil {
		in, out := &in.SourceName, &out.SourceName
		*out = new(string)
		**out = **in
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionObservation.
func (in *DefinitionObservation) DeepCopy() *DefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(DefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionParameters) DeepCopyInto(out *DefinitionParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionParameters.
func (in *DefinitionParameters) DeepCopy() *DefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(DefinitionParameters)
	in.DeepCopyInto(out)
	return out
}
