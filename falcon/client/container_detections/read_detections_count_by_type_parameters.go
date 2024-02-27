// Code generated by go-swagger; DO NOT EDIT.

package container_detections

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
)

// NewReadDetectionsCountByTypeParams creates a new ReadDetectionsCountByTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadDetectionsCountByTypeParams() *ReadDetectionsCountByTypeParams {
	return &ReadDetectionsCountByTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadDetectionsCountByTypeParamsWithTimeout creates a new ReadDetectionsCountByTypeParams object
// with the ability to set a timeout on a request.
func NewReadDetectionsCountByTypeParamsWithTimeout(timeout time.Duration) *ReadDetectionsCountByTypeParams {
	return &ReadDetectionsCountByTypeParams{
		timeout: timeout,
	}
}

// NewReadDetectionsCountByTypeParamsWithContext creates a new ReadDetectionsCountByTypeParams object
// with the ability to set a context for a request.
func NewReadDetectionsCountByTypeParamsWithContext(ctx context.Context) *ReadDetectionsCountByTypeParams {
	return &ReadDetectionsCountByTypeParams{
		Context: ctx,
	}
}

// NewReadDetectionsCountByTypeParamsWithHTTPClient creates a new ReadDetectionsCountByTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadDetectionsCountByTypeParamsWithHTTPClient(client *http.Client) *ReadDetectionsCountByTypeParams {
	return &ReadDetectionsCountByTypeParams{
		HTTPClient: client,
	}
}

/*
ReadDetectionsCountByTypeParams contains all the parameters to send to the API endpoint

	for the read detections count by type operation.

	Typically these are written to a http.Request.
*/
type ReadDetectionsCountByTypeParams struct {

	/* Filter.

	   Filter images using a query in Falcon Query Language (FQL). Supported filters:  cid,container_id,detection_type,id,image_digest,image_id,image_registry,image_repository,image_tag,name,severity
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read detections count by type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadDetectionsCountByTypeParams) WithDefaults() *ReadDetectionsCountByTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read detections count by type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadDetectionsCountByTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read detections count by type params
func (o *ReadDetectionsCountByTypeParams) WithTimeout(timeout time.Duration) *ReadDetectionsCountByTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read detections count by type params
func (o *ReadDetectionsCountByTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read detections count by type params
func (o *ReadDetectionsCountByTypeParams) WithContext(ctx context.Context) *ReadDetectionsCountByTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read detections count by type params
func (o *ReadDetectionsCountByTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read detections count by type params
func (o *ReadDetectionsCountByTypeParams) WithHTTPClient(client *http.Client) *ReadDetectionsCountByTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read detections count by type params
func (o *ReadDetectionsCountByTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the read detections count by type params
func (o *ReadDetectionsCountByTypeParams) WithFilter(filter *string) *ReadDetectionsCountByTypeParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the read detections count by type params
func (o *ReadDetectionsCountByTypeParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *ReadDetectionsCountByTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
