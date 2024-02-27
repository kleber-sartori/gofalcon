// Code generated by go-swagger; DO NOT EDIT.

package container_images

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

// ReadCombinedImagesExportReader is a Reader for the ReadCombinedImagesExport structure.
type ReadCombinedImagesExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadCombinedImagesExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadCombinedImagesExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadCombinedImagesExportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewReadCombinedImagesExportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadCombinedImagesExportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/combined/images/export/v1] ReadCombinedImagesExport", response, response.Code())
	}
}

// NewReadCombinedImagesExportOK creates a ReadCombinedImagesExportOK with default headers values
func NewReadCombinedImagesExportOK() *ReadCombinedImagesExportOK {
	return &ReadCombinedImagesExportOK{}
}

/*
ReadCombinedImagesExportOK describes a response with status code 200, with default header values.

OK
*/
type ReadCombinedImagesExportOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ImagesAPICombinedImageExport
}

// IsSuccess returns true when this read combined images export o k response has a 2xx status code
func (o *ReadCombinedImagesExportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read combined images export o k response has a 3xx status code
func (o *ReadCombinedImagesExportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read combined images export o k response has a 4xx status code
func (o *ReadCombinedImagesExportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read combined images export o k response has a 5xx status code
func (o *ReadCombinedImagesExportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read combined images export o k response a status code equal to that given
func (o *ReadCombinedImagesExportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read combined images export o k response
func (o *ReadCombinedImagesExportOK) Code() int {
	return 200
}

func (o *ReadCombinedImagesExportOK) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/images/export/v1][%d] readCombinedImagesExportOK  %+v", 200, o.Payload)
}

func (o *ReadCombinedImagesExportOK) String() string {
	return fmt.Sprintf("[GET /container-security/combined/images/export/v1][%d] readCombinedImagesExportOK  %+v", 200, o.Payload)
}

func (o *ReadCombinedImagesExportOK) GetPayload() *models.ImagesAPICombinedImageExport {
	return o.Payload
}

func (o *ReadCombinedImagesExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ImagesAPICombinedImageExport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadCombinedImagesExportForbidden creates a ReadCombinedImagesExportForbidden with default headers values
func NewReadCombinedImagesExportForbidden() *ReadCombinedImagesExportForbidden {
	return &ReadCombinedImagesExportForbidden{}
}

/*
ReadCombinedImagesExportForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReadCombinedImagesExportForbidden struct {

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

// IsSuccess returns true when this read combined images export forbidden response has a 2xx status code
func (o *ReadCombinedImagesExportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read combined images export forbidden response has a 3xx status code
func (o *ReadCombinedImagesExportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read combined images export forbidden response has a 4xx status code
func (o *ReadCombinedImagesExportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this read combined images export forbidden response has a 5xx status code
func (o *ReadCombinedImagesExportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this read combined images export forbidden response a status code equal to that given
func (o *ReadCombinedImagesExportForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the read combined images export forbidden response
func (o *ReadCombinedImagesExportForbidden) Code() int {
	return 403
}

func (o *ReadCombinedImagesExportForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/images/export/v1][%d] readCombinedImagesExportForbidden  %+v", 403, o.Payload)
}

func (o *ReadCombinedImagesExportForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/combined/images/export/v1][%d] readCombinedImagesExportForbidden  %+v", 403, o.Payload)
}

func (o *ReadCombinedImagesExportForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadCombinedImagesExportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadCombinedImagesExportTooManyRequests creates a ReadCombinedImagesExportTooManyRequests with default headers values
func NewReadCombinedImagesExportTooManyRequests() *ReadCombinedImagesExportTooManyRequests {
	return &ReadCombinedImagesExportTooManyRequests{}
}

/*
ReadCombinedImagesExportTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ReadCombinedImagesExportTooManyRequests struct {

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

// IsSuccess returns true when this read combined images export too many requests response has a 2xx status code
func (o *ReadCombinedImagesExportTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read combined images export too many requests response has a 3xx status code
func (o *ReadCombinedImagesExportTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read combined images export too many requests response has a 4xx status code
func (o *ReadCombinedImagesExportTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this read combined images export too many requests response has a 5xx status code
func (o *ReadCombinedImagesExportTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this read combined images export too many requests response a status code equal to that given
func (o *ReadCombinedImagesExportTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the read combined images export too many requests response
func (o *ReadCombinedImagesExportTooManyRequests) Code() int {
	return 429
}

func (o *ReadCombinedImagesExportTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/images/export/v1][%d] readCombinedImagesExportTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadCombinedImagesExportTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/combined/images/export/v1][%d] readCombinedImagesExportTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadCombinedImagesExportTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadCombinedImagesExportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadCombinedImagesExportInternalServerError creates a ReadCombinedImagesExportInternalServerError with default headers values
func NewReadCombinedImagesExportInternalServerError() *ReadCombinedImagesExportInternalServerError {
	return &ReadCombinedImagesExportInternalServerError{}
}

/*
ReadCombinedImagesExportInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ReadCombinedImagesExportInternalServerError struct {

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

// IsSuccess returns true when this read combined images export internal server error response has a 2xx status code
func (o *ReadCombinedImagesExportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read combined images export internal server error response has a 3xx status code
func (o *ReadCombinedImagesExportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read combined images export internal server error response has a 4xx status code
func (o *ReadCombinedImagesExportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read combined images export internal server error response has a 5xx status code
func (o *ReadCombinedImagesExportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read combined images export internal server error response a status code equal to that given
func (o *ReadCombinedImagesExportInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read combined images export internal server error response
func (o *ReadCombinedImagesExportInternalServerError) Code() int {
	return 500
}

func (o *ReadCombinedImagesExportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/images/export/v1][%d] readCombinedImagesExportInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadCombinedImagesExportInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/combined/images/export/v1][%d] readCombinedImagesExportInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadCombinedImagesExportInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *ReadCombinedImagesExportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
