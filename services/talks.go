package services

import (
	"fmt"

	"github.com/sctskw/attend.io/db"
	"github.com/sctskw/attend.io/models"
)

type TalkService interface {
	GetAll() models.TalkList
	GetById(id string) (*models.Talk, error)
	GetAllById(ids ...string) (results models.TalkList)
	GetAttendees(id string) (models.AttendeeList, error)
	Create(m *models.Talk) (*models.Talk, error)
	AddAttendee(id string, attendees []string) (*models.Talk, error)
	RemoveAttendee(id, attendee string) error
	Delete(id string) error
}

type talkService struct {
	db db.DatabaseClient
}

func NewTalkService(client db.DatabaseClient) TalkService {
	return &talkService{db: client}
}

func (s *talkService) GetAll() models.TalkList {
	data := models.TalkList{}

	raw, err := s.db.FetchAll("talks")

	if err != nil {
		return data
	}

	for _, r := range raw {
		t := models.Talk{}
		err := t.UnmarshalBinary(r)

		if err != nil {
			continue
		}

		data = append(data, &t)
	}

	return data
}

func (s *talkService) GetById(id string) (*models.Talk, error) {
	raw, err := s.db.FetchById("talks", id)

	if err != nil {
		return nil, err
	}

	t := &models.Talk{}
	err = t.UnmarshalBinary(raw)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *talkService) GetAllById(ids ...string) (results models.TalkList) {

	talks, err := s.db.FetchAllById("talks", ids...)

	if err != nil {
		return results
	}

	for _, raw := range talks {
		t := &models.Talk{}
		err := t.UnmarshalBinary(raw)

		if err != nil {
			continue
		}

		results = append(results, t)
	}

	return results
}

func (s *talkService) GetAttendees(id string) (models.AttendeeList, error) {
	talk, err := s.GetById(id)

	if err != nil {
		return nil, err
	}

	return talk.RefAttendees, nil
}

func (s *talkService) Create(m *models.Talk) (*models.Talk, error) {

	//ensure this field is created
	data, _ := m.MarshalBinary()
	res, err := s.db.Insert("talks", data)

	if err != nil {
		return nil, err
	}

	t := &models.Talk{}
	err = t.UnmarshalBinary(res)

	if err != nil {
		return nil, err
	}

	return s.GetById(t.RefID)
}

func (s *talkService) Delete(id string) error {
	s.db.DeleteById("talks", id)
	return nil
}

func (s *talkService) AddAttendee(id string, attendees []string) (*models.Talk, error) {

	talk, err := s.GetById(id)

	fmt.Println(fmt.Sprintf("Found Talk: %s", id))

	if err != nil {
		return nil, err
	}

	refs := Attendees().GetAllById(attendees...)
	talk.RefAttendees = append(talk.RefAttendees, refs...)

	fmt.Println(fmt.Sprintf("Updated Talk: %+v", talk))

	b, _ := talk.MarshalBinary()

	res, err := s.db.Update("talks", id, b)

	if err != nil {
		return nil, err
	}

	//associate the Talk to the Attendee
	for _, a := range refs {
		_, err := Attendees().JoinTalk(a.RefID, id)

		//TODO: handle this
		if err != nil {
			continue
		}
	}

	t := &models.Talk{}
	err = t.UnmarshalBinary(res)

	if err != nil {
		return nil, err
	}

	return t, nil

}

func (s *talkService) RemoveAttendee(id, aid string) error {

	talk, err := s.GetById(id)

	if err != nil {
		return err
	}

	//remove the Attendee
	//TODO: this isn't efficient. could be a go routine or an async background task since it can happen later
	updated := models.AttendeeList{}
	for _, ref := range talk.RefAttendees {
		if ref.RefID != aid {
			updated = append(updated, ref)
		}
	}
	talk.RefAttendees = updated //update the list

	b, _ := talk.MarshalBinary()

	_, err = s.db.Update("talks", id, b)

	return err
}
