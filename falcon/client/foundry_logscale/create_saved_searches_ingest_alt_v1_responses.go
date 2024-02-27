// Code generated by go-swagger; DO NOT EDIT.

package foundry_logscale

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

// CreateSavedSearchesIngestAltV1Reader is a Reader for the CreateSavedSearchesIngestAltV1 structure.
type CreateSavedSearchesIngestAltV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSavedSearchesIngestAltV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSavedSearchesIngestAltV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateSavedSearchesIngestAltV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateSavedSearchesIngestAltV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /loggingapi/entities/saved-searches-ingest/v1] CreateSavedSearchesIngestAltV1", response, response.Code())
	}
}

// NewCreateSavedSearchesIngestAltV1OK creates a CreateSavedSearchesIngestAltV1OK with default headers values
func NewCreateSavedSearchesIngestAltV1OK() *CreateSavedSearchesIngestAltV1OK {
	return &CreateSavedSearchesIngestAltV1OK{}
}

/*
CreateSavedSearchesIngestAltV1OK describes a response with status code 200, with default header values.

OK
*/
type CreateSavedSearchesIngestAltV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientDataIngestResponseWrapperV1
}

// IsSuccess returns true when this create saved searches ingest alt v1 o k response has a 2xx status code
func (o *CreateSavedSearchesIngestAltV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create saved searches ingest alt v1 o k response has a 3xx status code
func (o *CreateSavedSearchesIngestAltV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saved searches ingest alt v1 o k response has a 4xx status code
func (o *CreateSavedSearchesIngestAltV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create saved searches ingest alt v1 o k response has a 5xx status code
func (o *CreateSavedSearchesIngestAltV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this create saved searches ingest alt v1 o k response a status code equal to that given
func (o *CreateSavedSearchesIngestAltV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create saved searches ingest alt v1 o k response
func (o *CreateSavedSearchesIngestAltV1OK) Code() int {
	return 200
}

func (o *CreateSavedSearchesIngestAltV1OK) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-ingest/v1][%d] createSavedSearchesIngestAltV1OK  %+v", 200, o.Payload)
}

func (o *CreateSavedSearchesIngestAltV1OK) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-ingest/v1][%d] createSavedSearchesIngestAltV1OK  %+v", 200, o.Payload)
}

func (o *CreateSavedSearchesIngestAltV1OK) GetPayload() *models.ClientDataIngestResponseWrapperV1 {
	return o.Payload
}

func (o *CreateSavedSearchesIngestAltV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientDataIngestResponseWrapperV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSavedSearchesIngestAltV1Forbidden creates a CreateSavedSearchesIngestAltV1Forbidden with default headers values
func NewCreateSavedSearchesIngestAltV1Forbidden() *CreateSavedSearchesIngestAltV1Forbidden {
	return &CreateSavedSearchesIngestAltV1Forbidden{}
}

/*
CreateSavedSearchesIngestAltV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateSavedSearchesIngestAltV1Forbidden struct {

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

// IsSuccess returns true when this create saved searches ingest alt v1 forbidden response has a 2xx status code
func (o *CreateSavedSearchesIngestAltV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create saved searches ingest alt v1 forbidden response has a 3xx status code
func (o *CreateSavedSearchesIngestAltV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saved searches ingest alt v1 forbidden response has a 4xx status code
func (o *CreateSavedSearchesIngestAltV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create saved searches ingest alt v1 forbidden response has a 5xx status code
func (o *CreateSavedSearchesIngestAltV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create saved searches ingest alt v1 forbidden response a status code equal to that given
func (o *CreateSavedSearchesIngestAltV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create saved searches ingest alt v1 forbidden response
func (o *CreateSavedSearchesIngestAltV1Forbidden) Code() int {
	return 403
}

func (o *CreateSavedSearchesIngestAltV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-ingest/v1][%d] createSavedSearchesIngestAltV1Forbidden  %+v", 403, o.Payload)
}

func (o *CreateSavedSearchesIngestAltV1Forbidden) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-ingest/v1][%d] createSavedSearchesIngestAltV1Forbidden  %+v", 403, o.Payload)
}

func (o *CreateSavedSearchesIngestAltV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateSavedSearchesIngestAltV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSavedSearchesIngestAltV1TooManyRequests creates a CreateSavedSearchesIngestAltV1TooManyRequests with default headers values
func NewCreateSavedSearchesIngestAltV1TooManyRequests() *CreateSavedSearchesIngestAltV1TooManyRequests {
	return &CreateSavedSearchesIngestAltV1TooManyRequests{}
}

/*
CreateSavedSearchesIngestAltV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateSavedSearchesIngestAltV1TooManyRequests struct {

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

// IsSuccess returns true when this create saved searches ingest alt v1 too many requests response has a 2xx status code
func (o *CreateSavedSearchesIngestAltV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create saved searches ingest alt v1 too many requests response has a 3xx status code
func (o *CreateSavedSearchesIngestAltV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saved searches ingest alt v1 too many requests response has a 4xx status code
func (o *CreateSavedSearchesIngestAltV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create saved searches ingest alt v1 too many requests response has a 5xx status code
func (o *CreateSavedSearchesIngestAltV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create saved searches ingest alt v1 too many requests response a status code equal to that given
func (o *CreateSavedSearchesIngestAltV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create saved searches ingest alt v1 too many requests response
func (o *CreateSavedSearchesIngestAltV1TooManyRequests) Code() int {
	return 429
}

func (o *CreateSavedSearchesIngestAltV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-ingest/v1][%d] createSavedSearchesIngestAltV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateSavedSearchesIngestAltV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-ingest/v1][%d] createSavedSearchesIngestAltV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateSavedSearchesIngestAltV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateSavedSearchesIngestAltV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
