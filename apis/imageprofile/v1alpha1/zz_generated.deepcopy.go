//go:build !ignore_autogenerated

/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConstraintsInitParameters) DeepCopyInto(out *ConstraintsInitParameters) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConstraintsInitParameters.
func (in *ConstraintsInitParameters) DeepCopy() *ConstraintsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConstraintsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConstraintsObservation) DeepCopyInto(out *ConstraintsObservation) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConstraintsObservation.
func (in *ConstraintsObservation) DeepCopy() *ConstraintsObservation {
	if in == nil {
		return nil
	}
	out := new(ConstraintsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConstraintsParameters) DeepCopyInto(out *ConstraintsParameters) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConstraintsParameters.
func (in *ConstraintsParameters) DeepCopy() *ConstraintsParameters {
	if in == nil {
		return nil
	}
	out := new(ConstraintsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageMappingInitParameters) DeepCopyInto(out *ImageMappingInitParameters) {
	*out = *in
	if in.CloudConfig != nil {
		in, out := &in.CloudConfig, &out.CloudConfig
		*out = new(string)
		**out = **in
	}
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]ConstraintsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.ImageName != nil {
		in, out := &in.ImageName, &out.ImageName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageMappingInitParameters.
func (in *ImageMappingInitParameters) DeepCopy() *ImageMappingInitParameters {
	if in == nil {
		return nil
	}
	out := new(ImageMappingInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageMappingObservation) DeepCopyInto(out *ImageMappingObservation) {
	*out = *in
	if in.CloudConfig != nil {
		in, out := &in.CloudConfig, &out.CloudConfig
		*out = new(string)
		**out = **in
	}
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]ConstraintsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExternalID != nil {
		in, out := &in.ExternalID, &out.ExternalID
		*out = new(string)
		**out = **in
	}
	if in.ExternalRegionID != nil {
		in, out := &in.ExternalRegionID, &out.ExternalRegionID
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.ImageName != nil {
		in, out := &in.ImageName, &out.ImageName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Organization != nil {
		in, out := &in.Organization, &out.Organization
		*out = new(string)
		**out = **in
	}
	if in.OsFamily != nil {
		in, out := &in.OsFamily, &out.OsFamily
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.Private != nil {
		in, out := &in.Private, &out.Private
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageMappingObservation.
func (in *ImageMappingObservation) DeepCopy() *ImageMappingObservation {
	if in == nil {
		return nil
	}
	out := new(ImageMappingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageMappingParameters) DeepCopyInto(out *ImageMappingParameters) {
	*out = *in
	if in.CloudConfig != nil {
		in, out := &in.CloudConfig, &out.CloudConfig
		*out = new(string)
		**out = **in
	}
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]ConstraintsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.ImageName != nil {
		in, out := &in.ImageName, &out.ImageName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageMappingParameters.
func (in *ImageMappingParameters) DeepCopy() *ImageMappingParameters {
	if in == nil {
		return nil
	}
	out := new(ImageMappingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Profile) DeepCopyInto(out *Profile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Profile.
func (in *Profile) DeepCopy() *Profile {
	if in == nil {
		return nil
	}
	out := new(Profile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Profile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileInitParameters) DeepCopyInto(out *ProfileInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ImageMapping != nil {
		in, out := &in.ImageMapping, &out.ImageMapping
		*out = make([]ImageMappingInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RegionID != nil {
		in, out := &in.RegionID, &out.RegionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileInitParameters.
func (in *ProfileInitParameters) DeepCopy() *ProfileInitParameters {
	if in == nil {
		return nil
	}
	out := new(ProfileInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileList) DeepCopyInto(out *ProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Profile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileList.
func (in *ProfileList) DeepCopy() *ProfileList {
	if in == nil {
		return nil
	}
	out := new(ProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileObservation) DeepCopyInto(out *ProfileObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExternalRegionID != nil {
		in, out := &in.ExternalRegionID, &out.ExternalRegionID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ImageMapping != nil {
		in, out := &in.ImageMapping, &out.ImageMapping
		*out = make([]ImageMappingObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.RegionID != nil {
		in, out := &in.RegionID, &out.RegionID
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileObservation.
func (in *ProfileObservation) DeepCopy() *ProfileObservation {
	if in == nil {
		return nil
	}
	out := new(ProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileParameters) DeepCopyInto(out *ProfileParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ImageMapping != nil {
		in, out := &in.ImageMapping, &out.ImageMapping
		*out = make([]ImageMappingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RegionID != nil {
		in, out := &in.RegionID, &out.RegionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileParameters.
func (in *ProfileParameters) DeepCopy() *ProfileParameters {
	if in == nil {
		return nil
	}
	out := new(ProfileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileSpec) DeepCopyInto(out *ProfileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileSpec.
func (in *ProfileSpec) DeepCopy() *ProfileSpec {
	if in == nil {
		return nil
	}
	out := new(ProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileStatus) DeepCopyInto(out *ProfileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileStatus.
func (in *ProfileStatus) DeepCopy() *ProfileStatus {
	if in == nil {
		return nil
	}
	out := new(ProfileStatus)
	in.DeepCopyInto(out)
	return out
}
