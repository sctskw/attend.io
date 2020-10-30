package services

import (
	"github.com/go-openapi/strfmt"
	"github.com/sctskw/attend.io/models"
)

func GetAttendeesByEventId(eventId strfmt.UUID) models.AttendeesList {
	return models.NewAttendeeList(
		models.NewAttendee("John", "Smith", "john.smith@test.com"),
		models.NewAttendee("Jane", "Doe", "jane.doe@test.com"),
	)
}

func GetAttendeeById(id strfmt.UUID) *models.Attendee {
	return models.NewAttendee("Bob", "Stevens", "bob@test.com")
}

func GetAttendeeByEmail(email strfmt.Email) *models.Attendee {
	return models.NewAttendee("Sam", "Wise", "sam.wise@test.com")
}
