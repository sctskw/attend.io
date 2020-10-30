package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/google/uuid"
)

func NewEvent(name string, talk strfmt.UUID, attendees []*Attendee) *Event {
	return &Event{
		ID:   strfmt.UUID(uuid.New().String()),
		Name: swag.String(name),
		Talk: talk,
	}
}

func NewEventList(e ...*Event) EventList {
	events := EventList{}
	events = append(events, e...)
	return events
}
