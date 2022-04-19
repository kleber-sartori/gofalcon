// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new hosts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hosts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDeviceDetails(params *GetDeviceDetailsParams, opts ...ClientOption) (*GetDeviceDetailsOK, error)

	GetOnlineStateV1(params *GetOnlineStateV1Params, opts ...ClientOption) (*GetOnlineStateV1OK, error)

	PerformActionV2(params *PerformActionV2Params, opts ...ClientOption) (*PerformActionV2Accepted, error)

	QueryDeviceLoginHistory(params *QueryDeviceLoginHistoryParams, opts ...ClientOption) (*QueryDeviceLoginHistoryOK, error)

	QueryDevicesByFilter(params *QueryDevicesByFilterParams, opts ...ClientOption) (*QueryDevicesByFilterOK, error)

	QueryDevicesByFilterScroll(params *QueryDevicesByFilterScrollParams, opts ...ClientOption) (*QueryDevicesByFilterScrollOK, error)

	QueryGetNetworkAddressHistoryV1(params *QueryGetNetworkAddressHistoryV1Params, opts ...ClientOption) (*QueryGetNetworkAddressHistoryV1OK, error)

	QueryHiddenDevices(params *QueryHiddenDevicesParams, opts ...ClientOption) (*QueryHiddenDevicesOK, error)

	UpdateDeviceTags(params *UpdateDeviceTagsParams, opts ...ClientOption) (*UpdateDeviceTagsOK, error)

	EntitiesPerformAction(params *EntitiesPerformActionParams, opts ...ClientOption) (*EntitiesPerformActionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDeviceDetails gets details on one or more hosts by providing agent i ds a ID you can get a host s agent i ds a i ds from the devices queries devices v1 endpoint the falcon console or the streaming API
*/
func (a *Client) GetDeviceDetails(params *GetDeviceDetailsParams, opts ...ClientOption) (*GetDeviceDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeviceDetails",
		Method:             "GET",
		PathPattern:        "/devices/entities/devices/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceDetailsReader{formats: a.formats},
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
	success, ok := result.(*GetDeviceDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDeviceDetailsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetOnlineStateV1 gets the online status for one or more hosts by specifying each host s unique ID successful requests return an HTTP 200 response and the status for each host identified by a state of online offline or unknown for each host identified by host id make a g e t request to devices queries devices v1 to get a list of host i ds
*/
func (a *Client) GetOnlineStateV1(params *GetOnlineStateV1Params, opts ...ClientOption) (*GetOnlineStateV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOnlineStateV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOnlineState.V1",
		Method:             "GET",
		PathPattern:        "/devices/entities/online-state/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOnlineStateV1Reader{formats: a.formats},
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
	success, ok := result.(*GetOnlineStateV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOnlineStateV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PerformActionV2 takes various actions on the hosts in your environment contain or lift containment on a host delete or restore a host
*/
func (a *Client) PerformActionV2(params *PerformActionV2Params, opts ...ClientOption) (*PerformActionV2Accepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformActionV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PerformActionV2",
		Method:             "POST",
		PathPattern:        "/devices/entities/devices-actions/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PerformActionV2Reader{formats: a.formats},
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
	success, ok := result.(*PerformActionV2Accepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PerformActionV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QueryDeviceLoginHistory retrieves details about recent login sessions for a set of devices
*/
func (a *Client) QueryDeviceLoginHistory(params *QueryDeviceLoginHistoryParams, opts ...ClientOption) (*QueryDeviceLoginHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDeviceLoginHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryDeviceLoginHistory",
		Method:             "POST",
		PathPattern:        "/devices/combined/devices/login-history/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryDeviceLoginHistoryReader{formats: a.formats},
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
	success, ok := result.(*QueryDeviceLoginHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDeviceLoginHistoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryDevicesByFilter searches for hosts in your environment by platform hostname IP and other criteria
*/
func (a *Client) QueryDevicesByFilter(params *QueryDevicesByFilterParams, opts ...ClientOption) (*QueryDevicesByFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDevicesByFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryDevicesByFilter",
		Method:             "GET",
		PathPattern:        "/devices/queries/devices/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryDevicesByFilterReader{formats: a.formats},
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
	success, ok := result.(*QueryDevicesByFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDevicesByFilterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryDevicesByFilterScroll searches for hosts in your environment by platform hostname IP and other criteria with continuous pagination capability based on offset pointer which expires after 2 minutes with no maximum limit
*/
func (a *Client) QueryDevicesByFilterScroll(params *QueryDevicesByFilterScrollParams, opts ...ClientOption) (*QueryDevicesByFilterScrollOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDevicesByFilterScrollParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryDevicesByFilterScroll",
		Method:             "GET",
		PathPattern:        "/devices/queries/devices-scroll/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryDevicesByFilterScrollReader{formats: a.formats},
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
	success, ok := result.(*QueryDevicesByFilterScrollOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDevicesByFilterScrollDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryGetNetworkAddressHistoryV1 retrieves history of IP and m a c addresses of devices
*/
func (a *Client) QueryGetNetworkAddressHistoryV1(params *QueryGetNetworkAddressHistoryV1Params, opts ...ClientOption) (*QueryGetNetworkAddressHistoryV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryGetNetworkAddressHistoryV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryGetNetworkAddressHistoryV1",
		Method:             "POST",
		PathPattern:        "/devices/combined/devices/network-address-history/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryGetNetworkAddressHistoryV1Reader{formats: a.formats},
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
	success, ok := result.(*QueryGetNetworkAddressHistoryV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryGetNetworkAddressHistoryV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryHiddenDevices retrieves hidden hosts that match the provided filter criteria
*/
func (a *Client) QueryHiddenDevices(params *QueryHiddenDevicesParams, opts ...ClientOption) (*QueryHiddenDevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryHiddenDevicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryHiddenDevices",
		Method:             "GET",
		PathPattern:        "/devices/queries/devices-hidden/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryHiddenDevicesReader{formats: a.formats},
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
	success, ok := result.(*QueryHiddenDevicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryHiddenDevicesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateDeviceTags appends or remove one or more falcon grouping tags on one or more hosts
*/
func (a *Client) UpdateDeviceTags(params *UpdateDeviceTagsParams, opts ...ClientOption) (*UpdateDeviceTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeviceTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateDeviceTags",
		Method:             "PATCH",
		PathPattern:        "/devices/entities/devices/tags/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDeviceTagsReader{formats: a.formats},
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
	success, ok := result.(*UpdateDeviceTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateDeviceTagsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EntitiesPerformAction performs the specified action on the provided prevention policy i ds
*/
func (a *Client) EntitiesPerformAction(params *EntitiesPerformActionParams, opts ...ClientOption) (*EntitiesPerformActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEntitiesPerformActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "entities.perform_action",
		Method:             "POST",
		PathPattern:        "/devices/entities/group-actions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EntitiesPerformActionReader{formats: a.formats},
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
	success, ok := result.(*EntitiesPerformActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EntitiesPerformActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
