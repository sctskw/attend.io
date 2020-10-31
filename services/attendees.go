package services

import (
	"github.com/go-openapi/strfmt"
	"github.com/sctskw/attend.io/db"
	"github.com/sctskw/attend.io/models"
)

type AttendeeService interface {
	GetAllById(ids ...string) models.AttendeesList
	GetById(id string) *models.Attendee
	GetByEmail(email strfmt.Email) *models.Attendee
	Create(m *models.Attendee) *models.Attendee
}

type attendeeService struct {
	db   db.DatabaseClient
	name string
}

func NewAttendeeService(client db.DatabaseClient) AttendeeService {
	return &attendeeService{db: client, name: "attendees"}
}

func (s *attendeeService) GetById(id string) *models.Attendee {
	raw := s.db.FetchById(s.name, id)

	a := &models.Attendee{}
	err := a.UnmarshalBinary(raw)

	if err != nil {
		panic(err)
	}

	return a

}

func (s *attendeeService) GetByEmail(email strfmt.Email) *models.Attendee {
	raw := s.db.FetchByField(s.name, "email", email.String())

	a := &models.Attendee{}
	err := a.UnmarshalBinary(raw)

	if err != nil {
		panic(err)
	}

	return a
}

func (s *attendeeService) GetAllById(ids ...string) (results models.AttendeesList) {

	attendees := s.db.FetchAllById("attendees", ids...)

	for _, raw := range attendees {
		a := &models.Attendee{}
		err := a.UnmarshalBinary(raw)

		if err != nil {
			continue
		}

		results = append(results, a)
	}

	return results
}

func (s *attendeeService) Create(m *models.Attendee) *models.Attendee {

	//ensure this field is created
	data, _ := m.MarshalBinary()
	res := s.db.Insert("attendees", data)

	//TODO
	if res == nil {
		return nil
	}

	a := &models.Attendee{}
	err := a.UnmarshalBinary(res)

	if err != nil {
		return nil
	}

	return s.GetById(a.ID)
}
