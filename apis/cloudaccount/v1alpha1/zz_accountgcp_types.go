/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccountGCPInitParameters struct {

	// GCP Client email.
	ClientEmail *string `json:"clientEmail,omitempty" tf:"client_email,omitempty"`

	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of this resource instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// GCP Private key ID.
	PrivateKeyID *string `json:"privateKeyId,omitempty" tf:"private_key_id,omitempty"`

	// GCP Private key.
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// GCP Project ID.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The set of region ids that will be enabled for this cloud account.
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	Tags []AccountGCPTagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AccountGCPLinksInitParameters struct {
}

type AccountGCPLinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// +listType=set
	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type AccountGCPLinksParameters struct {
}

type AccountGCPObservation struct {

	// GCP Client email.
	ClientEmail *string `json:"clientEmail,omitempty" tf:"client_email,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Links []AccountGCPLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// The name of this resource instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the organization this entity belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Email of the user that owns the entity.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// GCP Private key ID.
	PrivateKeyID *string `json:"privateKeyId,omitempty" tf:"private_key_id,omitempty"`

	// GCP Project ID.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The set of region ids that will be enabled for this cloud account.
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	Tags []AccountGCPTagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type AccountGCPParameters struct {

	// GCP Client email.
	// +kubebuilder:validation:Optional
	ClientEmail *string `json:"clientEmail,omitempty" tf:"client_email,omitempty"`

	// A human-friendly description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of this resource instance.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// GCP Private key ID.
	// +kubebuilder:validation:Optional
	PrivateKeyID *string `json:"privateKeyId,omitempty" tf:"private_key_id,omitempty"`

	// GCP Private key.
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// GCP Project ID.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The set of region ids that will be enabled for this cloud account.
	// +kubebuilder:validation:Optional
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []AccountGCPTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AccountGCPTagsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AccountGCPTagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AccountGCPTagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// AccountGCPSpec defines the desired state of AccountGCP
type AccountGCPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountGCPParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AccountGCPInitParameters `json:"initProvider,omitempty"`
}

// AccountGCPStatus defines the observed state of AccountGCP.
type AccountGCPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountGCPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccountGCP is the Schema for the AccountGCPs API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type AccountGCP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientEmail) || (has(self.initProvider) && has(self.initProvider.clientEmail))",message="spec.forProvider.clientEmail is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privateKeySecretRef)",message="spec.forProvider.privateKeySecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privateKeyId) || (has(self.initProvider) && has(self.initProvider.privateKeyId))",message="spec.forProvider.privateKeyId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectId) || (has(self.initProvider) && has(self.initProvider.projectId))",message="spec.forProvider.projectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.regions) || (has(self.initProvider) && has(self.initProvider.regions))",message="spec.forProvider.regions is a required parameter"
	Spec   AccountGCPSpec   `json:"spec"`
	Status AccountGCPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountGCPList contains a list of AccountGCPs
type AccountGCPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountGCP `json:"items"`
}

// Repository type metadata.
var (
	AccountGCP_Kind             = "AccountGCP"
	AccountGCP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccountGCP_Kind}.String()
	AccountGCP_KindAPIVersion   = AccountGCP_Kind + "." + CRDGroupVersion.String()
	AccountGCP_GroupVersionKind = CRDGroupVersion.WithKind(AccountGCP_Kind)
)

func init() {
	SchemeBuilder.Register(&AccountGCP{}, &AccountGCPList{})
}
