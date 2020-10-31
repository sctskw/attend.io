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

func (m *Attendee) GetTalkIds() (ids []string) {
	for _, t := range m.RefTalks {
		ids = append(ids, t.ID)
	}
	return ids
}
