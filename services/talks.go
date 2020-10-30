package services

import (
	"github.com/go-openapi/strfmt"
	"github.com/sctskw/attend.io/models"
)

func GetTalks() models.TalkList {
	return models.NewTalkList(
		models.NewTalk("Talk 1"),
		models.NewTalk("Talk 2"),
	)
}

func GetTalkById(id strfmt.UUID) *models.Talk {
	return models.NewTalk("Talk 3")
}
