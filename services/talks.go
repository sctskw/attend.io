package services

import (
	"github.com/sctskw/attend.io/models"
)

func GetTalks() []*models.Talk {
	//TODO: models.NewTalkList
	return []*models.Talk{
		models.NewTalk("Talk 1"),
		models.NewTalk("Talk 2"),
	}
}
