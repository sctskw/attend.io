package services

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/sctskw/attend.io/db"
	"github.com/sctskw/attend.io/models"
)

type EventService interface {
	GetAll() models.EventList
	GetById(id strfmt.UUID) *models.Event
	GetAttendees(id strfmt.UUID) models.AttendeesList
}

type eventService struct {
	db db.DatabaseClient
}

func NewEventService(client db.DatabaseClient) EventService {
	return &eventService{db: client}
}

func (s *eventService) GetAll() models.EventList {
	return models.EventList{
		models.NewEvent("Event 1", strfmt.UUID(uuid.New().String()), nil),
		models.NewEvent("Event 2", strfmt.UUID(uuid.New().String()), nil),
	}
}

func (s *eventService) GetById(id strfmt.UUID) *models.Event {
	return models.NewEvent("Event 3", strfmt.UUID(uuid.New().String()), nil)
}

func (s *eventService) GetAttendees(id strfmt.UUID) models.AttendeesList {
	return models.NewAttendeeList(
		models.NewAttendee("John", "Smith", "john.smith@test.com"),
		models.NewAttendee("Jane", "Doe", "jane.doe@test.com"),
	)
}
