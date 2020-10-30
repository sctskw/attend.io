package services

import (
	"github.com/sctskw/attend.io/db"
)

type ServiceRegistry struct {
	db        db.DatabaseClient
	Attendees AttendeeService
	Events    EventService
	System    SystemService
	Talks     TalkService
}

func NewServiceRegistry(client db.DatabaseClient) *ServiceRegistry {
	return &ServiceRegistry{
		db:        client,
		Attendees: NewAttendeeService(client),
		Events:    NewEventService(client),
		System:    NewSystemService(client),
		Talks:     NewTalkService(client),
	}
}

func Attendees() AttendeeService {
	return Get().Attendees
}

func Events() EventService {
	return Get().Events
}

func System() SystemService {
	return Get().System
}

func Talks() TalkService {
	return Get().Talks
}
