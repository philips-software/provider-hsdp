package cdr

/*
Copyright 2022 Koninklijke Philips N.V., https://www.philips.com
*/

import (
	"github.com/crossplane/upjet/pkg/config"
)

const (
	shortGroup = "cdr"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// Organization
	p.AddResourceConfigurator("hsdp_cdr_org", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Organization"
		r.References["org_id"] = config.Reference{
			Type:         "github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization",
			RefFieldName: "OrganizationRef",
		}
	})
	// Subscription
	p.AddResourceConfigurator("hsdp_cdr_subscription", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Subscription"
	})
}
