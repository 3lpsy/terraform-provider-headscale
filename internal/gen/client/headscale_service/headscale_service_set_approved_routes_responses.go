// Code generated by go-swagger; DO NOT EDIT.

package headscale_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/models"
)

// HeadscaleServiceSetApprovedRoutesReader is a Reader for the HeadscaleServiceSetApprovedRoutes structure.
type HeadscaleServiceSetApprovedRoutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceSetApprovedRoutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceSetApprovedRoutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceSetApprovedRoutesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceSetApprovedRoutesOK creates a HeadscaleServiceSetApprovedRoutesOK with default headers values
func NewHeadscaleServiceSetApprovedRoutesOK() *HeadscaleServiceSetApprovedRoutesOK {
	return &HeadscaleServiceSetApprovedRoutesOK{}
}

/*
HeadscaleServiceSetApprovedRoutesOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceSetApprovedRoutesOK struct {
	Payload *models.V1SetApprovedRoutesResponse
}

// IsSuccess returns true when this headscale service set approved routes o k response has a 2xx status code
func (o *HeadscaleServiceSetApprovedRoutesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service set approved routes o k response has a 3xx status code
func (o *HeadscaleServiceSetApprovedRoutesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service set approved routes o k response has a 4xx status code
func (o *HeadscaleServiceSetApprovedRoutesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service set approved routes o k response has a 5xx status code
func (o *HeadscaleServiceSetApprovedRoutesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service set approved routes o k response a status code equal to that given
func (o *HeadscaleServiceSetApprovedRoutesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the headscale service set approved routes o k response
func (o *HeadscaleServiceSetApprovedRoutesOK) Code() int {
	return 200
}

func (o *HeadscaleServiceSetApprovedRoutesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/node/{nodeId}/approve_routes][%d] headscaleServiceSetApprovedRoutesOK %s", 200, payload)
}

func (o *HeadscaleServiceSetApprovedRoutesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/node/{nodeId}/approve_routes][%d] headscaleServiceSetApprovedRoutesOK %s", 200, payload)
}

func (o *HeadscaleServiceSetApprovedRoutesOK) GetPayload() *models.V1SetApprovedRoutesResponse {
	return o.Payload
}

func (o *HeadscaleServiceSetApprovedRoutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SetApprovedRoutesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceSetApprovedRoutesDefault creates a HeadscaleServiceSetApprovedRoutesDefault with default headers values
func NewHeadscaleServiceSetApprovedRoutesDefault(code int) *HeadscaleServiceSetApprovedRoutesDefault {
	return &HeadscaleServiceSetApprovedRoutesDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceSetApprovedRoutesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceSetApprovedRoutesDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this headscale service set approved routes default response has a 2xx status code
func (o *HeadscaleServiceSetApprovedRoutesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service set approved routes default response has a 3xx status code
func (o *HeadscaleServiceSetApprovedRoutesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service set approved routes default response has a 4xx status code
func (o *HeadscaleServiceSetApprovedRoutesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service set approved routes default response has a 5xx status code
func (o *HeadscaleServiceSetApprovedRoutesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service set approved routes default response a status code equal to that given
func (o *HeadscaleServiceSetApprovedRoutesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the headscale service set approved routes default response
func (o *HeadscaleServiceSetApprovedRoutesDefault) Code() int {
	return o._statusCode
}

func (o *HeadscaleServiceSetApprovedRoutesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/node/{nodeId}/approve_routes][%d] HeadscaleService_SetApprovedRoutes default %s", o._statusCode, payload)
}

func (o *HeadscaleServiceSetApprovedRoutesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/node/{nodeId}/approve_routes][%d] HeadscaleService_SetApprovedRoutes default %s", o._statusCode, payload)
}

func (o *HeadscaleServiceSetApprovedRoutesDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceSetApprovedRoutesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
