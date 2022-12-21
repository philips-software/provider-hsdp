/*
Copyright 2022 Koninklijke Philips N.V., https://www.philips.com
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"hsdp_iam_group":               config.IdentifierFromProvider,
	"hsdp_iam_org":                 config.IdentifierFromProvider,
	"hsdp_iam_application":         config.IdentifierFromProvider,
	"hsdp_iam_client":              config.IdentifierFromProvider,
	"hsdp_iam_email_template":      config.IdentifierFromProvider,
	"hsdp_iam_proposition":         config.IdentifierFromProvider,
	"hsdp_iam_service":             config.IdentifierFromProvider,
	"hsdp_iam_role":                config.IdentifierFromProvider,
	"hsdp_iam_password_policy":     config.IdentifierFromProvider,
	"hsdp_iam_role_sharing_policy": config.IdentifierFromProvider,
	"hsdp_iam_user":                config.IdentifierFromProvider,
	"hsdp_cdr_org":                 config.IdentifierFromProvider,
	"hsdp_cdr_subscription":        config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
