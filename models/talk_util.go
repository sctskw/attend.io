package models

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

func NewTalk(name, presenter, description string, start, end time.Time) *Talk {
	return &Talk{
		Name:          swag.String(name),
		Presenter:     presenter,
		Description:   description,
		DateTimeEnd:   strfmt.DateTime(end),
		DateTimeStart: strfmt.DateTime(start),
	}
}

func NewTalkList(t ...*Talk) TalkList {
	list := TalkList{}
	list = append(list, t...)
	return list
}
