// Code generated by go-swagger; DO NOT EDIT.

package runtime_detections

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

// NewGetRuntimeDetectionsCombinedV2Params creates a new GetRuntimeDetectionsCombinedV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRuntimeDetectionsCombinedV2Params() *GetRuntimeDetectionsCombinedV2Params {
	return &GetRuntimeDetectionsCombinedV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRuntimeDetectionsCombinedV2ParamsWithTimeout creates a new GetRuntimeDetectionsCombinedV2Params object
// with the ability to set a timeout on a request.
func NewGetRuntimeDetectionsCombinedV2ParamsWithTimeout(timeout time.Duration) *GetRuntimeDetectionsCombinedV2Params {
	return &GetRuntimeDetectionsCombinedV2Params{
		timeout: timeout,
	}
}

// NewGetRuntimeDetectionsCombinedV2ParamsWithContext creates a new GetRuntimeDetectionsCombinedV2Params object
// with the ability to set a context for a request.
func NewGetRuntimeDetectionsCombinedV2ParamsWithContext(ctx context.Context) *GetRuntimeDetectionsCombinedV2Params {
	return &GetRuntimeDetectionsCombinedV2Params{
		Context: ctx,
	}
}

// NewGetRuntimeDetectionsCombinedV2ParamsWithHTTPClient creates a new GetRuntimeDetectionsCombinedV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetRuntimeDetectionsCombinedV2ParamsWithHTTPClient(client *http.Client) *GetRuntimeDetectionsCombinedV2Params {
	return &GetRuntimeDetectionsCombinedV2Params{
		HTTPClient: client,
	}
}

/*
GetRuntimeDetectionsCombinedV2Params contains all the parameters to send to the API endpoint

	for the get runtime detections combined v2 operation.

	Typically these are written to a http.Request.
*/
type GetRuntimeDetectionsCombinedV2Params struct {

	/* Filter.

	   Filter Container Runtime Detections using a query in Falcon Query Language (FQL). Supported filters:  action_taken,aid,cid,cloud,cluster_name,command_line,computer_name,container_id,detect_timestamp,detection_description,detection_id,file_name,file_path,host_id,host_type,image_id,name,namespace,pod_name,severity,tactic
	*/
	Filter *string

	/* Limit.

	   The upper-bound on the number of records to retrieve.
	*/
	Limit *int64

	/* Offset.

	   The offset from where to begin.
	*/
	Offset *int64

	/* Sort.

	   The field to sort the records on.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get runtime detections combined v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuntimeDetectionsCombinedV2Params) WithDefaults() *GetRuntimeDetectionsCombinedV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get runtime detections combined v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuntimeDetectionsCombinedV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) WithTimeout(timeout time.Duration) *GetRuntimeDetectionsCombinedV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) WithContext(ctx context.Context) *GetRuntimeDetectionsCombinedV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) WithHTTPClient(client *http.Client) *GetRuntimeDetectionsCombinedV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) WithFilter(filter *string) *GetRuntimeDetectionsCombinedV2Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) WithLimit(limit *int64) *GetRuntimeDetectionsCombinedV2Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) WithOffset(offset *int64) *GetRuntimeDetectionsCombinedV2Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) WithSort(sort *string) *GetRuntimeDetectionsCombinedV2Params {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get runtime detections combined v2 params
func (o *GetRuntimeDetectionsCombinedV2Params) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetRuntimeDetectionsCombinedV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
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
