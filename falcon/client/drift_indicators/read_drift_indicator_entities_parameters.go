// Code generated by go-swagger; DO NOT EDIT.

package drift_indicators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewReadDriftIndicatorEntitiesParams creates a new ReadDriftIndicatorEntitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadDriftIndicatorEntitiesParams() *ReadDriftIndicatorEntitiesParams {
	return &ReadDriftIndicatorEntitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadDriftIndicatorEntitiesParamsWithTimeout creates a new ReadDriftIndicatorEntitiesParams object
// with the ability to set a timeout on a request.
func NewReadDriftIndicatorEntitiesParamsWithTimeout(timeout time.Duration) *ReadDriftIndicatorEntitiesParams {
	return &ReadDriftIndicatorEntitiesParams{
		timeout: timeout,
	}
}

// NewReadDriftIndicatorEntitiesParamsWithContext creates a new ReadDriftIndicatorEntitiesParams object
// with the ability to set a context for a request.
func NewReadDriftIndicatorEntitiesParamsWithContext(ctx context.Context) *ReadDriftIndicatorEntitiesParams {
	return &ReadDriftIndicatorEntitiesParams{
		Context: ctx,
	}
}

// NewReadDriftIndicatorEntitiesParamsWithHTTPClient creates a new ReadDriftIndicatorEntitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadDriftIndicatorEntitiesParamsWithHTTPClient(client *http.Client) *ReadDriftIndicatorEntitiesParams {
	return &ReadDriftIndicatorEntitiesParams{
		HTTPClient: client,
	}
}

/*
ReadDriftIndicatorEntitiesParams contains all the parameters to send to the API endpoint

	for the read drift indicator entities operation.

	Typically these are written to a http.Request.
*/
type ReadDriftIndicatorEntitiesParams struct {

	/* Ids.

	   Search Drift Indicators by ids - The maximum amount is 100 IDs
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read drift indicator entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadDriftIndicatorEntitiesParams) WithDefaults() *ReadDriftIndicatorEntitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read drift indicator entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadDriftIndicatorEntitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read drift indicator entities params
func (o *ReadDriftIndicatorEntitiesParams) WithTimeout(timeout time.Duration) *ReadDriftIndicatorEntitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read drift indicator entities params
func (o *ReadDriftIndicatorEntitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read drift indicator entities params
func (o *ReadDriftIndicatorEntitiesParams) WithContext(ctx context.Context) *ReadDriftIndicatorEntitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read drift indicator entities params
func (o *ReadDriftIndicatorEntitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read drift indicator entities params
func (o *ReadDriftIndicatorEntitiesParams) WithHTTPClient(client *http.Client) *ReadDriftIndicatorEntitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read drift indicator entities params
func (o *ReadDriftIndicatorEntitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the read drift indicator entities params
func (o *ReadDriftIndicatorEntitiesParams) WithIds(ids []string) *ReadDriftIndicatorEntitiesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the read drift indicator entities params
func (o *ReadDriftIndicatorEntitiesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *ReadDriftIndicatorEntitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamReadDriftIndicatorEntities binds the parameter ids
func (o *ReadDriftIndicatorEntitiesParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "csv"
	idsIS := swag.JoinByFormat(idsIC, "csv")

	return idsIS
}
