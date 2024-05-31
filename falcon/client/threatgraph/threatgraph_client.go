// Code generated by go-swagger; DO NOT EDIT.

package threatgraph

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new threatgraph API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for threatgraph API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CombinedEdgesGet(params *CombinedEdgesGetParams, opts ...ClientOption) (*CombinedEdgesGetOK, error)

	CombinedRanOnGet(params *CombinedRanOnGetParams, opts ...ClientOption) (*CombinedRanOnGetOK, error)

	CombinedSummaryGet(params *CombinedSummaryGetParams, opts ...ClientOption) (*CombinedSummaryGetOK, error)

	EntitiesVerticesGet(params *EntitiesVerticesGetParams, opts ...ClientOption) (*EntitiesVerticesGetOK, error)

	EntitiesVerticesGetv2(params *EntitiesVerticesGetv2Params, opts ...ClientOption) (*EntitiesVerticesGetv2OK, error)

	QueriesEdgetypesGet(params *QueriesEdgetypesGetParams, opts ...ClientOption) (*QueriesEdgetypesGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CombinedEdgesGet retrieves edges for a given vertex id one edge type must be specified
*/
func (a *Client) CombinedEdgesGet(params *CombinedEdgesGetParams, opts ...ClientOption) (*CombinedEdgesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCombinedEdgesGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "combined_edges_get",
		Method:             "GET",
		PathPattern:        "/threatgraph/combined/edges/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CombinedEdgesGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CombinedEdgesGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for combined_edges_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CombinedRanOnGet looks up instances of indicators such as hashes domain names and ip addresses that have been seen on devices in your environment
*/
func (a *Client) CombinedRanOnGet(params *CombinedRanOnGetParams, opts ...ClientOption) (*CombinedRanOnGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCombinedRanOnGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "combined_ran_on_get",
		Method:             "GET",
		PathPattern:        "/threatgraph/combined/ran-on/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CombinedRanOnGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CombinedRanOnGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for combined_ran_on_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CombinedSummaryGet retrieves summary for a given vertex ID
*/
func (a *Client) CombinedSummaryGet(params *CombinedSummaryGetParams, opts ...ClientOption) (*CombinedSummaryGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCombinedSummaryGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "combined_summary_get",
		Method:             "GET",
		PathPattern:        "/threatgraph/combined/{vertex-type}/summary/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CombinedSummaryGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CombinedSummaryGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for combined_summary_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EntitiesVerticesGet retrieves metadata for a given vertex ID
*/
func (a *Client) EntitiesVerticesGet(params *EntitiesVerticesGetParams, opts ...ClientOption) (*EntitiesVerticesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEntitiesVerticesGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "entities_vertices_get",
		Method:             "GET",
		PathPattern:        "/threatgraph/entities/{vertex-type}/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EntitiesVerticesGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EntitiesVerticesGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for entities_vertices_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EntitiesVerticesGetv2 retrieves metadata for a given vertex ID
*/
func (a *Client) EntitiesVerticesGetv2(params *EntitiesVerticesGetv2Params, opts ...ClientOption) (*EntitiesVerticesGetv2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEntitiesVerticesGetv2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "entities_vertices_getv2",
		Method:             "GET",
		PathPattern:        "/threatgraph/entities/{vertex-type}/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EntitiesVerticesGetv2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EntitiesVerticesGetv2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for entities_vertices_getv2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueriesEdgetypesGet shows all available edge types
*/
func (a *Client) QueriesEdgetypesGet(params *QueriesEdgetypesGetParams, opts ...ClientOption) (*QueriesEdgetypesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueriesEdgetypesGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queries_edgetypes_get",
		Method:             "GET",
		PathPattern:        "/threatgraph/queries/edge-types/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueriesEdgetypesGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueriesEdgetypesGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queries_edgetypes_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
