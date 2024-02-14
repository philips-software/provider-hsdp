// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SubscriptionInitParameters struct {

	// +crossplane:generate:reference:type=github.com/philips-software/provider-hsdp/apis/mdm/v1alpha1.DataType
	// +crossplane:generate:reference:extractor=github.com/philips-software/provider-hsdp/config/common.ExtractResourceName()
	// +crossplane:generate:reference:refFieldName=DataTypeRef
	// +crossplane:generate:reference:selectorFieldName=DataTypeSelector
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Reference to a DataType in mdm to populate dataType.
	// +kubebuilder:validation:Optional
	DataTypeRef *v1.Reference `json:"dataTypeRef,omitempty" tf:"-"`

	// Selector for a DataType in mdm to populate dataType.
	// +kubebuilder:validation:Optional
	DataTypeSelector *v1.Selector `json:"dataTypeSelector,omitempty" tf:"-"`

	DeliverDataOnly *bool `json:"deliverDataOnly,omitempty" tf:"deliver_data_only,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	KinesisStreamPartitionKey *string `json:"kinesisStreamPartitionKey,omitempty" tf:"kinesis_stream_partition_key,omitempty"`

	NameInfix *string `json:"nameInfix,omitempty" tf:"name_infix,omitempty"`

	// +crossplane:generate:reference:type=SqsSubscriber
	// +crossplane:generate:reference:refFieldName=SubscriberRef
	// +crossplane:generate:reference:selectorFieldName=SubscriberSelector
	SubscriberID *string `json:"subscriberId,omitempty" tf:"subscriber_id,omitempty"`

	// Reference to a SqsSubscriber to populate subscriberId.
	// +kubebuilder:validation:Optional
	SubscriberRef *v1.Reference `json:"subscriberRef,omitempty" tf:"-"`

	// Selector for a SqsSubscriber to populate subscriberId.
	// +kubebuilder:validation:Optional
	SubscriberSelector *v1.Selector `json:"subscriberSelector,omitempty" tf:"-"`
}

type SubscriptionObservation struct {
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	DeliverDataOnly *bool `json:"deliverDataOnly,omitempty" tf:"deliver_data_only,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KinesisStreamPartitionKey *string `json:"kinesisStreamPartitionKey,omitempty" tf:"kinesis_stream_partition_key,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NameInfix *string `json:"nameInfix,omitempty" tf:"name_infix,omitempty"`

	RuleName *string `json:"ruleName,omitempty" tf:"rule_name,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	SubscriberID *string `json:"subscriberId,omitempty" tf:"subscriber_id,omitempty"`
}

type SubscriptionParameters struct {

	// +crossplane:generate:reference:type=github.com/philips-software/provider-hsdp/apis/mdm/v1alpha1.DataType
	// +crossplane:generate:reference:extractor=github.com/philips-software/provider-hsdp/config/common.ExtractResourceName()
	// +crossplane:generate:reference:refFieldName=DataTypeRef
	// +crossplane:generate:reference:selectorFieldName=DataTypeSelector
	// +kubebuilder:validation:Optional
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Reference to a DataType in mdm to populate dataType.
	// +kubebuilder:validation:Optional
	DataTypeRef *v1.Reference `json:"dataTypeRef,omitempty" tf:"-"`

	// Selector for a DataType in mdm to populate dataType.
	// +kubebuilder:validation:Optional
	DataTypeSelector *v1.Selector `json:"dataTypeSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DeliverDataOnly *bool `json:"deliverDataOnly,omitempty" tf:"deliver_data_only,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	KinesisStreamPartitionKey *string `json:"kinesisStreamPartitionKey,omitempty" tf:"kinesis_stream_partition_key,omitempty"`

	// +kubebuilder:validation:Optional
	NameInfix *string `json:"nameInfix,omitempty" tf:"name_infix,omitempty"`

	// +crossplane:generate:reference:type=SqsSubscriber
	// +crossplane:generate:reference:refFieldName=SubscriberRef
	// +crossplane:generate:reference:selectorFieldName=SubscriberSelector
	// +kubebuilder:validation:Optional
	SubscriberID *string `json:"subscriberId,omitempty" tf:"subscriber_id,omitempty"`

	// Reference to a SqsSubscriber to populate subscriberId.
	// +kubebuilder:validation:Optional
	SubscriberRef *v1.Reference `json:"subscriberRef,omitempty" tf:"-"`

	// Selector for a SqsSubscriber to populate subscriberId.
	// +kubebuilder:validation:Optional
	SubscriberSelector *v1.Selector `json:"subscriberSelector,omitempty" tf:"-"`
}

// SubscriptionSpec defines the desired state of Subscription
type SubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SubscriptionInitParameters `json:"initProvider,omitempty"`
}

// SubscriptionStatus defines the observed state of Subscription.
type SubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subscription is the Schema for the Subscriptions API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nameInfix) || (has(self.initProvider) && has(self.initProvider.nameInfix))",message="spec.forProvider.nameInfix is a required parameter"
	Spec   SubscriptionSpec   `json:"spec"`
	Status SubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionList contains a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

// Repository type metadata.
var (
	Subscription_Kind             = "Subscription"
	Subscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subscription_Kind}.String()
	Subscription_KindAPIVersion   = Subscription_Kind + "." + CRDGroupVersion.String()
	Subscription_GroupVersionKind = CRDGroupVersion.WithKind(Subscription_Kind)
)

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}