// Code generated by go-swagger; DO NOT EDIT.

package workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workflows API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workflows API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Deprovision(params *DeprovisionParams, opts ...ClientOption) (*DeprovisionOK, error)

	Execute(params *ExecuteParams, opts ...ClientOption) (*ExecuteOK, error)

	ExecutionAction(params *ExecutionActionParams, opts ...ClientOption) (*ExecutionActionOK, error)

	ExecutionResults(params *ExecutionResultsParams, opts ...ClientOption) (*ExecutionResultsOK, error)

	Promote(params *PromoteParams, opts ...ClientOption) (*PromoteOK, error)

	Provision(params *ProvisionParams, opts ...ClientOption) (*ProvisionOK, error)

	WorkflowDefinitionsCombined(params *WorkflowDefinitionsCombinedParams, opts ...ClientOption) (*WorkflowDefinitionsCombinedOK, error)

	WorkflowDefinitionsCreate(params *WorkflowDefinitionsCreateParams, opts ...ClientOption) (*WorkflowDefinitionsCreateOK, error)

	WorkflowDefinitionsExport(params *WorkflowDefinitionsExportParams, opts ...ClientOption) (*WorkflowDefinitionsExportOK, *WorkflowDefinitionsExportStatus299, error)

	WorkflowDefinitionsImport(params *WorkflowDefinitionsImportParams, opts ...ClientOption) (*WorkflowDefinitionsImportOK, error)

	WorkflowDefinitionsUpdate(params *WorkflowDefinitionsUpdateParams, opts ...ClientOption) (*WorkflowDefinitionsUpdateOK, error)

	WorkflowExecutionsCombined(params *WorkflowExecutionsCombinedParams, opts ...ClientOption) (*WorkflowExecutionsCombinedOK, error)

	WorkflowGetHumanInputV1(params *WorkflowGetHumanInputV1Params, opts ...ClientOption) (*WorkflowGetHumanInputV1OK, error)

	WorkflowUpdateHumanInputV1(params *WorkflowUpdateHumanInputV1Params, opts ...ClientOption) (*WorkflowUpdateHumanInputV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Deprovision deprovisions a system definition that was previously provisioned on the target c ID
*/
func (a *Client) Deprovision(params *DeprovisionParams, opts ...ClientOption) (*DeprovisionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeprovisionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Deprovision",
		Method:             "POST",
		PathPattern:        "/workflows/system-definitions/deprovision/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeprovisionReader{formats: a.formats},
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
	success, ok := result.(*DeprovisionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Deprovision: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Execute executes an on demand workflow the body is JSON used to trigger the execution the response the execution ID s
*/
func (a *Client) Execute(params *ExecuteParams, opts ...ClientOption) (*ExecuteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecuteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Execute",
		Method:             "POST",
		PathPattern:        "/workflows/entities/execute/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecuteReader{formats: a.formats},
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
	success, ok := result.(*ExecuteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Execute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExecutionAction allows a user to resume retry a failed workflow execution
*/
func (a *Client) ExecutionAction(params *ExecutionActionParams, opts ...ClientOption) (*ExecutionActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecutionActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExecutionAction",
		Method:             "POST",
		PathPattern:        "/workflows/entities/execution-actions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecutionActionReader{formats: a.formats},
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
	success, ok := result.(*ExecutionActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExecutionAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExecutionResults gets execution result of a given execution
*/
func (a *Client) ExecutionResults(params *ExecutionResultsParams, opts ...ClientOption) (*ExecutionResultsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecutionResultsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExecutionResults",
		Method:             "GET",
		PathPattern:        "/workflows/entities/execution-results/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecutionResultsReader{formats: a.formats},
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
	success, ok := result.(*ExecutionResultsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExecutionResults: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Promote promotes a version of a system definition for a customer the customer must already have been provisioned this allows the caller to apply an updated template version to a specific cid and expects all parameters to be supplied if the template supports multi instance the customer scope definition ID must be supplied to determine which customer workflow should be updated
*/
func (a *Client) Promote(params *PromoteParams, opts ...ClientOption) (*PromoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPromoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Promote",
		Method:             "POST",
		PathPattern:        "/workflows/system-definitions/promote/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PromoteReader{formats: a.formats},
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
	success, ok := result.(*PromoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Promote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Provision provisions a system definition onto the target c ID by using the template and provided parameters
*/
func (a *Client) Provision(params *ProvisionParams, opts ...ClientOption) (*ProvisionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProvisionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Provision",
		Method:             "POST",
		PathPattern:        "/workflows/system-definitions/provision/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProvisionReader{formats: a.formats},
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
	success, ok := result.(*ProvisionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Provision: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
WorkflowDefinitionsCombined searches workflow definitions based on the provided filter
*/
func (a *Client) WorkflowDefinitionsCombined(params *WorkflowDefinitionsCombinedParams, opts ...ClientOption) (*WorkflowDefinitionsCombinedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowDefinitionsCombinedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkflowDefinitionsCombined",
		Method:             "GET",
		PathPattern:        "/workflows/combined/definitions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowDefinitionsCombinedReader{formats: a.formats},
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
	success, ok := result.(*WorkflowDefinitionsCombinedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for WorkflowDefinitionsCombined: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
WorkflowDefinitionsCreate creates a workflow definition based on the provided model
*/
func (a *Client) WorkflowDefinitionsCreate(params *WorkflowDefinitionsCreateParams, opts ...ClientOption) (*WorkflowDefinitionsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowDefinitionsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkflowDefinitionsCreate",
		Method:             "POST",
		PathPattern:        "/workflows/entities/definitions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowDefinitionsCreateReader{formats: a.formats},
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
	success, ok := result.(*WorkflowDefinitionsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for WorkflowDefinitionsCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
WorkflowDefinitionsExport exports a workflow definition for the given definition ID
*/
func (a *Client) WorkflowDefinitionsExport(params *WorkflowDefinitionsExportParams, opts ...ClientOption) (*WorkflowDefinitionsExportOK, *WorkflowDefinitionsExportStatus299, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowDefinitionsExportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkflowDefinitionsExport",
		Method:             "GET",
		PathPattern:        "/workflows/entities/definitions/export/v1",
		ProducesMediaTypes: []string{"application/yaml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowDefinitionsExportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *WorkflowDefinitionsExportOK:
		return value, nil, nil
	case *WorkflowDefinitionsExportStatus299:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for workflows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
WorkflowDefinitionsImport imports a workflow definition based on the provided model
*/
func (a *Client) WorkflowDefinitionsImport(params *WorkflowDefinitionsImportParams, opts ...ClientOption) (*WorkflowDefinitionsImportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowDefinitionsImportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkflowDefinitionsImport",
		Method:             "POST",
		PathPattern:        "/workflows/entities/definitions/import/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowDefinitionsImportReader{formats: a.formats},
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
	success, ok := result.(*WorkflowDefinitionsImportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for WorkflowDefinitionsImport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
WorkflowDefinitionsUpdate updates a workflow definition based on the provided model
*/
func (a *Client) WorkflowDefinitionsUpdate(params *WorkflowDefinitionsUpdateParams, opts ...ClientOption) (*WorkflowDefinitionsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowDefinitionsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkflowDefinitionsUpdate",
		Method:             "PUT",
		PathPattern:        "/workflows/entities/definitions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowDefinitionsUpdateReader{formats: a.formats},
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
	success, ok := result.(*WorkflowDefinitionsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for WorkflowDefinitionsUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
WorkflowExecutionsCombined searches workflow executions based on the provided filter
*/
func (a *Client) WorkflowExecutionsCombined(params *WorkflowExecutionsCombinedParams, opts ...ClientOption) (*WorkflowExecutionsCombinedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowExecutionsCombinedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkflowExecutionsCombined",
		Method:             "GET",
		PathPattern:        "/workflows/combined/executions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowExecutionsCombinedReader{formats: a.formats},
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
	success, ok := result.(*WorkflowExecutionsCombinedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for WorkflowExecutionsCombined: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
WorkflowGetHumanInputV1 gets one or more specific human inputs by their i ds
*/
func (a *Client) WorkflowGetHumanInputV1(params *WorkflowGetHumanInputV1Params, opts ...ClientOption) (*WorkflowGetHumanInputV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowGetHumanInputV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkflowGetHumanInputV1",
		Method:             "GET",
		PathPattern:        "/workflows/entities/human-inputs/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowGetHumanInputV1Reader{formats: a.formats},
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
	success, ok := result.(*WorkflowGetHumanInputV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for WorkflowGetHumanInputV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
WorkflowUpdateHumanInputV1 provides an input in response to a human input action depending on action configuration one or more of approve decline and or escalate are permitted
*/
func (a *Client) WorkflowUpdateHumanInputV1(params *WorkflowUpdateHumanInputV1Params, opts ...ClientOption) (*WorkflowUpdateHumanInputV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowUpdateHumanInputV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkflowUpdateHumanInputV1",
		Method:             "PATCH",
		PathPattern:        "/workflows/entities/human-inputs/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowUpdateHumanInputV1Reader{formats: a.formats},
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
	success, ok := result.(*WorkflowUpdateHumanInputV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for WorkflowUpdateHumanInputV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
