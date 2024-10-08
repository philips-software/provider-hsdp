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

type EmailTemplateInitParameters struct {

	// The template format. Must be HTML currently
	// The template format. Must be 'HTML' currently.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The From field of the email. Default value is default
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// A clickable link, depends on the template type
	// A clickable link, depends on the template type.
	Link *string `json:"link,omitempty" tf:"link,omitempty"`

	// The locale of the template. When not specified the template will become the default. Only a single default template is allowed of course.
	// The locale of the template. When not specified the template will become the default. Only a single default template is allowed of course.
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	// The UUID of the IAM Org to apply this email template to
	// The Id of the IAM Org to apply this email template to.
	// +crossplane:generate:reference:type=github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	ManagingOrganization *string `json:"managingOrganization,omitempty" tf:"managing_organization,omitempty"`

	// Selector for a Organization in iam to populate managingOrganization.
	// +kubebuilder:validation:Optional
	ManagingOrganizationSelector *v1.Selector `json:"managingOrganizationSelector,omitempty" tf:"-"`

	// A boolean value indicating if challenges are enabled at organization level. If the value is set to true, challenge_policy attribute is mandatory.
	// The message body.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Reference to a Organization in iam to populate managingOrganization.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// The Subject line of the email. Default value is default
	// The Subject line of the email.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// The email template. See the Type table above for available values
	// The email template type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EmailTemplateObservation struct {

	// The template format. Must be HTML currently
	// The template format. Must be 'HTML' currently.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The From field of the email. Default value is default
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// The GUID of the email template
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A clickable link, depends on the template type
	// A clickable link, depends on the template type.
	Link *string `json:"link,omitempty" tf:"link,omitempty"`

	// The locale of the template. When not specified the template will become the default. Only a single default template is allowed of course.
	// The locale of the template. When not specified the template will become the default. Only a single default template is allowed of course.
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	// The UUID of the IAM Org to apply this email template to
	// The Id of the IAM Org to apply this email template to.
	ManagingOrganization *string `json:"managingOrganization,omitempty" tf:"managing_organization,omitempty"`

	// A boolean value indicating if challenges are enabled at organization level. If the value is set to true, challenge_policy attribute is mandatory.
	// The message body.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	MessageBase64 *string `json:"messageBase64,omitempty" tf:"message_base64,omitempty"`

	// The Subject line of the email. Default value is default
	// The Subject line of the email.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// The email template. See the Type table above for available values
	// The email template type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EmailTemplateParameters struct {

	// The template format. Must be HTML currently
	// The template format. Must be 'HTML' currently.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The From field of the email. Default value is default
	// +kubebuilder:validation:Optional
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// A clickable link, depends on the template type
	// A clickable link, depends on the template type.
	// +kubebuilder:validation:Optional
	Link *string `json:"link,omitempty" tf:"link,omitempty"`

	// The locale of the template. When not specified the template will become the default. Only a single default template is allowed of course.
	// The locale of the template. When not specified the template will become the default. Only a single default template is allowed of course.
	// +kubebuilder:validation:Optional
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	// The UUID of the IAM Org to apply this email template to
	// The Id of the IAM Org to apply this email template to.
	// +crossplane:generate:reference:type=github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +kubebuilder:validation:Optional
	ManagingOrganization *string `json:"managingOrganization,omitempty" tf:"managing_organization,omitempty"`

	// Selector for a Organization in iam to populate managingOrganization.
	// +kubebuilder:validation:Optional
	ManagingOrganizationSelector *v1.Selector `json:"managingOrganizationSelector,omitempty" tf:"-"`

	// A boolean value indicating if challenges are enabled at organization level. If the value is set to true, challenge_policy attribute is mandatory.
	// The message body.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Reference to a Organization in iam to populate managingOrganization.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// The Subject line of the email. Default value is default
	// The Subject line of the email.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// The email template. See the Type table above for available values
	// The email template type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// EmailTemplateSpec defines the desired state of EmailTemplate
type EmailTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EmailTemplateParameters `json:"forProvider"`
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
	InitProvider EmailTemplateInitParameters `json:"initProvider,omitempty"`
}

// EmailTemplateStatus defines the observed state of EmailTemplate.
type EmailTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmailTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EmailTemplate is the Schema for the EmailTemplates API. Manages HSDP IAM Email template resources
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type EmailTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.message) || (has(self.initProvider) && has(self.initProvider.message))",message="spec.forProvider.message is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   EmailTemplateSpec   `json:"spec"`
	Status EmailTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailTemplateList contains a list of EmailTemplates
type EmailTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailTemplate `json:"items"`
}

// Repository type metadata.
var (
	EmailTemplate_Kind             = "EmailTemplate"
	EmailTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EmailTemplate_Kind}.String()
	EmailTemplate_KindAPIVersion   = EmailTemplate_Kind + "." + CRDGroupVersion.String()
	EmailTemplate_GroupVersionKind = CRDGroupVersion.WithKind(EmailTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&EmailTemplate{}, &EmailTemplateList{})
}
