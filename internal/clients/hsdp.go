/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/philips-software/provider-hsdp/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal hsdp credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string, tfProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]any{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = map[string]interface{}{
			"region":              creds["region"],
			"environment":         creds["environment"],
			"service_id":          creds["service_id"],
			"service_private_key": creds["service_private_key"],
			"oauth2_client_id":    creds["oauth2_client_id"],
			"oauth2_password":     creds["oauth2_password"],
			"cartel_host":         creds["cartel_host"],
			"cartel_secret":       creds["cartel_secret"],
			"cartel_token":        creds["cartel_token"],
			"uaa_username":        creds["uaa_username"],
			"uaa_password":        creds["uaa_password"],
			"debug_log":           creds["debug_log"],
			"debug_stderr":        creds["debug_stderr"],
		}

		err = configureNoForkClient(ctx, &ps, *tfProvider)
		if err != nil {
			return ps, errors.Wrap(err, "Unable to set terraform provider setup.Meta")
		}

		return ps, nil
	}
}

func configureNoForkClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	diag := p.Configure(context.WithoutCancel(ctx), &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	return nil
}
