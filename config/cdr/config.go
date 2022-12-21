package cdr

/*
Copyright 2022 Koninklijke Philips N.V., https://www.philips.com
*/

import (
	"github.com/upbound/upjet/pkg/config"
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
	})
}
