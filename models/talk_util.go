package models

import (
	"github.com/go-openapi/swag"
	"github.com/google/uuid"
)

func NewTalk(name string) *Talk {
	return &Talk{
		ID:   uuid.New().String(),
		Name: swag.String(name),
	}
}

func NewTalkList(t ...*Talk) TalkList {
	list := TalkList{}
	list = append(list, t...)
	return list
}
