// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/tetrateio/training/samples/modernbank/microservices/user/pkg/model"
)

// ListAccountsReader is a Reader for the ListAccounts structure.
type ListAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewListAccountsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListAccountsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAccountsOK creates a ListAccountsOK with default headers values
func NewListAccountsOK() *ListAccountsOK {
	return &ListAccountsOK{}
}

/*ListAccountsOK handles this case with default header values.

Success!
*/
type ListAccountsOK struct {
	/*Version of the microservice that responded
	 */
	Version string

	Payload []*model.Account
}

func (o *ListAccountsOK) Error() string {
	return fmt.Sprintf("[GET /users/{username}/accounts][%d] listAccountsOK  %+v", 200, o.Payload)
}

func (o *ListAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountsNotFound creates a ListAccountsNotFound with default headers values
func NewListAccountsNotFound() *ListAccountsNotFound {
	return &ListAccountsNotFound{}
}

/*ListAccountsNotFound handles this case with default header values.

Owner not found
*/
type ListAccountsNotFound struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *ListAccountsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{username}/accounts][%d] listAccountsNotFound ", 404)
}

func (o *ListAccountsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}

// NewListAccountsInternalServerError creates a ListAccountsInternalServerError with default headers values
func NewListAccountsInternalServerError() *ListAccountsInternalServerError {
	return &ListAccountsInternalServerError{}
}

/*ListAccountsInternalServerError handles this case with default header values.

Internal server error
*/
type ListAccountsInternalServerError struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *ListAccountsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{username}/accounts][%d] listAccountsInternalServerError ", 500)
}

func (o *ListAccountsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}
