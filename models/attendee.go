// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Attendee attendee
//
// swagger:model Attendee
type Attendee struct {

	// ID
	ID string `json:"ID,omitempty"`

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// name display
	// Min Length: 1
	NameDisplay string `json:"name_display,omitempty"`

	// name first
	// Required: true
	// Min Length: 1
	NameFirst *string `json:"name_first"`

	// name last
	// Required: true
	// Min Length: 1
	NameLast *string `json:"name_last"`
}

// Validate validates this attendee
func (m *Attendee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameLast(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Attendee) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Attendee) validateNameDisplay(formats strfmt.Registry) error {

	if swag.IsZero(m.NameDisplay) { // not required
		return nil
	}

	if err := validate.MinLength("name_display", "body", string(m.NameDisplay), 1); err != nil {
		return err
	}

	return nil
}

func (m *Attendee) validateNameFirst(formats strfmt.Registry) error {

	if err := validate.Required("name_first", "body", m.NameFirst); err != nil {
		return err
	}

	if err := validate.MinLength("name_first", "body", string(*m.NameFirst), 1); err != nil {
		return err
	}

	return nil
}

func (m *Attendee) validateNameLast(formats strfmt.Registry) error {

	if err := validate.Required("name_last", "body", m.NameLast); err != nil {
		return err
	}

	if err := validate.MinLength("name_last", "body", string(*m.NameLast), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Attendee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Attendee) UnmarshalBinary(b []byte) error {
	var res Attendee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
