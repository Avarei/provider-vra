/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/

// Code generated by upjet. DO NOT EDIT.

package networkiprange

import (
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	tjcontroller "github.com/upbound/upjet/pkg/controller"
	"github.com/upbound/upjet/pkg/terraform"
	ctrl "sigs.k8s.io/controller-runtime"

	v1alpha1 "github.com/avarei/provider-vra/cmd/provider/apis/network/v1alpha1"
	features "github.com/avarei/provider-vra/cmd/provider/internal/features"
)

// Setup adds a controller that reconciles NetworkIPRange managed resources.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1alpha1.NetworkIPRange_GroupVersionKind.String())
	var initializers managed.InitializerChain
	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.SecretStoreConfigGVK != nil {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), *o.SecretStoreConfigGVK))
	}
	opts := []managed.ReconcilerOption{
		managed.WithExternalConnecter(tjcontroller.NewConnector(mgr.GetClient(), o.WorkspaceStore, o.SetupFn, o.Provider.Resources["vra_network_ip_range"], tjcontroller.WithLogger(o.Logger),
			tjcontroller.WithCallbackProvider(tjcontroller.NewAPICallbacks(mgr, xpresource.ManagedKind(v1alpha1.NetworkIPRange_GroupVersionKind))),
		)),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithFinalizer(terraform.NewWorkspaceFinalizer(o.WorkspaceStore, xpresource.NewAPIFinalizer(mgr.GetClient(), managed.FinalizerName))),
		managed.WithTimeout(3 * time.Minute),
		managed.WithInitializers(initializers),
		managed.WithConnectionPublishers(cps...),
		managed.WithPollInterval(o.PollInterval),
	}
	if o.Features.Enabled(features.EnableAlphaManagementPolicies) {
		opts = append(opts, managed.WithManagementPolicies())
	}
	r := managed.NewReconciler(mgr, xpresource.ManagedKind(v1alpha1.NetworkIPRange_GroupVersionKind), opts...)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha1.NetworkIPRange{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}
