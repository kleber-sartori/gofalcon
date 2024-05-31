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

// FwmgrMsaAggregationResult fwmgr msa aggregation result
//
// swagger:model fwmgr.msa.AggregationResult
type FwmgrMsaAggregationResult struct {

	// buckets
	Buckets []*FwmgrMsaAggregationResultItem `json:"buckets"`

	// doc count error upper bound
	DocCountErrorUpperBound int64 `json:"doc_count_error_upper_bound,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// sum other doc count
	SumOtherDocCount int64 `json:"sum_other_doc_count,omitempty"`
}

// Validate validates this fwmgr msa aggregation result
func (m *FwmgrMsaAggregationResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuckets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrMsaAggregationResult) validateBuckets(formats strfmt.Registry) error {
	if swag.IsZero(m.Buckets) { // not required
		return nil
	}

	for i := 0; i < len(m.Buckets); i++ {
		if swag.IsZero(m.Buckets[i]) { // not required
			continue
		}

		if m.Buckets[i] != nil {
			if err := m.Buckets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buckets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("buckets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FwmgrMsaAggregationResult) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fwmgr msa aggregation result based on the context it is used
func (m *FwmgrMsaAggregationResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuckets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrMsaAggregationResult) contextValidateBuckets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Buckets); i++ {

		if m.Buckets[i] != nil {

			if swag.IsZero(m.Buckets[i]) { // not required
				return nil
			}

			if err := m.Buckets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buckets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("buckets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrMsaAggregationResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrMsaAggregationResult) UnmarshalBinary(b []byte) error {
	var res FwmgrMsaAggregationResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
