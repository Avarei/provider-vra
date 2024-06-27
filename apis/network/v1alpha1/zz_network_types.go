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

type ConstraintsObservation struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Indicates whether this constraint should be strictly enforced or not.
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`
}

type ConstraintsParameters struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// Indicates whether this constraint should be strictly enforced or not.
	// +kubebuilder:validation:Required
	Mandatory *bool `json:"mandatory" tf:"mandatory,omitempty"`
}

type LinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type LinksParameters struct {
}

type NetworkObservation struct {
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Constraints that are used to drive placement policies for entities such as image, network, storage, etc. Constraint expressions are matched against tags on existing placement targets.
	Constraints []ConstraintsObservation `json:"constraints,omitempty" tf:"constraints,omitempty"`

	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	ExternalZoneID *string `json:"externalZoneId,omitempty" tf:"external_zone_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Links []LinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	OutboundAccess *bool `json:"outboundAccess,omitempty" tf:"outbound_access,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	Tags []TagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type NetworkParameters struct {

	// Constraints that are used to drive placement policies for entities such as image, network, storage, etc. Constraint expressions are matched against tags on existing placement targets.
	// +kubebuilder:validation:Optional
	Constraints []ConstraintsParameters `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// +kubebuilder:validation:Optional
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OutboundAccess *bool `json:"outboundAccess,omitempty" tf:"outbound_access,omitempty"`

	// +crossplane:generate:reference:type=github.com/avarei/provider-vra/cmd/provider/apis/project/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// NetworkSpec defines the desired state of Network
type NetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkParameters `json:"forProvider"`
}

// NetworkStatus defines the observed state of Network.
type NetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Network is the Schema for the Networks API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   NetworkSpec   `json:"spec"`
	Status NetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkList contains a list of Networks
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

// Repository type metadata.
var (
	Network_Kind             = "Network"
	Network_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Network_Kind}.String()
	Network_KindAPIVersion   = Network_Kind + "." + CRDGroupVersion.String()
	Network_GroupVersionKind = CRDGroupVersion.WithKind(Network_Kind)
)

func init() {
	SchemeBuilder.Register(&Network{}, &NetworkList{})
}
