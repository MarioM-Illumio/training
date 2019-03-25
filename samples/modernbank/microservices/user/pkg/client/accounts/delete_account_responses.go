// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAccountReader is a Reader for the DeleteAccount structure.
type DeleteAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAccountOK creates a DeleteAccountOK with default headers values
func NewDeleteAccountOK() *DeleteAccountOK {
	return &DeleteAccountOK{}
}

/*DeleteAccountOK handles this case with default header values.

Aaaaand it's gone!
*/
type DeleteAccountOK struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *DeleteAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /users/{owner}/accounts/{number}][%d] deleteAccountOK ", 200)
}

func (o *DeleteAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}

// NewDeleteAccountNotFound creates a DeleteAccountNotFound with default headers values
func NewDeleteAccountNotFound() *DeleteAccountNotFound {
	return &DeleteAccountNotFound{}
}

/*DeleteAccountNotFound handles this case with default header values.

Account not found
*/
type DeleteAccountNotFound struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *DeleteAccountNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{owner}/accounts/{number}][%d] deleteAccountNotFound ", 404)
}

func (o *DeleteAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}

// NewDeleteAccountInternalServerError creates a DeleteAccountInternalServerError with default headers values
func NewDeleteAccountInternalServerError() *DeleteAccountInternalServerError {
	return &DeleteAccountInternalServerError{}
}

/*DeleteAccountInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteAccountInternalServerError struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *DeleteAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /users/{owner}/accounts/{number}][%d] deleteAccountInternalServerError ", 500)
}

func (o *DeleteAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}
