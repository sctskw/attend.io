package models

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

func NewTalk(name, presenter, description string, start, end time.Time) *Talk {
	return &Talk{
		Name:          swag.String(name),
		Presenter:     presenter,
		Description:   description,
		DateTimeEnd:   strfmt.DateTime(end),
		DateTimeStart: strfmt.DateTime(start),
	}
}

func (m *Talk) GetAttendeeIds() (ids []string) {
	for _, a := range m.RefAttendees {
		ids = append(ids, a.RefID)
	}
	return ids
}
