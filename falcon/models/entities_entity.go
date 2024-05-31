// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EntitiesEntity entities entity
//
// swagger:model entities.Entity
type EntitiesEntity struct {

	// ad domain
	AdDomain string `json:"ad_domain,omitempty"`

	// archived
	// Required: true
	Archived *bool `json:"archived"`

	// attack paths
	// Required: true
	AttackPaths []*EntitiesAttackPath `json:"attack_paths"`

	// attributes
	// Required: true
	Attributes []string `json:"attributes"`

	// authorizers
	// Required: true
	Authorizers []string `json:"authorizers"`

	// azure containing roles
	// Required: true
	AzureContainingRoles []string `json:"azure_containing_roles"`

	// azure rbac subscription names
	// Required: true
	AzureRbacSubscriptionNames []string `json:"azure_rbac_subscription_names"`

	// azure rbac subscription roles
	// Required: true
	AzureRbacSubscriptionRoles []string `json:"azure_rbac_subscription_roles"`

	// business privileges
	// Required: true
	BusinessPrivileges []string `json:"business_privileges"`

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// classifications
	// Required: true
	Classifications []string `json:"classifications"`

	// container type
	ContainerType string `json:"container_type,omitempty"`

	// creation time
	// Required: true
	CreationTime *string `json:"creation_time"`

	// department
	Department string `json:"department,omitempty"`

	// email addresses
	// Required: true
	EmailAddresses []string `json:"email_addresses"`

	// entity id
	// Required: true
	EntityID *string `json:"entity_id"`

	// entra domain
	EntraDomain string `json:"entra_domain,omitempty"`

	// entra tenant
	EntraTenant string `json:"entra_tenant,omitempty"`

	// flattened containing group entities
	// Required: true
	FlattenedContainingGroupEntities []string `json:"flattened_containing_group_entities"`

	// last update time
	// Required: true
	LastUpdateTime *string `json:"last_update_time"`

	// most recent activity
	MostRecentActivity string `json:"most_recent_activity,omitempty"`

	// most recent on premise activity
	MostRecentOnPremiseActivity string `json:"most_recent_on_premise_activity,omitempty"`

	// most recent sso activity
	MostRecentSsoActivity string `json:"most_recent_sso_activity,omitempty"`

	// object sid
	ObjectSid string `json:"object_sid,omitempty"`

	// ou
	Ou string `json:"ou,omitempty"`

	// password last change
	PasswordLastChange string `json:"password_last_change,omitempty"`

	// primary display name
	// Required: true
	PrimaryDisplayName *string `json:"primary_display_name"`

	// privileges
	// Required: true
	Privileges []string `json:"privileges"`

	// risk factors
	// Required: true
	RiskFactors []string `json:"risk_factors"`

	// risk score
	// Required: true
	RiskScore *float64 `json:"risk_score"`

	// secondary display name
	// Required: true
	SecondaryDisplayName *string `json:"secondary_display_name"`

	// stealth privileges paths
	// Required: true
	StealthPrivilegesPaths []*EntitiesAttackPath `json:"stealth_privileges_paths"`

	// title
	Title string `json:"title,omitempty"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this entities entity
func (m *EntitiesEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchived(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttackPaths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorizers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureContainingRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureRbacSubscriptionNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureRbacSubscriptionRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusinessPrivileges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClassifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlattenedContainingGroupEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivileges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRiskFactors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRiskScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStealthPrivilegesPaths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntitiesEntity) validateArchived(formats strfmt.Registry) error {

	if err := validate.Required("archived", "body", m.Archived); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateAttackPaths(formats strfmt.Registry) error {

	if err := validate.Required("attack_paths", "body", m.AttackPaths); err != nil {
		return err
	}

	for i := 0; i < len(m.AttackPaths); i++ {
		if swag.IsZero(m.AttackPaths[i]) { // not required
			continue
		}

		if m.AttackPaths[i] != nil {
			if err := m.AttackPaths[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attack_paths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attack_paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EntitiesEntity) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateAuthorizers(formats strfmt.Registry) error {

	if err := validate.Required("authorizers", "body", m.Authorizers); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateAzureContainingRoles(formats strfmt.Registry) error {

	if err := validate.Required("azure_containing_roles", "body", m.AzureContainingRoles); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateAzureRbacSubscriptionNames(formats strfmt.Registry) error {

	if err := validate.Required("azure_rbac_subscription_names", "body", m.AzureRbacSubscriptionNames); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateAzureRbacSubscriptionRoles(formats strfmt.Registry) error {

	if err := validate.Required("azure_rbac_subscription_roles", "body", m.AzureRbacSubscriptionRoles); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateBusinessPrivileges(formats strfmt.Registry) error {

	if err := validate.Required("business_privileges", "body", m.BusinessPrivileges); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateClassifications(formats strfmt.Registry) error {

	if err := validate.Required("classifications", "body", m.Classifications); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateCreationTime(formats strfmt.Registry) error {

	if err := validate.Required("creation_time", "body", m.CreationTime); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateEmailAddresses(formats strfmt.Registry) error {

	if err := validate.Required("email_addresses", "body", m.EmailAddresses); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateEntityID(formats strfmt.Registry) error {

	if err := validate.Required("entity_id", "body", m.EntityID); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateFlattenedContainingGroupEntities(formats strfmt.Registry) error {

	if err := validate.Required("flattened_containing_group_entities", "body", m.FlattenedContainingGroupEntities); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateLastUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("last_update_time", "body", m.LastUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validatePrimaryDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("primary_display_name", "body", m.PrimaryDisplayName); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validatePrivileges(formats strfmt.Registry) error {

	if err := validate.Required("privileges", "body", m.Privileges); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateRiskFactors(formats strfmt.Registry) error {

	if err := validate.Required("risk_factors", "body", m.RiskFactors); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateRiskScore(formats strfmt.Registry) error {

	if err := validate.Required("risk_score", "body", m.RiskScore); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateSecondaryDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("secondary_display_name", "body", m.SecondaryDisplayName); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesEntity) validateStealthPrivilegesPaths(formats strfmt.Registry) error {

	if err := validate.Required("stealth_privileges_paths", "body", m.StealthPrivilegesPaths); err != nil {
		return err
	}

	for i := 0; i < len(m.StealthPrivilegesPaths); i++ {
		if swag.IsZero(m.StealthPrivilegesPaths[i]) { // not required
			continue
		}

		if m.StealthPrivilegesPaths[i] != nil {
			if err := m.StealthPrivilegesPaths[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stealth_privileges_paths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stealth_privileges_paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EntitiesEntity) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this entities entity based on the context it is used
func (m *EntitiesEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttackPaths(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStealthPrivilegesPaths(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntitiesEntity) contextValidateAttackPaths(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AttackPaths); i++ {

		if m.AttackPaths[i] != nil {

			if swag.IsZero(m.AttackPaths[i]) { // not required
				return nil
			}

			if err := m.AttackPaths[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attack_paths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attack_paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EntitiesEntity) contextValidateStealthPrivilegesPaths(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StealthPrivilegesPaths); i++ {

		if m.StealthPrivilegesPaths[i] != nil {

			if swag.IsZero(m.StealthPrivilegesPaths[i]) { // not required
				return nil
			}

			if err := m.StealthPrivilegesPaths[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stealth_privileges_paths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stealth_privileges_paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntitiesEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntitiesEntity) UnmarshalBinary(b []byte) error {
	var res EntitiesEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
