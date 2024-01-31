package mdm

/*
Copyright 2024 Koninklijke Philips N.V., https://www.philips.com
*/

import (
	"github.com/crossplane/upjet/pkg/config"
)

const (
	shortGroup = "mdm"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// MDM Proposition
	p.AddResourceConfigurator("hsdp_connect_mdm_proposition", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Proposition"
		r.References["organization_id"] = config.Reference{
			Type:              "github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization",
			RefFieldName:      "OrganizationRef",
			SelectorFieldName: "OrganizationSelector",
		}
	})
	// MDM DataType
	p.AddResourceConfigurator("hsdp_connect_mdm_data_type", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DataType"
		r.References["proposition_id"] = config.Reference{
			Type:              "Proposition",
			RefFieldName:      "PropositionRef",
			SelectorFieldName: "PropositionSelector",
		}
	})
}
