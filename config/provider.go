/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	blockdevice "github.com/avarei/provider-vra/config/block_device"
	blueprint "github.com/avarei/provider-vra/config/blueprint"
	catalogitementitlement "github.com/avarei/provider-vra/config/catalog_item"
	catalogsource "github.com/avarei/provider-vra/config/catalog_source"
	cloudaccount "github.com/avarei/provider-vra/config/cloud_account"
	contentsource "github.com/avarei/provider-vra/config/content_source"
	deployment "github.com/avarei/provider-vra/config/deployment"
	fabric "github.com/avarei/provider-vra/config/fabric"
	flavorprofile "github.com/avarei/provider-vra/config/flavor_profile"
	imageprofile "github.com/avarei/provider-vra/config/image_profile"
	integration "github.com/avarei/provider-vra/config/integration"
	loadbalancer "github.com/avarei/provider-vra/config/load_balancer"
	machine "github.com/avarei/provider-vra/config/machine"
	network "github.com/avarei/provider-vra/config/network"
	project "github.com/avarei/provider-vra/config/project"
	storage "github.com/avarei/provider-vra/config/storage"
	zone "github.com/avarei/provider-vra/config/zone"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "vra"
	modulePath     = "github.com/avarei/provider-vra"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		project.Configure,
		blueprint.Configure,
		deployment.Configure,
		fabric.Configure,
		blockdevice.Configure,
		flavorprofile.Configure,
		imageprofile.Configure,
		storage.Configure,
		catalogsource.Configure,
		catalogitementitlement.Configure,
		cloudaccount.Configure,
		contentsource.Configure,
		integration.Configure,
		loadbalancer.Configure,
		machine.Configure,
		network.Configure,
		zone.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
