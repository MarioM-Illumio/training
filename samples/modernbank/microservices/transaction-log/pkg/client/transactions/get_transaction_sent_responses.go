// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/tetrateio/training/samples/modernbank/microservices/transaction-log/pkg/model"
)

// GetTransactionSentReader is a Reader for the GetTransactionSent structure.
type GetTransactionSentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionSentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionSentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetTransactionSentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetTransactionSentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionSentOK creates a GetTransactionSentOK with default headers values
func NewGetTransactionSentOK() *GetTransactionSentOK {
	return &GetTransactionSentOK{}
}

/*GetTransactionSentOK handles this case with default header values.

Success!
*/
type GetTransactionSentOK struct {
	/*Version of the microservice that responded
	 */
	Version string

	Payload *model.Transaction
}

func (o *GetTransactionSentOK) Error() string {
	return fmt.Sprintf("[GET /transactions/account/{sender}/sent/{transaction}][%d] getTransactionSentOK  %+v", 200, o.Payload)
}

func (o *GetTransactionSentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	o.Payload = new(model.Transaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionSentNotFound creates a GetTransactionSentNotFound with default headers values
func NewGetTransactionSentNotFound() *GetTransactionSentNotFound {
	return &GetTransactionSentNotFound{}
}

/*GetTransactionSentNotFound handles this case with default header values.

Transaction not found
*/
type GetTransactionSentNotFound struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *GetTransactionSentNotFound) Error() string {
	return fmt.Sprintf("[GET /transactions/account/{sender}/sent/{transaction}][%d] getTransactionSentNotFound ", 404)
}

func (o *GetTransactionSentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}

// NewGetTransactionSentInternalServerError creates a GetTransactionSentInternalServerError with default headers values
func NewGetTransactionSentInternalServerError() *GetTransactionSentInternalServerError {
	return &GetTransactionSentInternalServerError{}
}

/*GetTransactionSentInternalServerError handles this case with default header values.

Internal server error
*/
type GetTransactionSentInternalServerError struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *GetTransactionSentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /transactions/account/{sender}/sent/{transaction}][%d] getTransactionSentInternalServerError ", 500)
}

func (o *GetTransactionSentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}
