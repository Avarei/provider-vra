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
func (in *BlockDevice) DeepCopyInto(out *BlockDevice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDevice.
func (in *BlockDevice) DeepCopy() *BlockDevice {
	if in == nil {
		return nil
	}
	out := new(BlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlockDevice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceInitParameters) DeepCopyInto(out *BlockDeviceInitParameters) {
	*out = *in
	if in.CapacityInGb != nil {
		in, out := &in.CapacityInGb, &out.CapacityInGb
		*out = new(float64)
		**out = **in
	}
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]ConstraintsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomProperties != nil {
		in, out := &in.CustomProperties, &out.CustomProperties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DeploymentID != nil {
		in, out := &in.DeploymentID, &out.DeploymentID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DiskContentBase64 != nil {
		in, out := &in.DiskContentBase64, &out.DiskContentBase64
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.ExpandSnapshots != nil {
		in, out := &in.ExpandSnapshots, &out.ExpandSnapshots
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Persistent != nil {
		in, out := &in.Persistent, &out.Persistent
		*out = new(bool)
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
	if in.Purge != nil {
		in, out := &in.Purge, &out.Purge
		*out = new(bool)
		**out = **in
	}
	if in.SourceReference != nil {
		in, out := &in.SourceReference, &out.SourceReference
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceInitParameters.
func (in *BlockDeviceInitParameters) DeepCopy() *BlockDeviceInitParameters {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceList) DeepCopyInto(out *BlockDeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BlockDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceList.
func (in *BlockDeviceList) DeepCopy() *BlockDeviceList {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlockDeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceObservation) DeepCopyInto(out *BlockDeviceObservation) {
	*out = *in
	if in.CapacityInGb != nil {
		in, out := &in.CapacityInGb, &out.CapacityInGb
		*out = new(float64)
		**out = **in
	}
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]ConstraintsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CustomProperties != nil {
		in, out := &in.CustomProperties, &out.CustomProperties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DeploymentID != nil {
		in, out := &in.DeploymentID, &out.DeploymentID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DiskContentBase64 != nil {
		in, out := &in.DiskContentBase64, &out.DiskContentBase64
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.ExpandSnapshots != nil {
		in, out := &in.ExpandSnapshots, &out.ExpandSnapshots
		*out = new(bool)
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
	if in.ExternalZoneID != nil {
		in, out := &in.ExternalZoneID, &out.ExternalZoneID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]LinksObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.Persistent != nil {
		in, out := &in.Persistent, &out.Persistent
		*out = new(bool)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Purge != nil {
		in, out := &in.Purge, &out.Purge
		*out = new(bool)
		**out = **in
	}
	if in.Snapshots != nil {
		in, out := &in.Snapshots, &out.Snapshots
		*out = make([]SnapshotsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SourceReference != nil {
		in, out := &in.SourceReference, &out.SourceReference
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceObservation.
func (in *BlockDeviceObservation) DeepCopy() *BlockDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceParameters) DeepCopyInto(out *BlockDeviceParameters) {
	*out = *in
	if in.CapacityInGb != nil {
		in, out := &in.CapacityInGb, &out.CapacityInGb
		*out = new(float64)
		**out = **in
	}
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]ConstraintsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomProperties != nil {
		in, out := &in.CustomProperties, &out.CustomProperties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DeploymentID != nil {
		in, out := &in.DeploymentID, &out.DeploymentID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DiskContentBase64 != nil {
		in, out := &in.DiskContentBase64, &out.DiskContentBase64
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.ExpandSnapshots != nil {
		in, out := &in.ExpandSnapshots, &out.ExpandSnapshots
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Persistent != nil {
		in, out := &in.Persistent, &out.Persistent
		*out = new(bool)
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
	if in.Purge != nil {
		in, out := &in.Purge, &out.Purge
		*out = new(bool)
		**out = **in
	}
	if in.SourceReference != nil {
		in, out := &in.SourceReference, &out.SourceReference
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceParameters.
func (in *BlockDeviceParameters) DeepCopy() *BlockDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceSnapshot) DeepCopyInto(out *BlockDeviceSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceSnapshot.
func (in *BlockDeviceSnapshot) DeepCopy() *BlockDeviceSnapshot {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlockDeviceSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceSnapshotInitParameters) DeepCopyInto(out *BlockDeviceSnapshotInitParameters) {
	*out = *in
	if in.BlockDeviceID != nil {
		in, out := &in.BlockDeviceID, &out.BlockDeviceID
		*out = new(string)
		**out = **in
	}
	if in.BlockDeviceIDRef != nil {
		in, out := &in.BlockDeviceIDRef, &out.BlockDeviceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.BlockDeviceIDSelector != nil {
		in, out := &in.BlockDeviceIDSelector, &out.BlockDeviceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceSnapshotInitParameters.
func (in *BlockDeviceSnapshotInitParameters) DeepCopy() *BlockDeviceSnapshotInitParameters {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceSnapshotInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceSnapshotLinksInitParameters) DeepCopyInto(out *BlockDeviceSnapshotLinksInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceSnapshotLinksInitParameters.
func (in *BlockDeviceSnapshotLinksInitParameters) DeepCopy() *BlockDeviceSnapshotLinksInitParameters {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceSnapshotLinksInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceSnapshotLinksObservation) DeepCopyInto(out *BlockDeviceSnapshotLinksObservation) {
	*out = *in
	if in.Href != nil {
		in, out := &in.Href, &out.Href
		*out = new(string)
		**out = **in
	}
	if in.Hrefs != nil {
		in, out := &in.Hrefs, &out.Hrefs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Rel != nil {
		in, out := &in.Rel, &out.Rel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceSnapshotLinksObservation.
func (in *BlockDeviceSnapshotLinksObservation) DeepCopy() *BlockDeviceSnapshotLinksObservation {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceSnapshotLinksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceSnapshotLinksParameters) DeepCopyInto(out *BlockDeviceSnapshotLinksParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceSnapshotLinksParameters.
func (in *BlockDeviceSnapshotLinksParameters) DeepCopy() *BlockDeviceSnapshotLinksParameters {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceSnapshotLinksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceSnapshotList) DeepCopyInto(out *BlockDeviceSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BlockDeviceSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceSnapshotList.
func (in *BlockDeviceSnapshotList) DeepCopy() *BlockDeviceSnapshotList {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlockDeviceSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceSnapshotObservation) DeepCopyInto(out *BlockDeviceSnapshotObservation) {
	*out = *in
	if in.BlockDeviceID != nil {
		in, out := &in.BlockDeviceID, &out.BlockDeviceID
		*out = new(string)
		**out = **in
	}
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
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsCurrent != nil {
		in, out := &in.IsCurrent, &out.IsCurrent
		*out = new(bool)
		**out = **in
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]BlockDeviceSnapshotLinksObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceSnapshotObservation.
func (in *BlockDeviceSnapshotObservation) DeepCopy() *BlockDeviceSnapshotObservation {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceSnapshotObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceSnapshotParameters) DeepCopyInto(out *BlockDeviceSnapshotParameters) {
	*out = *in
	if in.BlockDeviceID != nil {
		in, out := &in.BlockDeviceID, &out.BlockDeviceID
		*out = new(string)
		**out = **in
	}
	if in.BlockDeviceIDRef != nil {
		in, out := &in.BlockDeviceIDRef, &out.BlockDeviceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.BlockDeviceIDSelector != nil {
		in, out := &in.BlockDeviceIDSelector, &out.BlockDeviceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceSnapshotParameters.
func (in *BlockDeviceSnapshotParameters) DeepCopy() *BlockDeviceSnapshotParameters {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceSnapshotParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceSnapshotSpec) DeepCopyInto(out *BlockDeviceSnapshotSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceSnapshotSpec.
func (in *BlockDeviceSnapshotSpec) DeepCopy() *BlockDeviceSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceSnapshotStatus) DeepCopyInto(out *BlockDeviceSnapshotStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceSnapshotStatus.
func (in *BlockDeviceSnapshotStatus) DeepCopy() *BlockDeviceSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceSpec) DeepCopyInto(out *BlockDeviceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceSpec.
func (in *BlockDeviceSpec) DeepCopy() *BlockDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceStatus) DeepCopyInto(out *BlockDeviceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceStatus.
func (in *BlockDeviceStatus) DeepCopy() *BlockDeviceStatus {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceStatus)
	in.DeepCopyInto(out)
	return out
}

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
func (in *LinksInitParameters) DeepCopyInto(out *LinksInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksInitParameters.
func (in *LinksInitParameters) DeepCopy() *LinksInitParameters {
	if in == nil {
		return nil
	}
	out := new(LinksInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinksObservation) DeepCopyInto(out *LinksObservation) {
	*out = *in
	if in.Href != nil {
		in, out := &in.Href, &out.Href
		*out = new(string)
		**out = **in
	}
	if in.Hrefs != nil {
		in, out := &in.Hrefs, &out.Hrefs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Rel != nil {
		in, out := &in.Rel, &out.Rel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksObservation.
func (in *LinksObservation) DeepCopy() *LinksObservation {
	if in == nil {
		return nil
	}
	out := new(LinksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinksParameters) DeepCopyInto(out *LinksParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksParameters.
func (in *LinksParameters) DeepCopy() *LinksParameters {
	if in == nil {
		return nil
	}
	out := new(LinksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotsInitParameters) DeepCopyInto(out *SnapshotsInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotsInitParameters.
func (in *SnapshotsInitParameters) DeepCopy() *SnapshotsInitParameters {
	if in == nil {
		return nil
	}
	out := new(SnapshotsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotsLinksInitParameters) DeepCopyInto(out *SnapshotsLinksInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotsLinksInitParameters.
func (in *SnapshotsLinksInitParameters) DeepCopy() *SnapshotsLinksInitParameters {
	if in == nil {
		return nil
	}
	out := new(SnapshotsLinksInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotsLinksObservation) DeepCopyInto(out *SnapshotsLinksObservation) {
	*out = *in
	if in.Href != nil {
		in, out := &in.Href, &out.Href
		*out = new(string)
		**out = **in
	}
	if in.Hrefs != nil {
		in, out := &in.Hrefs, &out.Hrefs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Rel != nil {
		in, out := &in.Rel, &out.Rel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotsLinksObservation.
func (in *SnapshotsLinksObservation) DeepCopy() *SnapshotsLinksObservation {
	if in == nil {
		return nil
	}
	out := new(SnapshotsLinksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotsLinksParameters) DeepCopyInto(out *SnapshotsLinksParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotsLinksParameters.
func (in *SnapshotsLinksParameters) DeepCopy() *SnapshotsLinksParameters {
	if in == nil {
		return nil
	}
	out := new(SnapshotsLinksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotsObservation) DeepCopyInto(out *SnapshotsObservation) {
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
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsCurrent != nil {
		in, out := &in.IsCurrent, &out.IsCurrent
		*out = new(bool)
		**out = **in
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]SnapshotsLinksObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotsObservation.
func (in *SnapshotsObservation) DeepCopy() *SnapshotsObservation {
	if in == nil {
		return nil
	}
	out := new(SnapshotsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotsParameters) DeepCopyInto(out *SnapshotsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotsParameters.
func (in *SnapshotsParameters) DeepCopy() *SnapshotsParameters {
	if in == nil {
		return nil
	}
	out := new(SnapshotsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagsInitParameters) DeepCopyInto(out *TagsInitParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagsInitParameters.
func (in *TagsInitParameters) DeepCopy() *TagsInitParameters {
	if in == nil {
		return nil
	}
	out := new(TagsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagsObservation) DeepCopyInto(out *TagsObservation) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagsObservation.
func (in *TagsObservation) DeepCopy() *TagsObservation {
	if in == nil {
		return nil
	}
	out := new(TagsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagsParameters) DeepCopyInto(out *TagsParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagsParameters.
func (in *TagsParameters) DeepCopy() *TagsParameters {
	if in == nil {
		return nil
	}
	out := new(TagsParameters)
	in.DeepCopyInto(out)
	return out
}
