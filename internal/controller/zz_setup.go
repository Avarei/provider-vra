/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	blockdevice "github.com/avarei/provider-vra/cmd/provider/internal/controller/blockdevice/blockdevice"
	blockdevicesnapshot "github.com/avarei/provider-vra/cmd/provider/internal/controller/blockdevice/blockdevicesnapshot"
	blueprint "github.com/avarei/provider-vra/cmd/provider/internal/controller/blueprint/blueprint"
	version "github.com/avarei/provider-vra/cmd/provider/internal/controller/blueprint/version"
	catalogitementitlement "github.com/avarei/provider-vra/cmd/provider/internal/controller/catalogitementitlement/catalogitementitlement"
	catalogsourceblueprint "github.com/avarei/provider-vra/cmd/provider/internal/controller/catalogsourceblueprint/catalogsourceblueprint"
	catalogsourceentitlement "github.com/avarei/provider-vra/cmd/provider/internal/controller/catalogsourceentitlement/catalogsourceentitlement"
	accountaws "github.com/avarei/provider-vra/cmd/provider/internal/controller/cloudaccount/accountaws"
	accountazure "github.com/avarei/provider-vra/cmd/provider/internal/controller/cloudaccount/accountazure"
	accountgcp "github.com/avarei/provider-vra/cmd/provider/internal/controller/cloudaccount/accountgcp"
	accountnsxt "github.com/avarei/provider-vra/cmd/provider/internal/controller/cloudaccount/accountnsxt"
	accountvmc "github.com/avarei/provider-vra/cmd/provider/internal/controller/cloudaccount/accountvmc"
	accountvsphere "github.com/avarei/provider-vra/cmd/provider/internal/controller/cloudaccount/accountvsphere"
	contentsource "github.com/avarei/provider-vra/cmd/provider/internal/controller/contentsource/contentsource"
	deployment "github.com/avarei/provider-vra/cmd/provider/internal/controller/deployment/deployment"
	compute "github.com/avarei/provider-vra/cmd/provider/internal/controller/fabric/compute"
	datastorevsphere "github.com/avarei/provider-vra/cmd/provider/internal/controller/fabric/datastorevsphere"
	networkvsphere "github.com/avarei/provider-vra/cmd/provider/internal/controller/fabric/networkvsphere"
	profile "github.com/avarei/provider-vra/cmd/provider/internal/controller/flavorprofile/profile"
	profileimageprofile "github.com/avarei/provider-vra/cmd/provider/internal/controller/imageprofile/profile"
	integration "github.com/avarei/provider-vra/cmd/provider/internal/controller/integration/integration"
	loadbalancer "github.com/avarei/provider-vra/cmd/provider/internal/controller/loadbalancer/loadbalancer"
	machine "github.com/avarei/provider-vra/cmd/provider/internal/controller/machine/machine"
	network "github.com/avarei/provider-vra/cmd/provider/internal/controller/network/network"
	networkiprange "github.com/avarei/provider-vra/cmd/provider/internal/controller/network/networkiprange"
	networkprofile "github.com/avarei/provider-vra/cmd/provider/internal/controller/network/networkprofile"
	project "github.com/avarei/provider-vra/cmd/provider/internal/controller/project/project"
	providerconfig "github.com/avarei/provider-vra/cmd/provider/internal/controller/providerconfig"
	profilestorage "github.com/avarei/provider-vra/cmd/provider/internal/controller/storage/profile"
	profileaws "github.com/avarei/provider-vra/cmd/provider/internal/controller/storage/profileaws"
	profileazure "github.com/avarei/provider-vra/cmd/provider/internal/controller/storage/profileazure"
	profilevsphere "github.com/avarei/provider-vra/cmd/provider/internal/controller/storage/profilevsphere"
	zone "github.com/avarei/provider-vra/cmd/provider/internal/controller/zone/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		blockdevice.Setup,
		blockdevicesnapshot.Setup,
		blueprint.Setup,
		version.Setup,
		catalogitementitlement.Setup,
		catalogsourceblueprint.Setup,
		catalogsourceentitlement.Setup,
		accountaws.Setup,
		accountazure.Setup,
		accountgcp.Setup,
		accountnsxt.Setup,
		accountvmc.Setup,
		accountvsphere.Setup,
		contentsource.Setup,
		deployment.Setup,
		compute.Setup,
		datastorevsphere.Setup,
		networkvsphere.Setup,
		profile.Setup,
		profileimageprofile.Setup,
		integration.Setup,
		loadbalancer.Setup,
		machine.Setup,
		network.Setup,
		networkiprange.Setup,
		networkprofile.Setup,
		project.Setup,
		providerconfig.Setup,
		profilestorage.Setup,
		profileaws.Setup,
		profileazure.Setup,
		profilevsphere.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
