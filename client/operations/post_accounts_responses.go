// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// PostAccountsReader is a Reader for the PostAccounts structure.
type PostAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAccountsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPostAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAccountsCreated creates a PostAccountsCreated with default headers values
func NewPostAccountsCreated() *PostAccountsCreated {
	return &PostAccountsCreated{}
}

/*PostAccountsCreated handles this case with default header values.

Account Details
*/
type PostAccountsCreated struct {
	AccessControlAllowHeaders string

	AccessControlAllowMethods string

	AccessControlAllowOrigin string
}

func (o *PostAccountsCreated) Error() string {
	return fmt.Sprintf("[POST /accounts][%d] postAccountsCreated ", 201)
}

func (o *PostAccountsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Allow-Headers
	o.AccessControlAllowHeaders = response.GetHeader("Access-Control-Allow-Headers")

	// response header Access-Control-Allow-Methods
	o.AccessControlAllowMethods = response.GetHeader("Access-Control-Allow-Methods")

	// response header Access-Control-Allow-Origin
	o.AccessControlAllowOrigin = response.GetHeader("Access-Control-Allow-Origin")

	return nil
}

// NewPostAccountsForbidden creates a PostAccountsForbidden with default headers values
func NewPostAccountsForbidden() *PostAccountsForbidden {
	return &PostAccountsForbidden{}
}

/*PostAccountsForbidden handles this case with default header values.

Failed to authenticate request
*/
type PostAccountsForbidden struct {
}

func (o *PostAccountsForbidden) Error() string {
	return fmt.Sprintf("[POST /accounts][%d] postAccountsForbidden ", 403)
}

func (o *PostAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostAccountsBody post accounts body
swagger:model PostAccountsBody
*/
type PostAccountsBody struct {

	// ARN for an IAM role within this AWS account. The DCE master account will assume this IAM role to execute operations within this AWS account. This IAM role is configured by the client, and must be configured with [a Trust Relationship with the DCE master account.](/https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html)
	//
	// Required: true
	AdminRoleArn *string `json:"adminRoleArn"`

	// AWS Account ID
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this post accounts body
func (o *PostAccountsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAdminRoleArn(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostAccountsBody) validateAdminRoleArn(formats strfmt.Registry) error {

	if err := validate.Required("account"+"."+"adminRoleArn", "body", o.AdminRoleArn); err != nil {
		return err
	}

	return nil
}

func (o *PostAccountsBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("account"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostAccountsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAccountsBody) UnmarshalBinary(b []byte) error {
	var res PostAccountsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}