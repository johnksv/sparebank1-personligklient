// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/johnksv/sparebank1-pm/models"
)

// GetDefaultBalanceAccountReader is a Reader for the GetDefaultBalanceAccount structure.
type GetDefaultBalanceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDefaultBalanceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDefaultBalanceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetDefaultBalanceAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetDefaultBalanceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDefaultBalanceAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDefaultBalanceAccountOK creates a GetDefaultBalanceAccountOK with default headers values
func NewGetDefaultBalanceAccountOK() *GetDefaultBalanceAccountOK {
	return &GetDefaultBalanceAccountOK{}
}

/*GetDefaultBalanceAccountOK handles this case with default header values.

Successful response
*/
type GetDefaultBalanceAccountOK struct {
	Payload *models.AccountDto
}

func (o *GetDefaultBalanceAccountOK) Error() string {
	return fmt.Sprintf("[GET /default][%d] getDefaultBalanceAccountOK  %+v", 200, o.Payload)
}

func (o *GetDefaultBalanceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDefaultBalanceAccountUnauthorized creates a GetDefaultBalanceAccountUnauthorized with default headers values
func NewGetDefaultBalanceAccountUnauthorized() *GetDefaultBalanceAccountUnauthorized {
	return &GetDefaultBalanceAccountUnauthorized{}
}

/*GetDefaultBalanceAccountUnauthorized handles this case with default header values.

User not authenticated
*/
type GetDefaultBalanceAccountUnauthorized struct {
	Payload *models.ErrorsDTO
}

func (o *GetDefaultBalanceAccountUnauthorized) Error() string {
	return fmt.Sprintf("[GET /default][%d] getDefaultBalanceAccountUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDefaultBalanceAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDefaultBalanceAccountForbidden creates a GetDefaultBalanceAccountForbidden with default headers values
func NewGetDefaultBalanceAccountForbidden() *GetDefaultBalanceAccountForbidden {
	return &GetDefaultBalanceAccountForbidden{}
}

/*GetDefaultBalanceAccountForbidden handles this case with default header values.

User not authorized
*/
type GetDefaultBalanceAccountForbidden struct {
	Payload *models.ErrorsDTO
}

func (o *GetDefaultBalanceAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /default][%d] getDefaultBalanceAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetDefaultBalanceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDefaultBalanceAccountNotFound creates a GetDefaultBalanceAccountNotFound with default headers values
func NewGetDefaultBalanceAccountNotFound() *GetDefaultBalanceAccountNotFound {
	return &GetDefaultBalanceAccountNotFound{}
}

/*GetDefaultBalanceAccountNotFound handles this case with default header values.

Resource not found
*/
type GetDefaultBalanceAccountNotFound struct {
	Payload *models.ErrorsDTO
}

func (o *GetDefaultBalanceAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /default][%d] getDefaultBalanceAccountNotFound  %+v", 404, o.Payload)
}

func (o *GetDefaultBalanceAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}