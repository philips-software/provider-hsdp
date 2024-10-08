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

type RoleInitParameters struct {

	// The description of the group
	// The role description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The managing organization ID of this role
	// The managing organization of the role.
	// +crossplane:generate:reference:type=github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	ManagingOrganization *string `json:"managingOrganization,omitempty" tf:"managing_organization,omitempty"`

	// Selector for a Organization in iam to populate managingOrganization.
	// +kubebuilder:validation:Optional
	ManagingOrganizationSelector *v1.Selector `json:"managingOrganizationSelector,omitempty" tf:"-"`

	// The name of the group
	// The role name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Organization in iam to populate managingOrganization.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// The list of permission to assign to this role
	// List of permissions IDs assigned to this role.
	// +listType=set
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Defaults to true. Setting to false will remove e.g. CLIENT.SCOPES permission which is only addable using a HSDP support ticket.
	// Removal protection of some ticket only permissions.
	TicketProtection *bool `json:"ticketProtection,omitempty" tf:"ticket_protection,omitempty"`
}

type RoleObservation struct {

	// The description of the group
	// The role description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The GUID of the role
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The managing organization ID of this role
	// The managing organization of the role.
	ManagingOrganization *string `json:"managingOrganization,omitempty" tf:"managing_organization,omitempty"`

	// The name of the group
	// The role name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The list of permission to assign to this role
	// List of permissions IDs assigned to this role.
	// +listType=set
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Defaults to true. Setting to false will remove e.g. CLIENT.SCOPES permission which is only addable using a HSDP support ticket.
	// Removal protection of some ticket only permissions.
	TicketProtection *bool `json:"ticketProtection,omitempty" tf:"ticket_protection,omitempty"`
}

type RoleParameters struct {

	// The description of the group
	// The role description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The managing organization ID of this role
	// The managing organization of the role.
	// +crossplane:generate:reference:type=github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +kubebuilder:validation:Optional
	ManagingOrganization *string `json:"managingOrganization,omitempty" tf:"managing_organization,omitempty"`

	// Selector for a Organization in iam to populate managingOrganization.
	// +kubebuilder:validation:Optional
	ManagingOrganizationSelector *v1.Selector `json:"managingOrganizationSelector,omitempty" tf:"-"`

	// The name of the group
	// The role name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Organization in iam to populate managingOrganization.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// The list of permission to assign to this role
	// List of permissions IDs assigned to this role.
	// +kubebuilder:validation:Optional
	// +listType=set
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Defaults to true. Setting to false will remove e.g. CLIENT.SCOPES permission which is only addable using a HSDP support ticket.
	// Removal protection of some ticket only permissions.
	// +kubebuilder:validation:Optional
	TicketProtection *bool `json:"ticketProtection,omitempty" tf:"ticket_protection,omitempty"`
}

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleParameters `json:"forProvider"`
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
	InitProvider RoleInitParameters `json:"initProvider,omitempty"`
}

// RoleStatus defines the observed state of Role.
type RoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Role is the Schema for the Roles API. Manages HSDP IAM Role resources
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permissions) || (has(self.initProvider) && has(self.initProvider.permissions))",message="spec.forProvider.permissions is a required parameter"
	Spec   RoleSpec   `json:"spec"`
	Status RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Roles
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// Repository type metadata.
var (
	Role_Kind             = "Role"
	Role_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Role_Kind}.String()
	Role_KindAPIVersion   = Role_Kind + "." + CRDGroupVersion.String()
	Role_GroupVersionKind = CRDGroupVersion.WithKind(Role_Kind)
)

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}
