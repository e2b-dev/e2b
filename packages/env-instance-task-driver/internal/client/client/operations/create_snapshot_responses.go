// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/e2b-dev/api/packages/env-instance-task-driver/internal/client/models"
)

// CreateSnapshotReader is a Reader for the CreateSnapshot structure.
type CreateSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCreateSnapshotNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSnapshotNoContent creates a CreateSnapshotNoContent with default headers values
func NewCreateSnapshotNoContent() *CreateSnapshotNoContent {
	return &CreateSnapshotNoContent{}
}

/*
CreateSnapshotNoContent describes a response with status code 204, with default header values.

Snapshot created
*/
type CreateSnapshotNoContent struct {
}

// IsSuccess returns true when this create snapshot no content response has a 2xx status code
func (o *CreateSnapshotNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create snapshot no content response has a 3xx status code
func (o *CreateSnapshotNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create snapshot no content response has a 4xx status code
func (o *CreateSnapshotNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this create snapshot no content response has a 5xx status code
func (o *CreateSnapshotNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this create snapshot no content response a status code equal to that given
func (o *CreateSnapshotNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the create snapshot no content response
func (o *CreateSnapshotNoContent) Code() int {
	return 204
}

func (o *CreateSnapshotNoContent) Error() string {
	return fmt.Sprintf("[PUT /snapshot/create][%d] createSnapshotNoContent ", 204)
}

func (o *CreateSnapshotNoContent) String() string {
	return fmt.Sprintf("[PUT /snapshot/create][%d] createSnapshotNoContent ", 204)
}

func (o *CreateSnapshotNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSnapshotBadRequest creates a CreateSnapshotBadRequest with default headers values
func NewCreateSnapshotBadRequest() *CreateSnapshotBadRequest {
	return &CreateSnapshotBadRequest{}
}

/*
CreateSnapshotBadRequest describes a response with status code 400, with default header values.

Snapshot cannot be created due to bad input
*/
type CreateSnapshotBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create snapshot bad request response has a 2xx status code
func (o *CreateSnapshotBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create snapshot bad request response has a 3xx status code
func (o *CreateSnapshotBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create snapshot bad request response has a 4xx status code
func (o *CreateSnapshotBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create snapshot bad request response has a 5xx status code
func (o *CreateSnapshotBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create snapshot bad request response a status code equal to that given
func (o *CreateSnapshotBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create snapshot bad request response
func (o *CreateSnapshotBadRequest) Code() int {
	return 400
}

func (o *CreateSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[PUT /snapshot/create][%d] createSnapshotBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSnapshotBadRequest) String() string {
	return fmt.Sprintf("[PUT /snapshot/create][%d] createSnapshotBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSnapshotBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnapshotDefault creates a CreateSnapshotDefault with default headers values
func NewCreateSnapshotDefault(code int) *CreateSnapshotDefault {
	return &CreateSnapshotDefault{
		_statusCode: code,
	}
}

/*
CreateSnapshotDefault describes a response with status code -1, with default header values.

Internal server error
*/
type CreateSnapshotDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create snapshot default response has a 2xx status code
func (o *CreateSnapshotDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create snapshot default response has a 3xx status code
func (o *CreateSnapshotDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create snapshot default response has a 4xx status code
func (o *CreateSnapshotDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create snapshot default response has a 5xx status code
func (o *CreateSnapshotDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create snapshot default response a status code equal to that given
func (o *CreateSnapshotDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create snapshot default response
func (o *CreateSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *CreateSnapshotDefault) Error() string {
	return fmt.Sprintf("[PUT /snapshot/create][%d] createSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSnapshotDefault) String() string {
	return fmt.Sprintf("[PUT /snapshot/create][%d] createSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSnapshotDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
