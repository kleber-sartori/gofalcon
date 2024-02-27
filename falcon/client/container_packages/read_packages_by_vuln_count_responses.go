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

// ReadPackagesByVulnCountReader is a Reader for the ReadPackagesByVulnCount structure.
type ReadPackagesByVulnCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadPackagesByVulnCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadPackagesByVulnCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadPackagesByVulnCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewReadPackagesByVulnCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadPackagesByVulnCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/combined/packages/by-vulnerability-count/v1] ReadPackagesByVulnCount", response, response.Code())
	}
}

// NewReadPackagesByVulnCountOK creates a ReadPackagesByVulnCountOK with default headers values
func NewReadPackagesByVulnCountOK() *ReadPackagesByVulnCountOK {
	return &ReadPackagesByVulnCountOK{}
}

/*
ReadPackagesByVulnCountOK describes a response with status code 200, with default header values.

OK
*/
type ReadPackagesByVulnCountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.PackagesAPIPackagesByVulnCount
}

// IsSuccess returns true when this read packages by vuln count o k response has a 2xx status code
func (o *ReadPackagesByVulnCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read packages by vuln count o k response has a 3xx status code
func (o *ReadPackagesByVulnCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read packages by vuln count o k response has a 4xx status code
func (o *ReadPackagesByVulnCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read packages by vuln count o k response has a 5xx status code
func (o *ReadPackagesByVulnCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read packages by vuln count o k response a status code equal to that given
func (o *ReadPackagesByVulnCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read packages by vuln count o k response
func (o *ReadPackagesByVulnCountOK) Code() int {
	return 200
}

func (o *ReadPackagesByVulnCountOK) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/by-vulnerability-count/v1][%d] readPackagesByVulnCountOK  %+v", 200, o.Payload)
}

func (o *ReadPackagesByVulnCountOK) String() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/by-vulnerability-count/v1][%d] readPackagesByVulnCountOK  %+v", 200, o.Payload)
}

func (o *ReadPackagesByVulnCountOK) GetPayload() *models.PackagesAPIPackagesByVulnCount {
	return o.Payload
}

func (o *ReadPackagesByVulnCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.PackagesAPIPackagesByVulnCount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPackagesByVulnCountForbidden creates a ReadPackagesByVulnCountForbidden with default headers values
func NewReadPackagesByVulnCountForbidden() *ReadPackagesByVulnCountForbidden {
	return &ReadPackagesByVulnCountForbidden{}
}

/*
ReadPackagesByVulnCountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReadPackagesByVulnCountForbidden struct {

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

// IsSuccess returns true when this read packages by vuln count forbidden response has a 2xx status code
func (o *ReadPackagesByVulnCountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read packages by vuln count forbidden response has a 3xx status code
func (o *ReadPackagesByVulnCountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read packages by vuln count forbidden response has a 4xx status code
func (o *ReadPackagesByVulnCountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this read packages by vuln count forbidden response has a 5xx status code
func (o *ReadPackagesByVulnCountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this read packages by vuln count forbidden response a status code equal to that given
func (o *ReadPackagesByVulnCountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the read packages by vuln count forbidden response
func (o *ReadPackagesByVulnCountForbidden) Code() int {
	return 403
}

func (o *ReadPackagesByVulnCountForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/by-vulnerability-count/v1][%d] readPackagesByVulnCountForbidden  %+v", 403, o.Payload)
}

func (o *ReadPackagesByVulnCountForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/by-vulnerability-count/v1][%d] readPackagesByVulnCountForbidden  %+v", 403, o.Payload)
}

func (o *ReadPackagesByVulnCountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadPackagesByVulnCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadPackagesByVulnCountTooManyRequests creates a ReadPackagesByVulnCountTooManyRequests with default headers values
func NewReadPackagesByVulnCountTooManyRequests() *ReadPackagesByVulnCountTooManyRequests {
	return &ReadPackagesByVulnCountTooManyRequests{}
}

/*
ReadPackagesByVulnCountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ReadPackagesByVulnCountTooManyRequests struct {

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

// IsSuccess returns true when this read packages by vuln count too many requests response has a 2xx status code
func (o *ReadPackagesByVulnCountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read packages by vuln count too many requests response has a 3xx status code
func (o *ReadPackagesByVulnCountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read packages by vuln count too many requests response has a 4xx status code
func (o *ReadPackagesByVulnCountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this read packages by vuln count too many requests response has a 5xx status code
func (o *ReadPackagesByVulnCountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this read packages by vuln count too many requests response a status code equal to that given
func (o *ReadPackagesByVulnCountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the read packages by vuln count too many requests response
func (o *ReadPackagesByVulnCountTooManyRequests) Code() int {
	return 429
}

func (o *ReadPackagesByVulnCountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/by-vulnerability-count/v1][%d] readPackagesByVulnCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadPackagesByVulnCountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/by-vulnerability-count/v1][%d] readPackagesByVulnCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadPackagesByVulnCountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadPackagesByVulnCountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadPackagesByVulnCountInternalServerError creates a ReadPackagesByVulnCountInternalServerError with default headers values
func NewReadPackagesByVulnCountInternalServerError() *ReadPackagesByVulnCountInternalServerError {
	return &ReadPackagesByVulnCountInternalServerError{}
}

/*
ReadPackagesByVulnCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ReadPackagesByVulnCountInternalServerError struct {

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

// IsSuccess returns true when this read packages by vuln count internal server error response has a 2xx status code
func (o *ReadPackagesByVulnCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read packages by vuln count internal server error response has a 3xx status code
func (o *ReadPackagesByVulnCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read packages by vuln count internal server error response has a 4xx status code
func (o *ReadPackagesByVulnCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read packages by vuln count internal server error response has a 5xx status code
func (o *ReadPackagesByVulnCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read packages by vuln count internal server error response a status code equal to that given
func (o *ReadPackagesByVulnCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read packages by vuln count internal server error response
func (o *ReadPackagesByVulnCountInternalServerError) Code() int {
	return 500
}

func (o *ReadPackagesByVulnCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/by-vulnerability-count/v1][%d] readPackagesByVulnCountInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadPackagesByVulnCountInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/combined/packages/by-vulnerability-count/v1][%d] readPackagesByVulnCountInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadPackagesByVulnCountInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *ReadPackagesByVulnCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
