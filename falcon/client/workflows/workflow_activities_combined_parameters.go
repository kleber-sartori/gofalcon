// Code generated by go-swagger; DO NOT EDIT.

package workflows

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

// NewWorkflowActivitiesCombinedParams creates a new WorkflowActivitiesCombinedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowActivitiesCombinedParams() *WorkflowActivitiesCombinedParams {
	return &WorkflowActivitiesCombinedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowActivitiesCombinedParamsWithTimeout creates a new WorkflowActivitiesCombinedParams object
// with the ability to set a timeout on a request.
func NewWorkflowActivitiesCombinedParamsWithTimeout(timeout time.Duration) *WorkflowActivitiesCombinedParams {
	return &WorkflowActivitiesCombinedParams{
		timeout: timeout,
	}
}

// NewWorkflowActivitiesCombinedParamsWithContext creates a new WorkflowActivitiesCombinedParams object
// with the ability to set a context for a request.
func NewWorkflowActivitiesCombinedParamsWithContext(ctx context.Context) *WorkflowActivitiesCombinedParams {
	return &WorkflowActivitiesCombinedParams{
		Context: ctx,
	}
}

// NewWorkflowActivitiesCombinedParamsWithHTTPClient creates a new WorkflowActivitiesCombinedParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowActivitiesCombinedParamsWithHTTPClient(client *http.Client) *WorkflowActivitiesCombinedParams {
	return &WorkflowActivitiesCombinedParams{
		HTTPClient: client,
	}
}

/*
WorkflowActivitiesCombinedParams contains all the parameters to send to the API endpoint

	for the workflow activities combined operation.

	Typically these are written to a http.Request.
*/
type WorkflowActivitiesCombinedParams struct {

	/* Filter.

	   FQL query specifying filter parameters.
	*/
	Filter string

	/* Limit.

	   Maximum number of records to return.
	*/
	Limit *int64

	/* Offset.

	   Starting pagination offset of records to return.
	*/
	Offset *string

	/* Sort.

	   Sort items by providing a comma separated list of property and direction (eg name.desc,time.asc). If direction is omitted, defaults to descending.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workflow activities combined params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowActivitiesCombinedParams) WithDefaults() *WorkflowActivitiesCombinedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow activities combined params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowActivitiesCombinedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) WithTimeout(timeout time.Duration) *WorkflowActivitiesCombinedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) WithContext(ctx context.Context) *WorkflowActivitiesCombinedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) WithHTTPClient(client *http.Client) *WorkflowActivitiesCombinedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) WithFilter(filter string) *WorkflowActivitiesCombinedParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) SetFilter(filter string) {
	o.Filter = filter
}

// WithLimit adds the limit to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) WithLimit(limit *int64) *WorkflowActivitiesCombinedParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) WithOffset(offset *string) *WorkflowActivitiesCombinedParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithSort adds the sort to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) WithSort(sort *string) *WorkflowActivitiesCombinedParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the workflow activities combined params
func (o *WorkflowActivitiesCombinedParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowActivitiesCombinedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param filter
	qrFilter := o.Filter
	qFilter := qrFilter

	if err := r.SetQueryParam("filter", qFilter); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
