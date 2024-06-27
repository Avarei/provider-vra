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

type CatalogItemEntitlementObservation struct {

	// Catalog item id.
	CatalogItemID *string `json:"catalogItemId,omitempty" tf:"catalog_item_id,omitempty"`

	Definition []DefinitionObservation `json:"definition,omitempty" tf:"definition,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Project id.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type CatalogItemEntitlementParameters struct {

	// Catalog item id.
	// +kubebuilder:validation:Optional
	CatalogItemID *string `json:"catalogItemId,omitempty" tf:"catalog_item_id,omitempty"`

	// Project id.
	// +crossplane:generate:reference:type=github.com/avarei/provider-vra/apis/project/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

type DefinitionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IconID *string `json:"iconId,omitempty" tf:"icon_id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NumberOfItems *float64 `json:"numberOfItems,omitempty" tf:"number_of_items,omitempty"`

	SourceName *string `json:"sourceName,omitempty" tf:"source_name,omitempty"`

	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DefinitionParameters struct {
}

// CatalogItemEntitlementSpec defines the desired state of CatalogItemEntitlement
type CatalogItemEntitlementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatalogItemEntitlementParameters `json:"forProvider"`
}

// CatalogItemEntitlementStatus defines the observed state of CatalogItemEntitlement.
type CatalogItemEntitlementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatalogItemEntitlementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogItemEntitlement is the Schema for the CatalogItemEntitlements API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type CatalogItemEntitlement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.catalogItemId)",message="catalogItemId is a required parameter"
	Spec   CatalogItemEntitlementSpec   `json:"spec"`
	Status CatalogItemEntitlementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogItemEntitlementList contains a list of CatalogItemEntitlements
type CatalogItemEntitlementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogItemEntitlement `json:"items"`
}

// Repository type metadata.
var (
	CatalogItemEntitlement_Kind             = "CatalogItemEntitlement"
	CatalogItemEntitlement_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CatalogItemEntitlement_Kind}.String()
	CatalogItemEntitlement_KindAPIVersion   = CatalogItemEntitlement_Kind + "." + CRDGroupVersion.String()
	CatalogItemEntitlement_GroupVersionKind = CRDGroupVersion.WithKind(CatalogItemEntitlement_Kind)
)

func init() {
	SchemeBuilder.Register(&CatalogItemEntitlement{}, &CatalogItemEntitlementList{})
}
