package services

import (
	"github.com/sctskw/attend.io/db"
	"github.com/sctskw/attend.io/models"
)

type TalkService interface {
	Create(m *models.Talk) *models.Talk
	Delete(id string)
	GetAll() models.TalkList
	GetById(id string) *models.Talk
	GetAttendees(id string) models.AttendeesList
}

type talkService struct {
	db db.DatabaseClient
}

func NewTalkService(client db.DatabaseClient) TalkService {
	return &talkService{db: client}
}

func (s *talkService) GetAll() models.TalkList {
	raw := s.db.FetchAll("talks")
	data := models.TalkList{}

	for _, r := range raw {
		t := &models.Talk{}
		err := t.UnmarshalBinary(r)

		//TODO: this could create data inconsistency
		if err != nil {
			panic(err)
		}

		data = append(data, t)
	}

	return data
}

func (s *talkService) GetById(id string) *models.Talk {
	raw := s.db.FetchById("talks", id)

	t := &models.Talk{}
	err := t.UnmarshalBinary(raw)

	if err != nil {
		panic(err)
	}

	return t
}

func (s *talkService) GetAttendees(id string) models.AttendeesList {
	ids := make([]string, 0)

	talk := s.GetById(id)

	for _, ref := range talk.RefAttendees {
		ids = append(ids, ref.ID)
	}

	return Attendees().GetAllById(ids...)
}

func (s *talkService) Create(m *models.Talk) *models.Talk {
	_ = s.db.Insert("talks", m)
	return nil
}

func (s *talkService) Delete(id string) {

}
