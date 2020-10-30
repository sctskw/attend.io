package services

import (
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
	raw := s.db.FetchAll("events")
	data := models.EventList{}

	for _, r := range raw {
		t := &models.Event{}
		err := t.UnmarshalBinary(r)

		//TODO: this could create data inconsistency
		if err != nil {
			panic(err)
		}

		data = append(data, t)
	}

	return data
}

func (s *eventService) GetById(id string) *models.Event {
	raw := s.db.FetchById("events", id)

	t := &models.Event{}
	err := t.UnmarshalBinary(raw)

	if err != nil {
		panic(err)
	}

	return t
}

func (s *eventService) GetAttendees(id string) models.AttendeesList {
	ids := make([]string, 0)
	event := s.GetById(id)

	if event == nil {
		return nil
	}

	for _, ref := range event.RefAttendees {
		ids = append(ids, ref.ID)
	}

	return Attendees().GetAllById(ids...)
}
