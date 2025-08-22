package iam

/*
Copyright 2022 Koninklijke Philips N.V., https://www.philips.com
*/

import (
	"github.com/crossplane/upjet/pkg/config"
)

const (
	shortGroup = "iam"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// Application
	p.AddResourceConfigurator("hsdp_iam_application", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["proposition_id"] = config.Reference{
			TerraformName: "hsdp_iam_proposition",
			RefFieldName:  "PropositionRef",
		}
	})

	// Client
	p.AddResourceConfigurator("hsdp_iam_client", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["application_id"] = config.Reference{
			TerraformName: "hsdp_iam_application",
			RefFieldName:  "ApplicationRef",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["client_id"].(string); ok {
				conn["client_id"] = []byte(a)
			}
			return conn, nil
		}
	})

	// EmailTemplate
	p.AddResourceConfigurator("hsdp_iam_email_template", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["managing_organization"] = config.Reference{
			TerraformName: "hsdp_iam_org",
			RefFieldName:  "OrganizationRef",
		}
	})

	// Group
	p.AddResourceConfigurator("hsdp_iam_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["managing_organization"] = config.Reference{
			TerraformName: "hsdp_iam_org",
			RefFieldName:  "OrganizationRef",
		}
		r.References["users"] = config.Reference{
			TerraformName: "hsdp_iam_user",
			RefFieldName:  "UserRef",
		}
		r.References["services"] = config.Reference{
			TerraformName: "hsdp_iam_service",
			RefFieldName:  "ServiceRef",
		}
		r.References["roles"] = config.Reference{
			TerraformName: "hsdp_iam_role",
			RefFieldName:  "RoleRef",
		}
	})

	// Organization
	p.AddResourceConfigurator("hsdp_iam_org", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Organization"
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["id"].(string); ok {
				conn["organization_id"] = []byte(a)
			}
			return conn, nil
		}
	})

	// PasswordPolicy
	p.AddResourceConfigurator("hsdp_iam_password_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["managing_organization"] = config.Reference{
			TerraformName: "hsdp_iam_org",
			RefFieldName:  "OrganizationRef",
		}
	})

	// Proposition
	p.AddResourceConfigurator("hsdp_iam_proposition", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["organization_id"] = config.Reference{
			TerraformName: "hsdp_iam_org",
			RefFieldName:  "OrganizationRef",
		}
	})

	// Role
	p.AddResourceConfigurator("hsdp_iam_role", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["managing_organization"] = config.Reference{
			TerraformName: "hsdp_iam_org",
			RefFieldName:  "OrganizationRef",
		}
	})

	// RoleSharingPolicy
	p.AddResourceConfigurator("hsdp_iam_role_sharing_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["role_id"] = config.Reference{
			TerraformName: "hsdp_iam_role",
			RefFieldName:  "RoleRef",
		}
		r.References["target_organization_id"] = config.Reference{
			TerraformName: "hsdp_iam_org",
			RefFieldName:  "OrganizationRef",
		}
	})

	// Service
	p.AddResourceConfigurator("hsdp_iam_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["application_id"] = config.Reference{
			TerraformName: "hsdp_iam_application",
			RefFieldName:  "ApplicationRef",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["service_id"].(string); ok {
				conn["service_id"] = []byte(a)
			}
			if a, ok := attr["self_managed_certificate"].(string); ok {
				conn["self_managed_certificate"] = []byte(a)
			}
			return conn, nil
		}
	})

	// User
	p.AddResourceConfigurator("hsdp_iam_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["organization_id"] = config.Reference{
			TerraformName: "hsdp_iam_org",
			RefFieldName:  "OrganizationRef",
		}
	})
}
