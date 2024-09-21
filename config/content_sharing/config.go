package contentsharingpolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_content_sharing_policy", func(r *config.Resource) {
		r.ShortGroup = "contentsharing"
		r.Kind = "ContentSharingPolicy"
		r.Version = "v1alpha1"
		r.References["project_id"] = config.Reference{
			Type: "github.com/avarei/provider-vra/apis/project/v1alpha1.Project",
		}
	})
}
