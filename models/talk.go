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

// Talk talk
//
// swagger:model Talk
type Talk struct {

	// date time end
	// Format: date-time
	DateTimeEnd strfmt.DateTime `json:"date_time_end,omitempty"`

	// date time start
	// Format: date-time
	DateTimeStart strfmt.DateTime `json:"date_time_start,omitempty"`

	// description
	// Min Length: 1
	Description string `json:"description,omitempty"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// presenter
	// Min Length: 1
	Presenter string `json:"presenter,omitempty"`

	// ref attendees
	RefAttendees AttendeesList `json:"ref_attendees,omitempty"`

	// ref talk
	RefTalk *Talk `json:"ref_talk,omitempty"`
}

// Validate validates this talk
func (m *Talk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateTimeEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateTimeStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresenter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefAttendees(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefTalk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Talk) validateDateTimeEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTimeEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("date_time_end", "body", "date-time", m.DateTimeEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Talk) validateDateTimeStart(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTimeStart) { // not required
		return nil
	}

	if err := validate.FormatOf("date_time_start", "body", "date-time", m.DateTimeStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Talk) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *Talk) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *Talk) validatePresenter(formats strfmt.Registry) error {

	if swag.IsZero(m.Presenter) { // not required
		return nil
	}

	if err := validate.MinLength("presenter", "body", string(m.Presenter), 1); err != nil {
		return err
	}

	return nil
}

func (m *Talk) validateRefAttendees(formats strfmt.Registry) error {

	if swag.IsZero(m.RefAttendees) { // not required
		return nil
	}

	if err := m.RefAttendees.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ref_attendees")
		}
		return err
	}

	return nil
}

func (m *Talk) validateRefTalk(formats strfmt.Registry) error {

	if swag.IsZero(m.RefTalk) { // not required
		return nil
	}

	if m.RefTalk != nil {
		if err := m.RefTalk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ref_talk")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Talk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Talk) UnmarshalBinary(b []byte) error {
	var res Talk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
