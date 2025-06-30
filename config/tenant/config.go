package tenant

import (
	"github.com/crossplane/upjet/pkg/config"
)

const (
	shortGroup = "tenant"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// Key
	p.AddResourceConfigurator("hsdp_tenant_key", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
