// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainExecuteCommandV1 domain execute command v1
//
// swagger:model domain.ExecuteCommandV1
type DomainExecuteCommandV1 struct {

	// Config auth type for plugin to execute. Only applicable for oneOf security scheme plugins. If not provided, it will use the default auth type on the config
	// Required: true
	ConfigAuthType *string `json:"config_auth_type"`

	// ConfigID for plugin to execute. If omitted, the oldest config will be used as part of the execution.
	// Required: true
	ConfigID *string `json:"config_id"`

	// ID of the definition containing the operation to execute.'
	// Required: true
	DefinitionID *string `json:"definition_id"`

	// ID of the specific plugin to execute, in the format 'definition_name.operation_name'
	// Required: true
	ID *string `json:"id"`

	// The specific operation to execute.
	// Required: true
	OperationID *string `json:"operation_id"`

	// The request params/body to execute the command. The data model for this section is determined at runtime via the request schema.
	// Required: true
	Request *DomainRequest `json:"request"`

	// The version of the definition to execute.
	// Required: true
	Version *int32 `json:"version"`
}

// Validate validates this domain execute command v1
func (m *DomainExecuteCommandV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigAuthType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefinitionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainExecuteCommandV1) validateConfigAuthType(formats strfmt.Registry) error {

	if err := validate.Required("config_auth_type", "body", m.ConfigAuthType); err != nil {
		return err
	}

	return nil
}

func (m *DomainExecuteCommandV1) validateConfigID(formats strfmt.Registry) error {

	if err := validate.Required("config_id", "body", m.ConfigID); err != nil {
		return err
	}

	return nil
}

func (m *DomainExecuteCommandV1) validateDefinitionID(formats strfmt.Registry) error {

	if err := validate.Required("definition_id", "body", m.DefinitionID); err != nil {
		return err
	}

	return nil
}

func (m *DomainExecuteCommandV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainExecuteCommandV1) validateOperationID(formats strfmt.Registry) error {

	if err := validate.Required("operation_id", "body", m.OperationID); err != nil {
		return err
	}

	return nil
}

func (m *DomainExecuteCommandV1) validateRequest(formats strfmt.Registry) error {

	if err := validate.Required("request", "body", m.Request); err != nil {
		return err
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *DomainExecuteCommandV1) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain execute command v1 based on the context it is used
func (m *DomainExecuteCommandV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainExecuteCommandV1) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.Request != nil {

		if err := m.Request.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainExecuteCommandV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainExecuteCommandV1) UnmarshalBinary(b []byte) error {
	var res DomainExecuteCommandV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
