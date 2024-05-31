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

// DomainExecutionMetadataSummary domain execution metadata summary
//
// swagger:model domain.ExecutionMetadataSummary
type DomainExecutionMetadataSummary struct {

	// report params
	// Required: true
	ReportParams *DomainReportParams `json:"report_params"`

	// subtype
	// Required: true
	Subtype *string `json:"subtype"`

	// unscheduled execution type
	// Required: true
	UnscheduledExecutionType *string `json:"unscheduled_execution_type"`

	// xdr data
	// Required: true
	XdrData *DomainXDRData `json:"xdr_data"`

	// xdr params
	// Required: true
	XdrParams *DomainXDRParams `json:"xdr_params"`
}

// Validate validates this domain execution metadata summary
func (m *DomainExecutionMetadataSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReportParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubtype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnscheduledExecutionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXdrData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXdrParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainExecutionMetadataSummary) validateReportParams(formats strfmt.Registry) error {

	if err := validate.Required("report_params", "body", m.ReportParams); err != nil {
		return err
	}

	if m.ReportParams != nil {
		if err := m.ReportParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("report_params")
			}
			return err
		}
	}

	return nil
}

func (m *DomainExecutionMetadataSummary) validateSubtype(formats strfmt.Registry) error {

	if err := validate.Required("subtype", "body", m.Subtype); err != nil {
		return err
	}

	return nil
}

func (m *DomainExecutionMetadataSummary) validateUnscheduledExecutionType(formats strfmt.Registry) error {

	if err := validate.Required("unscheduled_execution_type", "body", m.UnscheduledExecutionType); err != nil {
		return err
	}

	return nil
}

func (m *DomainExecutionMetadataSummary) validateXdrData(formats strfmt.Registry) error {

	if err := validate.Required("xdr_data", "body", m.XdrData); err != nil {
		return err
	}

	if m.XdrData != nil {
		if err := m.XdrData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("xdr_data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("xdr_data")
			}
			return err
		}
	}

	return nil
}

func (m *DomainExecutionMetadataSummary) validateXdrParams(formats strfmt.Registry) error {

	if err := validate.Required("xdr_params", "body", m.XdrParams); err != nil {
		return err
	}

	if m.XdrParams != nil {
		if err := m.XdrParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("xdr_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("xdr_params")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this domain execution metadata summary based on the context it is used
func (m *DomainExecutionMetadataSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReportParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateXdrData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateXdrParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainExecutionMetadataSummary) contextValidateReportParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ReportParams != nil {

		if err := m.ReportParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("report_params")
			}
			return err
		}
	}

	return nil
}

func (m *DomainExecutionMetadataSummary) contextValidateXdrData(ctx context.Context, formats strfmt.Registry) error {

	if m.XdrData != nil {

		if err := m.XdrData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("xdr_data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("xdr_data")
			}
			return err
		}
	}

	return nil
}

func (m *DomainExecutionMetadataSummary) contextValidateXdrParams(ctx context.Context, formats strfmt.Registry) error {

	if m.XdrParams != nil {

		if err := m.XdrParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("xdr_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("xdr_params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainExecutionMetadataSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainExecutionMetadataSummary) UnmarshalBinary(b []byte) error {
	var res DomainExecutionMetadataSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
