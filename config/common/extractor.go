package common

import (
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	jresource "github.com/upbound/upjet/pkg/resource"
)

const (
	configPackagePath = "github.com/philips-software/provider-hsdp/config"
	// ExtractResourceIDFuncPath contains the path to a custom extractor for Terraform managed resources
	ExtractResourceIDFuncPath = configPackagePath + "/common.ExtractResourceID()"
)

// ExtractResourceID extracts the value of `spec.atProvider.id`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractResourceID() reference.ExtractValueFn {
	return func(mr resource.Managed) string {
		tr, ok := mr.(jresource.Terraformed)
		if !ok {
			return ""
		}
		return tr.GetID()
	}
}
