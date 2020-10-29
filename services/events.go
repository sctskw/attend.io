package services

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/sctskw/attend.io/models"
)

func GetAllEvents() []*models.Event {
	//TODO: models.NewEventList
	return []*models.Event{
		models.NewEvent("Event 1", strfmt.UUID(uuid.New().String()), nil),
		models.NewEvent("Event 2", strfmt.UUID(uuid.New().String()), nil),
	}
}

func GetEventById(id string) *models.Event {
	return models.NewEvent("Event 3", strfmt.UUID(uuid.New().String()), nil)
}
