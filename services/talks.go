package services

import (
	"github.com/go-openapi/strfmt"
	"github.com/sctskw/attend.io/db"
	"github.com/sctskw/attend.io/models"
)

type TalkService interface {
	GetAll() models.TalkList
	GetById(id strfmt.UUID) *models.Talk
}

type talkService struct {
	dbClient db.DatabaseClient
}

func NewTalkService(client db.DatabaseClient) TalkService {
	return &talkService{dbClient: client}
}

func (s *talkService) GetAll() models.TalkList {
	return models.NewTalkList(
		models.NewTalk("Talk 1"),
		models.NewTalk("Talk 2"),
	)
}

func (s *talkService) GetById(id strfmt.UUID) *models.Talk {
	return models.NewTalk("Talk 3")
}
