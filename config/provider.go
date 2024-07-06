/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/philips-software/terraform-provider-hsdp/hsdp"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/philips-software/provider-hsdp/config/cdr"
	"github.com/philips-software/provider-hsdp/config/connect/dbs"
	"github.com/philips-software/provider-hsdp/config/connect/mdm"
	"github.com/philips-software/provider-hsdp/config/dicom"
	"github.com/philips-software/provider-hsdp/config/iam"
)

const (
	resourcePrefix = "hsdp"
	modulePath     = "github.com/philips-software/provider-hsdp"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithShortName("hsdp"),
		ujconfig.WithRootGroup("hsdp.crossplane.io"),
		ujconfig.WithTerraformPluginSDKIncludeList(ExternalNameConfigured()),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(hsdp.Provider("v0.47.0")),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		iam.Configure,
		cdr.Configure,
		dbs.Configure,
		mdm.Configure,
		dicom.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
