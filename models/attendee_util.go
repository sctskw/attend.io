package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

func NewAttendee(nameFirst, nameLast, email string) *Attendee {
	return &Attendee{
		NameDisplay: nameFirst,
		NameFirst:   swag.String(nameFirst),
		NameLast:    swag.String(nameLast),
		Email:       strfmt.Email(email),
	}
}

func (m *Attendee) GetTalkIds() (ids []string) {
	for _, t := range m.RefTalks {
		ids = append(ids, t.RefID)

	}
	return ids
}

func (m *Attendee) JoinTalks(talks TalkList) *Attendee {

	//TODO: need to check for duplicates
	for _, t := range talks {
		t.RefAttendees = nil //clean these up so we 're not overloading documents.
		m.RefTalks = append(m.RefTalks, t)
	}

	return m
}

func (m *Attendee) LeaveTalks(talks TalkList) *Attendee {
	//clean up talks
	//TODO: move to Attendee model
	updated := TalkList{}
	for _, talk := range m.RefTalks {
		for _, leaveTalk := range talks {
			if leaveTalk.RefID != talk.RefID {
				talk.RefAttendees = nil //sanitize
				updated = append(updated, talk)
			}
		}
	}

	//remove
	m.RefTalks = updated

	return m

}
