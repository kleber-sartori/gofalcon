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

// GraphDefinitionModel graph definition model
//
// swagger:model graph.DefinitionModel
type GraphDefinitionModel struct {

	// Details of all the activities within the model. Each activity has a unique node ID as the key, that is set by the caller.
	Activities map[string]GraphConfiguredActivity `json:"activities,omitempty"`

	// The end event which contains details to all inbound flow references.
	End *GraphEnd `json:"end,omitempty"`

	// Details all the sequence flows within the model. Each flow has a unique node ID as the key, that is set by the caller.
	// Required: true
	Flows map[string]GraphFlow `json:"flows"`

	// Details all the gateways within the model. Each gateway has a unique node ID as the key, that is set by the caller.
	Gateways map[string]GraphGateway `json:"gateways,omitempty"`

	// node registry
	// Required: true
	NodeRegistry interface{} `json:"nodeRegistry"`

	// Details of all sub-models within the model. Each sub-model has a unique node ID as the key set by the caller.
	SubModels map[string]GraphSubModel `json:"sub_models,omitempty"`

	// Selected trigger to invoke the workflow.
	// Required: true
	Trigger *GraphConfiguredTrigger `json:"trigger"`
}

// Validate validates this graph definition model
func (m *GraphDefinitionModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlows(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateways(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubModels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphDefinitionModel) validateActivities(formats strfmt.Registry) error {
	if swag.IsZero(m.Activities) { // not required
		return nil
	}

	for k := range m.Activities {

		if err := validate.Required("activities"+"."+k, "body", m.Activities[k]); err != nil {
			return err
		}
		if val, ok := m.Activities[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activities" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("activities" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *GraphDefinitionModel) validateEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.End) { // not required
		return nil
	}

	if m.End != nil {
		if err := m.End.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("end")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("end")
			}
			return err
		}
	}

	return nil
}

func (m *GraphDefinitionModel) validateFlows(formats strfmt.Registry) error {

	if err := validate.Required("flows", "body", m.Flows); err != nil {
		return err
	}

	for k := range m.Flows {

		if err := validate.Required("flows"+"."+k, "body", m.Flows[k]); err != nil {
			return err
		}
		if val, ok := m.Flows[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flows" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("flows" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *GraphDefinitionModel) validateGateways(formats strfmt.Registry) error {
	if swag.IsZero(m.Gateways) { // not required
		return nil
	}

	for k := range m.Gateways {

		if err := validate.Required("gateways"+"."+k, "body", m.Gateways[k]); err != nil {
			return err
		}
		if val, ok := m.Gateways[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gateways" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gateways" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *GraphDefinitionModel) validateNodeRegistry(formats strfmt.Registry) error {

	if m.NodeRegistry == nil {
		return errors.Required("nodeRegistry", "body", nil)
	}

	return nil
}

func (m *GraphDefinitionModel) validateSubModels(formats strfmt.Registry) error {
	if swag.IsZero(m.SubModels) { // not required
		return nil
	}

	for k := range m.SubModels {

		if err := validate.Required("sub_models"+"."+k, "body", m.SubModels[k]); err != nil {
			return err
		}
		if val, ok := m.SubModels[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sub_models" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sub_models" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *GraphDefinitionModel) validateTrigger(formats strfmt.Registry) error {

	if err := validate.Required("trigger", "body", m.Trigger); err != nil {
		return err
	}

	if m.Trigger != nil {
		if err := m.Trigger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this graph definition model based on the context it is used
func (m *GraphDefinitionModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActivities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGateways(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubModels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrigger(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphDefinitionModel) contextValidateActivities(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Activities {

		if val, ok := m.Activities[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *GraphDefinitionModel) contextValidateEnd(ctx context.Context, formats strfmt.Registry) error {

	if m.End != nil {

		if swag.IsZero(m.End) { // not required
			return nil
		}

		if err := m.End.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("end")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("end")
			}
			return err
		}
	}

	return nil
}

func (m *GraphDefinitionModel) contextValidateFlows(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("flows", "body", m.Flows); err != nil {
		return err
	}

	for k := range m.Flows {

		if val, ok := m.Flows[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *GraphDefinitionModel) contextValidateGateways(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Gateways {

		if val, ok := m.Gateways[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *GraphDefinitionModel) contextValidateSubModels(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.SubModels {

		if val, ok := m.SubModels[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *GraphDefinitionModel) contextValidateTrigger(ctx context.Context, formats strfmt.Registry) error {

	if m.Trigger != nil {

		if err := m.Trigger.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GraphDefinitionModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphDefinitionModel) UnmarshalBinary(b []byte) error {
	var res GraphDefinitionModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
