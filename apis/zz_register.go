/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/avarei/provider-vra/apis/blockdevice/v1alpha1"
	v1alpha1blueprint "github.com/avarei/provider-vra/apis/blueprint/v1alpha1"
	v1alpha1catalogitementitlement "github.com/avarei/provider-vra/apis/catalogitementitlement/v1alpha1"
	v1alpha1catalogsourceblueprint "github.com/avarei/provider-vra/apis/catalogsourceblueprint/v1alpha1"
	v1alpha1catalogsourceentitlement "github.com/avarei/provider-vra/apis/catalogsourceentitlement/v1alpha1"
	v1alpha1cloudaccount "github.com/avarei/provider-vra/apis/cloudaccount/v1alpha1"
	v1alpha1contentsharing "github.com/avarei/provider-vra/apis/contentsharing/v1alpha1"
	v1alpha1contentsource "github.com/avarei/provider-vra/apis/contentsource/v1alpha1"
	v1alpha1deployment "github.com/avarei/provider-vra/apis/deployment/v1alpha1"
	v1alpha1fabric "github.com/avarei/provider-vra/apis/fabric/v1alpha1"
	v1alpha1flavorprofile "github.com/avarei/provider-vra/apis/flavorprofile/v1alpha1"
	v1alpha1imageprofile "github.com/avarei/provider-vra/apis/imageprofile/v1alpha1"
	v1alpha1integration "github.com/avarei/provider-vra/apis/integration/v1alpha1"
	v1alpha1loadbalancer "github.com/avarei/provider-vra/apis/loadbalancer/v1alpha1"
	v1alpha1machine "github.com/avarei/provider-vra/apis/machine/v1alpha1"
	v1alpha1network "github.com/avarei/provider-vra/apis/network/v1alpha1"
	v1alpha1project "github.com/avarei/provider-vra/apis/project/v1alpha1"
	v1alpha1storage "github.com/avarei/provider-vra/apis/storage/v1alpha1"
	v1alpha1apis "github.com/avarei/provider-vra/apis/v1alpha1"
	v1beta1 "github.com/avarei/provider-vra/apis/v1beta1"
	v1alpha1zone "github.com/avarei/provider-vra/apis/zone/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1blueprint.SchemeBuilder.AddToScheme,
		v1alpha1catalogitementitlement.SchemeBuilder.AddToScheme,
		v1alpha1catalogsourceblueprint.SchemeBuilder.AddToScheme,
		v1alpha1catalogsourceentitlement.SchemeBuilder.AddToScheme,
		v1alpha1cloudaccount.SchemeBuilder.AddToScheme,
		v1alpha1contentsharing.SchemeBuilder.AddToScheme,
		v1alpha1contentsource.SchemeBuilder.AddToScheme,
		v1alpha1deployment.SchemeBuilder.AddToScheme,
		v1alpha1fabric.SchemeBuilder.AddToScheme,
		v1alpha1flavorprofile.SchemeBuilder.AddToScheme,
		v1alpha1imageprofile.SchemeBuilder.AddToScheme,
		v1alpha1integration.SchemeBuilder.AddToScheme,
		v1alpha1loadbalancer.SchemeBuilder.AddToScheme,
		v1alpha1machine.SchemeBuilder.AddToScheme,
		v1alpha1network.SchemeBuilder.AddToScheme,
		v1alpha1project.SchemeBuilder.AddToScheme,
		v1alpha1storage.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1zone.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
