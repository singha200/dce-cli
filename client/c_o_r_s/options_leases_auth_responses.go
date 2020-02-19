// Code generated by go-swagger; DO NOT EDIT.

package c_o_r_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// OptionsLeasesAuthReader is a Reader for the OptionsLeasesAuth structure.
type OptionsLeasesAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OptionsLeasesAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOptionsLeasesAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOptionsLeasesAuthOK creates a OptionsLeasesAuthOK with default headers values
func NewOptionsLeasesAuthOK() *OptionsLeasesAuthOK {
	return &OptionsLeasesAuthOK{}
}

/*OptionsLeasesAuthOK handles this case with default header values.

Default response for CORS method
*/
type OptionsLeasesAuthOK struct {
	AccessControlAllowHeaders string

	AccessControlAllowMethods string

	AccessControlAllowOrigin string
}

func (o *OptionsLeasesAuthOK) Error() string {
	return fmt.Sprintf("[OPTIONS /leases/auth][%d] optionsLeasesAuthOK ", 200)
}

func (o *OptionsLeasesAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Allow-Headers
	o.AccessControlAllowHeaders = response.GetHeader("Access-Control-Allow-Headers")

	// response header Access-Control-Allow-Methods
	o.AccessControlAllowMethods = response.GetHeader("Access-Control-Allow-Methods")

	// response header Access-Control-Allow-Origin
	o.AccessControlAllowOrigin = response.GetHeader("Access-Control-Allow-Origin")

	return nil
}
