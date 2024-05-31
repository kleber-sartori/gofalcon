// Code generated by go-swagger; DO NOT EDIT.

package filevantage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// StartActionsReader is a Reader for the StartActions structure.
type StartActionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartActionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStartActionsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStartActionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStartActionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewStartActionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewStartActionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartActionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /filevantage/entities/actions/v1] startActions", response, response.Code())
	}
}

// NewStartActionsAccepted creates a StartActionsAccepted with default headers values
func NewStartActionsAccepted() *StartActionsAccepted {
	return &StartActionsAccepted{}
}

/*
StartActionsAccepted describes a response with status code 202, with default header values.

Action has been accepted for processing; the returned `action_id` can be used to retrieve the details.
*/
type StartActionsAccepted struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ActionsActionResponse
}

// IsSuccess returns true when this start actions accepted response has a 2xx status code
func (o *StartActionsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start actions accepted response has a 3xx status code
func (o *StartActionsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start actions accepted response has a 4xx status code
func (o *StartActionsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this start actions accepted response has a 5xx status code
func (o *StartActionsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this start actions accepted response a status code equal to that given
func (o *StartActionsAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the start actions accepted response
func (o *StartActionsAccepted) Code() int {
	return 202
}

func (o *StartActionsAccepted) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/actions/v1][%d] startActionsAccepted  %+v", 202, o.Payload)
}

func (o *StartActionsAccepted) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/actions/v1][%d] startActionsAccepted  %+v", 202, o.Payload)
}

func (o *StartActionsAccepted) GetPayload() *models.ActionsActionResponse {
	return o.Payload
}

func (o *StartActionsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ActionsActionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartActionsBadRequest creates a StartActionsBadRequest with default headers values
func NewStartActionsBadRequest() *StartActionsBadRequest {
	return &StartActionsBadRequest{}
}

/*
StartActionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StartActionsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this start actions bad request response has a 2xx status code
func (o *StartActionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start actions bad request response has a 3xx status code
func (o *StartActionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start actions bad request response has a 4xx status code
func (o *StartActionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this start actions bad request response has a 5xx status code
func (o *StartActionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this start actions bad request response a status code equal to that given
func (o *StartActionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the start actions bad request response
func (o *StartActionsBadRequest) Code() int {
	return 400
}

func (o *StartActionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/actions/v1][%d] startActionsBadRequest  %+v", 400, o.Payload)
}

func (o *StartActionsBadRequest) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/actions/v1][%d] startActionsBadRequest  %+v", 400, o.Payload)
}

func (o *StartActionsBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *StartActionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartActionsForbidden creates a StartActionsForbidden with default headers values
func NewStartActionsForbidden() *StartActionsForbidden {
	return &StartActionsForbidden{}
}

/*
StartActionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StartActionsForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this start actions forbidden response has a 2xx status code
func (o *StartActionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start actions forbidden response has a 3xx status code
func (o *StartActionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start actions forbidden response has a 4xx status code
func (o *StartActionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this start actions forbidden response has a 5xx status code
func (o *StartActionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this start actions forbidden response a status code equal to that given
func (o *StartActionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the start actions forbidden response
func (o *StartActionsForbidden) Code() int {
	return 403
}

func (o *StartActionsForbidden) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/actions/v1][%d] startActionsForbidden  %+v", 403, o.Payload)
}

func (o *StartActionsForbidden) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/actions/v1][%d] startActionsForbidden  %+v", 403, o.Payload)
}

func (o *StartActionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *StartActionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartActionsConflict creates a StartActionsConflict with default headers values
func NewStartActionsConflict() *StartActionsConflict {
	return &StartActionsConflict{}
}

/*
StartActionsConflict describes a response with status code 409, with default header values.

Conflict
*/
type StartActionsConflict struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this start actions conflict response has a 2xx status code
func (o *StartActionsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start actions conflict response has a 3xx status code
func (o *StartActionsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start actions conflict response has a 4xx status code
func (o *StartActionsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this start actions conflict response has a 5xx status code
func (o *StartActionsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this start actions conflict response a status code equal to that given
func (o *StartActionsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the start actions conflict response
func (o *StartActionsConflict) Code() int {
	return 409
}

func (o *StartActionsConflict) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/actions/v1][%d] startActionsConflict  %+v", 409, o.Payload)
}

func (o *StartActionsConflict) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/actions/v1][%d] startActionsConflict  %+v", 409, o.Payload)
}

func (o *StartActionsConflict) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *StartActionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartActionsTooManyRequests creates a StartActionsTooManyRequests with default headers values
func NewStartActionsTooManyRequests() *StartActionsTooManyRequests {
	return &StartActionsTooManyRequests{}
}

/*
StartActionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type StartActionsTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this start actions too many requests response has a 2xx status code
func (o *StartActionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start actions too many requests response has a 3xx status code
func (o *StartActionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start actions too many requests response has a 4xx status code
func (o *StartActionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this start actions too many requests response has a 5xx status code
func (o *StartActionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this start actions too many requests response a status code equal to that given
func (o *StartActionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the start actions too many requests response
func (o *StartActionsTooManyRequests) Code() int {
	return 429
}

func (o *StartActionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/actions/v1][%d] startActionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *StartActionsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/actions/v1][%d] startActionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *StartActionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *StartActionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartActionsInternalServerError creates a StartActionsInternalServerError with default headers values
func NewStartActionsInternalServerError() *StartActionsInternalServerError {
	return &StartActionsInternalServerError{}
}

/*
StartActionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type StartActionsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this start actions internal server error response has a 2xx status code
func (o *StartActionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start actions internal server error response has a 3xx status code
func (o *StartActionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start actions internal server error response has a 4xx status code
func (o *StartActionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this start actions internal server error response has a 5xx status code
func (o *StartActionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this start actions internal server error response a status code equal to that given
func (o *StartActionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the start actions internal server error response
func (o *StartActionsInternalServerError) Code() int {
	return 500
}

func (o *StartActionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/actions/v1][%d] startActionsInternalServerError  %+v", 500, o.Payload)
}

func (o *StartActionsInternalServerError) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/actions/v1][%d] startActionsInternalServerError  %+v", 500, o.Payload)
}

func (o *StartActionsInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *StartActionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
