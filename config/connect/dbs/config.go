package dbs

/*
Copyright 2024 Koninklijke Philips N.V., https://www.philips.com
*/

import (
	"github.com/crossplane/upjet/pkg/config"
)

const (
	shortGroup = "dbs"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// SqsSubscriber
	p.AddResourceConfigurator("hsdp_dbs_sqs_subscriber", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SqsSubscriber"
	})
	// Subscription
	p.AddResourceConfigurator("hsdp_dbs_topic_subscription", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TopicSubscription"
		r.References["subscriber_id"] = config.Reference{
			Type:         "SqsSubscriber",
			RefFieldName: "SubscriberRef",
		}
	})
}
