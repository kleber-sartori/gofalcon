// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// NewNodesByContainerEngineVersionCountParams creates a new NodesByContainerEngineVersionCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodesByContainerEngineVersionCountParams() *NodesByContainerEngineVersionCountParams {
	return &NodesByContainerEngineVersionCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodesByContainerEngineVersionCountParamsWithTimeout creates a new NodesByContainerEngineVersionCountParams object
// with the ability to set a timeout on a request.
func NewNodesByContainerEngineVersionCountParamsWithTimeout(timeout time.Duration) *NodesByContainerEngineVersionCountParams {
	return &NodesByContainerEngineVersionCountParams{
		timeout: timeout,
	}
}

// NewNodesByContainerEngineVersionCountParamsWithContext creates a new NodesByContainerEngineVersionCountParams object
// with the ability to set a context for a request.
func NewNodesByContainerEngineVersionCountParamsWithContext(ctx context.Context) *NodesByContainerEngineVersionCountParams {
	return &NodesByContainerEngineVersionCountParams{
		Context: ctx,
	}
}

// NewNodesByContainerEngineVersionCountParamsWithHTTPClient creates a new NodesByContainerEngineVersionCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodesByContainerEngineVersionCountParamsWithHTTPClient(client *http.Client) *NodesByContainerEngineVersionCountParams {
	return &NodesByContainerEngineVersionCountParams{
		HTTPClient: client,
	}
}

/*
NodesByContainerEngineVersionCountParams contains all the parameters to send to the API endpoint

	for the nodes by container engine version count operation.

	Typically these are written to a http.Request.
*/
type NodesByContainerEngineVersionCountParams struct {

	/* Filter.

	   Search Kubernetes nodes using a query in Falcon Query Language (FQL). Supported filters:  aid,annotations_list,cid,cloud_account_id,cloud_name,cloud_region,cluster_id,cluster_name,container_count,container_runtime_version,first_seen,image_digest,ipv4,last_seen,linux_sensor_coverage,node_name,pod_count
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nodes by container engine version count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodesByContainerEngineVersionCountParams) WithDefaults() *NodesByContainerEngineVersionCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nodes by container engine version count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodesByContainerEngineVersionCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nodes by container engine version count params
func (o *NodesByContainerEngineVersionCountParams) WithTimeout(timeout time.Duration) *NodesByContainerEngineVersionCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodes by container engine version count params
func (o *NodesByContainerEngineVersionCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodes by container engine version count params
func (o *NodesByContainerEngineVersionCountParams) WithContext(ctx context.Context) *NodesByContainerEngineVersionCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodes by container engine version count params
func (o *NodesByContainerEngineVersionCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodes by container engine version count params
func (o *NodesByContainerEngineVersionCountParams) WithHTTPClient(client *http.Client) *NodesByContainerEngineVersionCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodes by container engine version count params
func (o *NodesByContainerEngineVersionCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the nodes by container engine version count params
func (o *NodesByContainerEngineVersionCountParams) WithFilter(filter *string) *NodesByContainerEngineVersionCountParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the nodes by container engine version count params
func (o *NodesByContainerEngineVersionCountParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *NodesByContainerEngineVersionCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
