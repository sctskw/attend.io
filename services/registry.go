package services

import (
	"github.com/sctskw/attend.io/db"
)

type ServiceRegistry interface {
	Attendees() AttendeeService
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
