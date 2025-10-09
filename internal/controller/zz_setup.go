/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	sqssubscriber "github.com/philips-software/provider-hsdp/internal/controller/dbs/sqssubscriber"
	subscription "github.com/philips-software/provider-hsdp/internal/controller/dbs/subscription"
	application "github.com/philips-software/provider-hsdp/internal/controller/iam/application"
	client "github.com/philips-software/provider-hsdp/internal/controller/iam/client"
	emailtemplate "github.com/philips-software/provider-hsdp/internal/controller/iam/emailtemplate"
	group "github.com/philips-software/provider-hsdp/internal/controller/iam/group"
	organization "github.com/philips-software/provider-hsdp/internal/controller/iam/organization"
	passwordpolicy "github.com/philips-software/provider-hsdp/internal/controller/iam/passwordpolicy"
	proposition "github.com/philips-software/provider-hsdp/internal/controller/iam/proposition"
	role "github.com/philips-software/provider-hsdp/internal/controller/iam/role"
	rolesharingpolicy "github.com/philips-software/provider-hsdp/internal/controller/iam/rolesharingpolicy"
	service "github.com/philips-software/provider-hsdp/internal/controller/iam/service"
	user "github.com/philips-software/provider-hsdp/internal/controller/iam/user"
	datatype "github.com/philips-software/provider-hsdp/internal/controller/mdm/datatype"
	propositionmdm "github.com/philips-software/provider-hsdp/internal/controller/mdm/proposition"
	providerconfig "github.com/philips-software/provider-hsdp/internal/controller/providerconfig"
	key "github.com/philips-software/provider-hsdp/internal/controller/tenant/key"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		sqssubscriber.Setup,
		subscription.Setup,
		application.Setup,
		client.Setup,
		emailtemplate.Setup,
		group.Setup,
		organization.Setup,
		passwordpolicy.Setup,
		proposition.Setup,
		role.Setup,
		rolesharingpolicy.Setup,
		service.Setup,
		user.Setup,
		datatype.Setup,
		propositionmdm.Setup,
		providerconfig.Setup,
		key.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
