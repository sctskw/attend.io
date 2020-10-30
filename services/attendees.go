package services

import (
	"github.com/go-openapi/strfmt"
	"github.com/sctskw/attend.io/db"
	"github.com/sctskw/attend.io/models"
)

type AttendeeService interface {
	GetByEventId(eventId strfmt.UUID) models.AttendeesList
	GetById(id strfmt.UUID) *models.Attendee
	GetByEmail(email strfmt.Email) *models.Attendee
}

type attendeeService struct {
	db db.DatabaseClient
}

func NewAttendeeService(client db.DatabaseClient) AttendeeService {
	return &attendeeService{db: client}
}

func (s *attendeeService) GetName() string {
	return "Attendees"
}

func (s *attendeeService) GetByEventId(eventId strfmt.UUID) models.AttendeesList {
	return models.NewAttendeeList(
		models.NewAttendee("John", "Smith", "john.smith@test.com"),
		models.NewAttendee("Jane", "Doe", "jane.doe@test.com"),
	)
}

func (s *attendeeService) GetById(id strfmt.UUID) *models.Attendee {
	return models.NewAttendee("Bob", "Stevens", "bob@test.com")
}

func (s *attendeeService) GetByEmail(email strfmt.Email) *models.Attendee {
	return models.NewAttendee("Sam", "Wise", "sam.wise@test.com")
}
