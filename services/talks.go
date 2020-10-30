package services

import (
	"github.com/sctskw/attend.io/db"
	"github.com/sctskw/attend.io/models"
)

type TalkService interface {
	GetAll() models.TalkList
	GetById(id string) *models.Talk
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
