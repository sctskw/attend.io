package services

import (
	"github.com/sctskw/attend.io/db"
)

type ServiceRegistry interface {
	Attendees() AttendeeService
	Talks() TalkService
	Events() EventService
	System() SystemService
}

type servicesRegistry struct {
	dbClient db.DatabaseClient
}

func NewServiceRegistry(dbClient db.DatabaseClient) ServiceRegistry {
	return &servicesRegistry{dbClient: dbClient}
}

func (r *servicesRegistry) Attendees() AttendeeService {
	return NewAttendeeService(r.dbClient)
}

func (r *servicesRegistry) Talks() TalkService {
	return NewTalkService(r.dbClient)
}

func (r *servicesRegistry) Events() EventService {
	return NewEventService(r.dbClient)
}

func (r *servicesRegistry) System() SystemService {
	return NewSystemService(r.dbClient)
}

//shortcuts
func Attendees() AttendeeService {
	return Get().Attendees()
}

func Talks() TalkService {
	return Get().Talks()
}

func Events() EventService {
	return Get().Events()
}

func System() SystemService {
	return Get().System()
}
