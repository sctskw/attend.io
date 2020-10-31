package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

func NewAttendee(nameFirst, nameLast, email string) *Attendee {
	return &Attendee{
		NameDisplay: nameFirst,
		NameFirst:   swag.String(nameFirst),
		NameLast:    swag.String(nameLast),
		Email:       strfmt.Email(email),
	}
}

func NewAttendeeList(attendees ...*Attendee) AttendeeList {
	l := AttendeeList{}
	l = append(l, attendees...)
	return l
}

func NewAttendeesListFromIds(ids ...string) AttendeeList {
	l := AttendeeList{}
	for _, id := range ids {
		l = append(l, &Attendee{ID: id})
	}
	return l
}
