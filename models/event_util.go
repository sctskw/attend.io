package models

import (
	"github.com/go-openapi/swag"
	"github.com/google/uuid"
)

func NewEvent(name string, talk string, attendees []*Attendee) *Event {
	return &Event{
		ID:   uuid.New().String(),
		Name: swag.String(name),
		Talk: talk,
	}
}

func NewEventList(e ...*Event) EventList {
	events := EventList{}
	events = append(events, e...)
	return events
}
