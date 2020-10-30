// Code generated by go-swagger; DO NOT EDIT.

package attendees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetAttendeeByFieldParams creates a new GetAttendeeByFieldParams object
// no default values defined in spec.
func NewGetAttendeeByFieldParams() GetAttendeeByFieldParams {

	return GetAttendeeByFieldParams{}
}

// GetAttendeeByFieldParams contains all the bound params for the get attendee by field operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAttendeeByField
type GetAttendeeByFieldParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Attendee Email
	  In: query
	*/
	Email *strfmt.Email
	/*Attendee ID
	  In: query
	*/
	ID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAttendeeByFieldParams() beforehand.
func (o *GetAttendeeByFieldParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qEmail, qhkEmail, _ := qs.GetOK("email")
	if err := o.bindEmail(qEmail, qhkEmail, route.Formats); err != nil {
		res = append(res, err)
	}

	qID, qhkID, _ := qs.GetOK("id")
	if err := o.bindID(qID, qhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEmail binds and validates parameter Email from query.
func (o *GetAttendeeByFieldParams) bindEmail(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: email
	value, err := formats.Parse("email", raw)
	if err != nil {
		return errors.InvalidType("email", "query", "strfmt.Email", raw)
	}
	o.Email = (value.(*strfmt.Email))

	if err := o.validateEmail(formats); err != nil {
		return err
	}

	return nil
}

// validateEmail carries on validations for parameter Email
func (o *GetAttendeeByFieldParams) validateEmail(formats strfmt.Registry) error {

	if err := validate.FormatOf("email", "query", "email", o.Email.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindID binds and validates parameter ID from query.
func (o *GetAttendeeByFieldParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ID = &raw

	return nil
}
