package dicom

/*
Copyright 2023 Koninklijke Philips N.V., https://www.philips.com
*/

import (
	"github.com/crossplane/upjet/pkg/config"
)

const (
	shortGroup = "dicom"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// Application
	p.AddResourceConfigurator("hsdp_dicom_notification", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["default_organization_id"] = config.Reference{
			Type:         "github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization",
			RefFieldName: "OrganizationRef",
		}
	})

	p.AddResourceConfigurator("hsdp_dicom_object_store", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["organization_id"] = config.Reference{
			Type:         "github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization",
			RefFieldName: "OrganizationRef",
		}
	})

	p.AddResourceConfigurator("hsdp_dicom_repository", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["organization_id"] = config.Reference{
			Type:         "github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization",
			RefFieldName: "OrganizationRef",
		}
	})

	p.AddResourceConfigurator("hsdp_dicom_store_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["organization_id"] = config.Reference{
			Type:         "github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization",
			RefFieldName: "OrganizationRef",
		}
	})

}
