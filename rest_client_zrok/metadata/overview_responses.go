// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti-test-kitchen/zrok/rest_model_zrok"
)

// OverviewReader is a Reader for the Overview structure.
type OverviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OverviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOverviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewOverviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOverviewOK creates a OverviewOK with default headers values
func NewOverviewOK() *OverviewOK {
	return &OverviewOK{}
}

/* OverviewOK describes a response with status code 200, with default header values.

overview returned
*/
type OverviewOK struct {
	Payload *rest_model_zrok.EnvironmentServices
}

func (o *OverviewOK) Error() string {
	return fmt.Sprintf("[GET /overview][%d] overviewOK  %+v", 200, o.Payload)
}
func (o *OverviewOK) GetPayload() *rest_model_zrok.EnvironmentServices {
	return o.Payload
}

func (o *OverviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model_zrok.EnvironmentServices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOverviewInternalServerError creates a OverviewInternalServerError with default headers values
func NewOverviewInternalServerError() *OverviewInternalServerError {
	return &OverviewInternalServerError{}
}

/* OverviewInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type OverviewInternalServerError struct {
	Payload rest_model_zrok.ErrorMessage
}

func (o *OverviewInternalServerError) Error() string {
	return fmt.Sprintf("[GET /overview][%d] overviewInternalServerError  %+v", 500, o.Payload)
}
func (o *OverviewInternalServerError) GetPayload() rest_model_zrok.ErrorMessage {
	return o.Payload
}

func (o *OverviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
