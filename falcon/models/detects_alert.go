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

// DetectsAlert detects alert
//
// swagger:model detects.Alert
type DetectsAlert struct {

	// Device or sensor ID for which the Alert was generated
	AgentID string `json:"agent_id,omitempty"`

	// Common linkage between multiple Alerts that belong to the same detection bouquet
	// Required: true
	AggregateID *string `json:"aggregate_id"`

	// Name of the person this Alert is assigned to
	AssignedToName string `json:"assigned_to_name,omitempty"`

	// UserID to which this Alert is assigned to
	AssignedToUID string `json:"assigned_to_uid,omitempty"`

	// UUID to which this Alert is assigned to
	AssignedToUUID string `json:"assigned_to_uuid,omitempty"`

	// Unique ID of CrowdStrike customers
	Cid string `json:"cid,omitempty"`

	// An opaque internal identifier that can uniquely identify an Alert
	CompositeID string `json:"composite_id,omitempty"`

	// Confidence is a 1-100 integer value denoting the confidence that, when this Alert fires, it is indicative of malicious activity
	Confidence int64 `json:"confidence,omitempty"`

	// internal only
	CrawlEdgeIds map[string][]string `json:"crawl_edge_ids,omitempty"`

	// internal only
	CrawlVertexIds map[string][]string `json:"crawl_vertex_ids,omitempty"`

	// indicates when ThreatGraph was crawled to gather info for this alert creation/update
	// Format: date-time
	CrawledTimestamp strfmt.DateTime `json:"crawled_timestamp,omitempty"`

	// indicates when the Alert was first written to backend store
	// Format: date-time
	CreatedTimestamp strfmt.DateTime `json:"created_timestamp,omitempty"`

	// Data Domains represents domains to which this alert belongs to
	DataDomains []string `json:"data_domains"`

	// Short, customer-visible summary of the detected activity
	Description string `json:"description,omitempty"`

	// Customer visible name for the Alert's pattern
	DisplayName string `json:"display_name,omitempty"`

	// Boolean to know if we sent email regarding this Alert
	EmailSent bool `json:"email_sent,omitempty"`

	// internal only
	EsDocID string `json:"es_doc_id,omitempty"`

	// internal only
	EsDocVersion int64 `json:"es_doc_version,omitempty"`

	// internal only
	EsRoutingID string `json:"es_routing_id,omitempty"`

	// internal only
	EsSourceSize int32 `json:"es_source_size,omitempty"`

	// Boolean indicating if this Alert is internal or external
	External bool `json:"external,omitempty"`

	// Vertex key which triggers the formation of the Alert
	ID string `json:"id,omitempty"`

	// Pattern Name coming either from Taxonomy or directly from the ingested Alert
	Name string `json:"name,omitempty"`

	// End goal that an attack adversary intends to achieve according to MITRE
	Objective string `json:"objective,omitempty"`

	// Taxonomy patternID for this Alert
	PatternID int64 `json:"pattern_id,omitempty"`

	// Platform that this Alert was triggered on e.g. Android, Windows, etc..
	Platform string `json:"platform,omitempty"`

	// poly id
	PolyID string `json:"poly_id,omitempty"`

	// Product specifies the SKU that this Alert belongs to e.g. mobile, idp, epp
	Product string `json:"product,omitempty"`

	// indicates when the Alert was marked as 'Resolved'
	// Format: date-time
	ResolvedTimestamp strfmt.DateTime `json:"resolved_timestamp,omitempty"`

	// Scenario was used pre-Handrails to display additional killchain context for UI alerts. With handrails, this field is mostly  obsolete in favor of tactic/technique. Still, it can be useful for determining specific pattern types that are not straightforward to distinguish from other fields alone
	Scenario string `json:"scenario,omitempty"`

	// Seconds To Resolved represents the seconds elapsed since this alert has been resolved
	// Required: true
	SecondsToResolved *int64 `json:"seconds_to_resolved"`

	// Seconds To Triage represents the seconds elapsed since this alert has been triaged
	// Required: true
	SecondsToTriaged *int64 `json:"seconds_to_triaged"`

	// Severity is also a 1-100 integer value, but unlike confidence severity impacts how a Alert is displayed in the UI
	Severity int64 `json:"severity,omitempty"`

	// Severity name is a UI friendly bucketing of the severity integer
	SeverityName string `json:"severity_name,omitempty"`

	// Boolean indicating if this Alert will be shown in the UI or if it's hidden'
	ShowInUI bool `json:"show_in_ui,omitempty"`

	// show in ui updated timestamp
	// Format: date-time
	ShowInUIUpdatedTimestamp strfmt.DateTime `json:"show_in_ui_updated_timestamp,omitempty"`

	// Source Products are products that produced events which contributed to this alert
	SourceProducts []string `json:"source_products"`

	// Source Vendors are vendors that produced events which contributed to this alert
	SourceVendors []string `json:"source_vendors"`

	// Could be one of the following - New, closed, in_progress, reopened
	Status string `json:"status,omitempty"`

	// Tactic and Technique are references to MITRE ATT&CK, which is a public framework for tracking and modeling adversary tools techniques and procedures
	Tactic string `json:"tactic,omitempty"`

	// Unique ID for the tactic seen in the Alert
	TacticID string `json:"tactic_id,omitempty"`

	// Tags are string values associated with the alert that can be added or removed through the API
	Tags []string `json:"tags"`

	// Tactic and Technique are references to MITRE ATT&CK, which is a public framework for tracking and modeling adversary tools techniques and procedures
	Technique string `json:"technique,omitempty"`

	// Unique ID for the technique seen in the Alert
	TechniqueID string `json:"technique_id,omitempty"`

	// stored value coming in directly from the ingested event or set by cloud in the absence of it
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// Type of definition Detections Extensibility use. Keyed-off of Pattern of the incoming events/Alerts
	// Required: true
	Type *string `json:"type"`

	// indicates when the Alert was last modified
	// Format: date-time
	UpdatedTimestamp strfmt.DateTime `json:"updated_timestamp,omitempty"`
}

// Validate validates this detects alert
func (m *DetectsAlert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrawledTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolvedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondsToResolved(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondsToTriaged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShowInUIUpdatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetectsAlert) validateAggregateID(formats strfmt.Registry) error {

	if err := validate.Required("aggregate_id", "body", m.AggregateID); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateCrawledTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CrawledTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("crawled_timestamp", "body", "date-time", m.CrawledTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateCreatedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("created_timestamp", "body", "date-time", m.CreatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateResolvedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.ResolvedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("resolved_timestamp", "body", "date-time", m.ResolvedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateSecondsToResolved(formats strfmt.Registry) error {

	if err := validate.Required("seconds_to_resolved", "body", m.SecondsToResolved); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateSecondsToTriaged(formats strfmt.Registry) error {

	if err := validate.Required("seconds_to_triaged", "body", m.SecondsToTriaged); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateShowInUIUpdatedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.ShowInUIUpdatedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("show_in_ui_updated_timestamp", "body", "date-time", m.ShowInUIUpdatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *DetectsAlert) validateUpdatedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_timestamp", "body", "date-time", m.UpdatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this detects alert based on context it is used
func (m *DetectsAlert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DetectsAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetectsAlert) UnmarshalBinary(b []byte) error {
	var res DetectsAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
