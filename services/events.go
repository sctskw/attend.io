package services

import (
	"github.com/google/uuid"
	"github.com/sctskw/attend.io/db"
	"github.com/sctskw/attend.io/models"
)

type EventService interface {
	GetAll() models.EventList
	GetById(id string) *models.Event
	GetAttendees(id string) models.AttendeesList
}

type eventService struct {
	db db.DatabaseClient
}

func NewEventService(client db.DatabaseClient) EventService {
	return &eventService{db: client}
}

func (s *eventService) GetAll() models.EventList {
	return models.EventList{
		models.NewEvent("Event 1", string(uuid.New().String()), nil),
		models.NewEvent("Event 2", string(uuid.New().String()), nil),
	}
}

func (s *eventService) GetById(id string) *models.Event {
	return models.NewEvent("Event 3", string(uuid.New().String()), nil)
}

func (s *eventService) GetAttendees(id string) models.AttendeesList {
	return models.NewAttendeeList(
		models.NewAttendee("John", "Smith", "john.smith@test.com"),
		models.NewAttendee("Jane", "Doe", "jane.doe@test.com"),
	)
}
