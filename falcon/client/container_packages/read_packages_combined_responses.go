// Code generated by go-swagger; DO NOT EDIT.

package container_packages

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

// ReadPackagesCombinedReader is a Reader for the ReadPackagesCombined structure.
type ReadPackagesCombinedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadPackagesCombinedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadPackagesCombinedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadPackagesCombinedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewReadPackagesCombinedTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadPackagesCombinedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/combined/packages/v1] ReadPackagesCombined", response, response.Code())
	}
}

// NewReadPackagesCombinedOK creates a ReadPackagesCombinedOK with default headers values
func NewReadPackagesCombinedOK() *ReadPackagesCombinedOK {
	return &ReadPackagesCombinedOK{}
}

/*
ReadPackagesCombinedOK describes a response with status code 200, with default header values.

OK
*/
type ReadPackagesCombinedOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.PackagesAPICombinedPackage
}

// IsSuccess returns true when this read packages combined o k response has a 2xx status code
func (o *ReadPackagesCombinedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read packages combined o k response has a 3xx status code
func (o *ReadPackagesCombinedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read packages combined o k response has a 4xx status code
func (o *ReadPackagesCombinedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read packages combined o k response has a 5xx status code
func (o *ReadPackagesCombinedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read packages combined o k response a status code equal to that given
func (o *ReadPackagesCombinedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read packages combined o k response
func (o *ReadPackagesCombinedOK) Code() int {
	return 200
}

func (o *ReadPackagesCombinedOK) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/v1][%d] readPackagesCombinedOK  %+v", 200, o.Payload)
}

func (o *ReadPackagesCombinedOK) String() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/v1][%d] readPackagesCombinedOK  %+v", 200, o.Payload)
}

func (o *ReadPackagesCombinedOK) GetPayload() *models.PackagesAPICombinedPackage {
	return o.Payload
}

func (o *ReadPackagesCombinedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.PackagesAPICombinedPackage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPackagesCombinedForbidden creates a ReadPackagesCombinedForbidden with default headers values
func NewReadPackagesCombinedForbidden() *ReadPackagesCombinedForbidden {
	return &ReadPackagesCombinedForbidden{}
}

/*
ReadPackagesCombinedForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReadPackagesCombinedForbidden struct {

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

// IsSuccess returns true when this read packages combined forbidden response has a 2xx status code
func (o *ReadPackagesCombinedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read packages combined forbidden response has a 3xx status code
func (o *ReadPackagesCombinedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read packages combined forbidden response has a 4xx status code
func (o *ReadPackagesCombinedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this read packages combined forbidden response has a 5xx status code
func (o *ReadPackagesCombinedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this read packages combined forbidden response a status code equal to that given
func (o *ReadPackagesCombinedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the read packages combined forbidden response
func (o *ReadPackagesCombinedForbidden) Code() int {
	return 403
}

func (o *ReadPackagesCombinedForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/v1][%d] readPackagesCombinedForbidden  %+v", 403, o.Payload)
}

func (o *ReadPackagesCombinedForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/v1][%d] readPackagesCombinedForbidden  %+v", 403, o.Payload)
}

func (o *ReadPackagesCombinedForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadPackagesCombinedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadPackagesCombinedTooManyRequests creates a ReadPackagesCombinedTooManyRequests with default headers values
func NewReadPackagesCombinedTooManyRequests() *ReadPackagesCombinedTooManyRequests {
	return &ReadPackagesCombinedTooManyRequests{}
}

/*
ReadPackagesCombinedTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ReadPackagesCombinedTooManyRequests struct {

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

// IsSuccess returns true when this read packages combined too many requests response has a 2xx status code
func (o *ReadPackagesCombinedTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read packages combined too many requests response has a 3xx status code
func (o *ReadPackagesCombinedTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read packages combined too many requests response has a 4xx status code
func (o *ReadPackagesCombinedTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this read packages combined too many requests response has a 5xx status code
func (o *ReadPackagesCombinedTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this read packages combined too many requests response a status code equal to that given
func (o *ReadPackagesCombinedTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the read packages combined too many requests response
func (o *ReadPackagesCombinedTooManyRequests) Code() int {
	return 429
}

func (o *ReadPackagesCombinedTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/v1][%d] readPackagesCombinedTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadPackagesCombinedTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/v1][%d] readPackagesCombinedTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadPackagesCombinedTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadPackagesCombinedTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadPackagesCombinedInternalServerError creates a ReadPackagesCombinedInternalServerError with default headers values
func NewReadPackagesCombinedInternalServerError() *ReadPackagesCombinedInternalServerError {
	return &ReadPackagesCombinedInternalServerError{}
}

/*
ReadPackagesCombinedInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ReadPackagesCombinedInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CoreEntitiesResponse
}

// IsSuccess returns true when this read packages combined internal server error response has a 2xx status code
func (o *ReadPackagesCombinedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read packages combined internal server error response has a 3xx status code
func (o *ReadPackagesCombinedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read packages combined internal server error response has a 4xx status code
func (o *ReadPackagesCombinedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read packages combined internal server error response has a 5xx status code
func (o *ReadPackagesCombinedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read packages combined internal server error response a status code equal to that given
func (o *ReadPackagesCombinedInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read packages combined internal server error response
func (o *ReadPackagesCombinedInternalServerError) Code() int {
	return 500
}

func (o *ReadPackagesCombinedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/v1][%d] readPackagesCombinedInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadPackagesCombinedInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/v1][%d] readPackagesCombinedInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadPackagesCombinedInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *ReadPackagesCombinedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CoreEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
