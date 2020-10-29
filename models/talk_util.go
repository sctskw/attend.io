package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/google/uuid"
)

func NewTalk(name string) *Talk {
	return &Talk{
		ID:   strfmt.UUID(uuid.New().String()),
		Name: swag.String(name),
	}
}

func (m *Talk) AddAttendee(a ...*Attendee) *Talk {
	m.Attendees = append(m.Attendees, a...)
	return m
}
