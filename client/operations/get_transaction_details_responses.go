// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/johnksv/sparebank1-personligklient/models"
)

// GetTransactionDetailsReader is a Reader for the GetTransactionDetails structure.
type GetTransactionDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetTransactionDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTransactionDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTransactionDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionDetailsOK creates a GetTransactionDetailsOK with default headers values
func NewGetTransactionDetailsOK() *GetTransactionDetailsOK {
	return &GetTransactionDetailsOK{}
}

/*GetTransactionDetailsOK handles this case with default header values.

Successful response
*/
type GetTransactionDetailsOK struct {
	Payload *models.TransactionDto
}

func (o *GetTransactionDetailsOK) Error() string {
	return fmt.Sprintf("[GET /{accountId}/transactions/{transactionId}][%d] getTransactionDetailsOK  %+v", 200, o.Payload)
}

func (o *GetTransactionDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransactionDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionDetailsUnauthorized creates a GetTransactionDetailsUnauthorized with default headers values
func NewGetTransactionDetailsUnauthorized() *GetTransactionDetailsUnauthorized {
	return &GetTransactionDetailsUnauthorized{}
}

/*GetTransactionDetailsUnauthorized handles this case with default header values.

User not authenticated
*/
type GetTransactionDetailsUnauthorized struct {
	Payload *models.ErrorsDTO
}

func (o *GetTransactionDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{accountId}/transactions/{transactionId}][%d] getTransactionDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTransactionDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionDetailsForbidden creates a GetTransactionDetailsForbidden with default headers values
func NewGetTransactionDetailsForbidden() *GetTransactionDetailsForbidden {
	return &GetTransactionDetailsForbidden{}
}

/*GetTransactionDetailsForbidden handles this case with default header values.

User not authorized
*/
type GetTransactionDetailsForbidden struct {
	Payload *models.ErrorsDTO
}

func (o *GetTransactionDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /{accountId}/transactions/{transactionId}][%d] getTransactionDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetTransactionDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionDetailsNotFound creates a GetTransactionDetailsNotFound with default headers values
func NewGetTransactionDetailsNotFound() *GetTransactionDetailsNotFound {
	return &GetTransactionDetailsNotFound{}
}

/*GetTransactionDetailsNotFound handles this case with default header values.

Resource not found
*/
type GetTransactionDetailsNotFound struct {
	Payload *models.ErrorsDTO
}

func (o *GetTransactionDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /{accountId}/transactions/{transactionId}][%d] getTransactionDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetTransactionDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
