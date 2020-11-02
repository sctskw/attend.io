package services

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/sctskw/attend.io/db"
	"github.com/sctskw/attend.io/models"
)

type AttendeeService interface {
	GetAllById(ids ...string) models.AttendeeList
	GetById(id string) (*models.Attendee, error)
	GetByEmail(email strfmt.Email) (*models.Attendee, error)
	Create(m *models.Attendee) (*models.Attendee, error)
	DeleteById(id string) error
	JoinTalk(id string, talks ...string) (*models.Attendee, error)
	LeaveTalk(id string, talks ...string) (*models.Attendee, error)
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

	return s.GetById(a.RefID)
}

func (s *attendeeService) DeleteById(id string) error {

	a, err := s.GetById(id)

	if err != nil {
		return err
	}

	//TODO:  check success before updating Talks
	s.db.DeleteById("attendees", id)

	//update the Talks
	//TODO: this isn't efficient. could be a go routine or an async background task since it can happen later
	for _, t := range a.GetTalkIds() {
		fmt.Println(fmt.Sprintf("%s DeJoining Talk: %s", id, t))
		err = Talks().RemoveAttendee(t, id)

		if err != nil {
			fmt.Errorf("could not dejoin: %s", err)
		}
	}

	return nil
}

func (s *attendeeService) JoinTalk(id string, talks ...string) (*models.Attendee, error) {

	attendee, err := s.GetById(id)

	if err != nil {
		return nil, err
	}

	refs := Talks().GetAllById(talks...)

	//sanitize
	for _, r := range refs {
		r.RefAttendees = nil //clean these up so we 're not overloading documents.
		attendee.RefTalks = append(attendee.RefTalks, r)
	}

	b, _ := attendee.MarshalBinary()

	res, err := s.db.Update("attendees", id, b)

	if err != nil {
		return nil, err
	}

	a := &models.Attendee{}
	err = a.UnmarshalBinary(res)

	if err != nil {
		return nil, err
	}

	return a, nil
}

//TODO: kind of duplicative of JoinTalk. Look to refactore
func (s *attendeeService) LeaveTalk(id string, talks ...string) (*models.Attendee, error) {

	attendee, err := s.GetById(id)

	if err != nil {
		return nil, err
	}

	refs := Talks().GetAllById(talks...)

	//remove from Talks first
	for _, ref := range refs {
		_ = Talks().RemoveAttendee(ref.RefID, attendee.RefID)
	}

	//clean up talks
	//TODO: move to Attendee model
	updated := models.TalkList{}
	for _, talk := range attendee.RefTalks {
		for _, ref := range refs {
			if ref.RefID != talk.RefID {
				talk.RefAttendees = nil //sanitize
				updated = append(updated, talk)
			}
		}
	}

	b, _ := attendee.MarshalBinary()

	res, err := s.db.Update("attendees", id, b)

	if err != nil {
		return nil, err
	}

	a := &models.Attendee{}
	err = a.UnmarshalBinary(res)

	if err != nil {
		return nil, err
	}

	return a, nil
}
