// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// InitializeSelfServiceSettingsViaBrowserFlowReader is a Reader for the InitializeSelfServiceSettingsViaBrowserFlow structure.
type InitializeSelfServiceSettingsViaBrowserFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitializeSelfServiceSettingsViaBrowserFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewInitializeSelfServiceSettingsViaBrowserFlowFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInitializeSelfServiceSettingsViaBrowserFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInitializeSelfServiceSettingsViaBrowserFlowFound creates a InitializeSelfServiceSettingsViaBrowserFlowFound with default headers values
func NewInitializeSelfServiceSettingsViaBrowserFlowFound() *InitializeSelfServiceSettingsViaBrowserFlowFound {
	return &InitializeSelfServiceSettingsViaBrowserFlowFound{}
}

/*InitializeSelfServiceSettingsViaBrowserFlowFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type InitializeSelfServiceSettingsViaBrowserFlowFound struct {
}

func (o *InitializeSelfServiceSettingsViaBrowserFlowFound) Error() string {
	return fmt.Sprintf("[GET /self-service/settings/browser/flows][%d] initializeSelfServiceSettingsViaBrowserFlowFound ", 302)
}

func (o *InitializeSelfServiceSettingsViaBrowserFlowFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitializeSelfServiceSettingsViaBrowserFlowInternalServerError creates a InitializeSelfServiceSettingsViaBrowserFlowInternalServerError with default headers values
func NewInitializeSelfServiceSettingsViaBrowserFlowInternalServerError() *InitializeSelfServiceSettingsViaBrowserFlowInternalServerError {
	return &InitializeSelfServiceSettingsViaBrowserFlowInternalServerError{}
}

/*InitializeSelfServiceSettingsViaBrowserFlowInternalServerError handles this case with default header values.

genericError
*/
type InitializeSelfServiceSettingsViaBrowserFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *InitializeSelfServiceSettingsViaBrowserFlowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/settings/browser/flows][%d] initializeSelfServiceSettingsViaBrowserFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *InitializeSelfServiceSettingsViaBrowserFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *InitializeSelfServiceSettingsViaBrowserFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
