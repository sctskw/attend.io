package services

import (
	"github.com/sctskw/attend.io/db"
)

type ServiceRegistry struct {
	db        db.DatabaseClient
	Attendees AttendeeService
	System    SystemService
	Talks     TalkService
}

func NewServiceRegistry(client db.DatabaseClient) *ServiceRegistry {
	return &ServiceRegistry{
		db:        client,
		Attendees: NewAttendeeService(client),
		System:    NewSystemService(client),
		Talks:     NewTalkService(client),
	}
}

func Attendees() AttendeeService {
	return Get().Attendees
}

func System() SystemService {
	return Get().System
}

func Talks() TalkService {
	return Get().Talks
}
