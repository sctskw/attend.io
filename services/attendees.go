package services

import (
	"github.com/go-openapi/strfmt"
	"github.com/sctskw/attend.io/db"
	"github.com/sctskw/attend.io/models"
)

type AttendeeService interface {
	GetAllRefsById(ids ...string) []string
	GetAllById(ids ...string) models.AttendeeList
	GetById(id string) (*models.Attendee, error)
	GetByEmail(email strfmt.Email) (*models.Attendee, error)
	Create(m *models.Attendee) (*models.Attendee, error)
}

type attendeeService struct {
	db   db.DatabaseClient
	name string
}

func NewAttendeeService(client db.DatabaseClient) AttendeeService {
	return &attendeeService{db: client, name: "attendees"}
}

func (s *attendeeService) GetById(id string) (*models.Attendee, error) {
	raw, err := s.db.FetchById(s.name, id)

	a := &models.Attendee{}
	err = a.UnmarshalBinary(raw)

	if err != nil {
		return nil, err
	}

	return a, nil

}

func (s *attendeeService) GetByEmail(email strfmt.Email) (*models.Attendee, error) {
	raw, err := s.db.FetchByField(s.name, "email", email.String())

	if err != nil {
		return nil, err
	}

	a := &models.Attendee{}
	err = a.UnmarshalBinary(raw)

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *attendeeService) GetAllRefsById(ids ...string) (results []string) {

	attendees, err := s.db.FetchAllById("attendees", ids...)

	if err != nil {
		return results
	}

	for _, raw := range attendees {
		a := &models.Attendee{}
		err := a.UnmarshalBinary(raw)

		if err != nil {
			continue
		}

		results = append(results, a.Ref)
	}

	return results
}

func (s *attendeeService) GetAllById(ids ...string) (results models.AttendeeList) {

	attendees, err := s.db.FetchAllById("attendees", ids...)

	if err != nil {
		return results
	}

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

func (s *attendeeService) Create(m *models.Attendee) (*models.Attendee, error) {

	//ensure this field is created
	data, _ := m.MarshalBinary()
	res, err := s.db.Insert("attendees", data)

	if err != nil {
		return nil, err
	}

	a := &models.Attendee{}
	err = a.UnmarshalBinary(res)

	if err != nil {
		return nil, err
	}

	return s.GetById(a.ID)
}
