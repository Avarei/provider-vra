/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	blockdevice "github.com/avarei/provider-vra/internal/controller/blockdevice/blockdevice"
	blockdevicesnapshot "github.com/avarei/provider-vra/internal/controller/blockdevice/blockdevicesnapshot"
	blueprint "github.com/avarei/provider-vra/internal/controller/blueprint/blueprint"
	version "github.com/avarei/provider-vra/internal/controller/blueprint/version"
	catalogitementitlement "github.com/avarei/provider-vra/internal/controller/catalogitementitlement/catalogitementitlement"
	catalogsourceblueprint "github.com/avarei/provider-vra/internal/controller/catalogsourceblueprint/catalogsourceblueprint"
	catalogsourceentitlement "github.com/avarei/provider-vra/internal/controller/catalogsourceentitlement/catalogsourceentitlement"
	accountaws "github.com/avarei/provider-vra/internal/controller/cloudaccount/accountaws"
	accountazure "github.com/avarei/provider-vra/internal/controller/cloudaccount/accountazure"
	accountgcp "github.com/avarei/provider-vra/internal/controller/cloudaccount/accountgcp"
	accountnsxt "github.com/avarei/provider-vra/internal/controller/cloudaccount/accountnsxt"
	accountvmc "github.com/avarei/provider-vra/internal/controller/cloudaccount/accountvmc"
	accountvsphere "github.com/avarei/provider-vra/internal/controller/cloudaccount/accountvsphere"
	contentsharingpolicy "github.com/avarei/provider-vra/internal/controller/contentsharing/contentsharingpolicy"
	contentsource "github.com/avarei/provider-vra/internal/controller/contentsource/contentsource"
	deployment "github.com/avarei/provider-vra/internal/controller/deployment/deployment"
	compute "github.com/avarei/provider-vra/internal/controller/fabric/compute"
	datastorevsphere "github.com/avarei/provider-vra/internal/controller/fabric/datastorevsphere"
	networkvsphere "github.com/avarei/provider-vra/internal/controller/fabric/networkvsphere"
	profile "github.com/avarei/provider-vra/internal/controller/flavorprofile/profile"
	profileimageprofile "github.com/avarei/provider-vra/internal/controller/imageprofile/profile"
	integration "github.com/avarei/provider-vra/internal/controller/integration/integration"
	loadbalancer "github.com/avarei/provider-vra/internal/controller/loadbalancer/loadbalancer"
	machine "github.com/avarei/provider-vra/internal/controller/machine/machine"
	network "github.com/avarei/provider-vra/internal/controller/network/network"
	networkiprange "github.com/avarei/provider-vra/internal/controller/network/networkiprange"
	networkprofile "github.com/avarei/provider-vra/internal/controller/network/networkprofile"
	project "github.com/avarei/provider-vra/internal/controller/project/project"
	providerconfig "github.com/avarei/provider-vra/internal/controller/providerconfig"
	profilestorage "github.com/avarei/provider-vra/internal/controller/storage/profile"
	profileaws "github.com/avarei/provider-vra/internal/controller/storage/profileaws"
	profileazure "github.com/avarei/provider-vra/internal/controller/storage/profileazure"
	profilevsphere "github.com/avarei/provider-vra/internal/controller/storage/profilevsphere"
	zone "github.com/avarei/provider-vra/internal/controller/zone/zone"
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
		contentsharingpolicy.Setup,
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
