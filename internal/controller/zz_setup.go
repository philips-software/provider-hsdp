// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	organization "github.com/philips-software/provider-hsdp/internal/controller/cdr/organization"
	subscription "github.com/philips-software/provider-hsdp/internal/controller/cdr/subscription"
	sqssubscriber "github.com/philips-software/provider-hsdp/internal/controller/dbs/sqssubscriber"
	topicsubscription "github.com/philips-software/provider-hsdp/internal/controller/dbs/topicsubscription"
	notification "github.com/philips-software/provider-hsdp/internal/controller/dicom/notification"
	objectstore "github.com/philips-software/provider-hsdp/internal/controller/dicom/objectstore"
	repository "github.com/philips-software/provider-hsdp/internal/controller/dicom/repository"
	storeconfig "github.com/philips-software/provider-hsdp/internal/controller/dicom/storeconfig"
	application "github.com/philips-software/provider-hsdp/internal/controller/iam/application"
	client "github.com/philips-software/provider-hsdp/internal/controller/iam/client"
	emailtemplate "github.com/philips-software/provider-hsdp/internal/controller/iam/emailtemplate"
	group "github.com/philips-software/provider-hsdp/internal/controller/iam/group"
	organizationiam "github.com/philips-software/provider-hsdp/internal/controller/iam/organization"
	passwordpolicy "github.com/philips-software/provider-hsdp/internal/controller/iam/passwordpolicy"
	proposition "github.com/philips-software/provider-hsdp/internal/controller/iam/proposition"
	role "github.com/philips-software/provider-hsdp/internal/controller/iam/role"
	rolesharingpolicy "github.com/philips-software/provider-hsdp/internal/controller/iam/rolesharingpolicy"
	service "github.com/philips-software/provider-hsdp/internal/controller/iam/service"
	user "github.com/philips-software/provider-hsdp/internal/controller/iam/user"
	providerconfig "github.com/philips-software/provider-hsdp/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		organization.Setup,
		subscription.Setup,
		sqssubscriber.Setup,
		topicsubscription.Setup,
		notification.Setup,
		objectstore.Setup,
		repository.Setup,
		storeconfig.Setup,
		application.Setup,
		client.Setup,
		emailtemplate.Setup,
		group.Setup,
		organizationiam.Setup,
		passwordpolicy.Setup,
		proposition.Setup,
		role.Setup,
		rolesharingpolicy.Setup,
		service.Setup,
		user.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
