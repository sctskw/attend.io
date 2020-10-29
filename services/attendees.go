package services

import (
	"github.com/sctskw/attend.io/models"
)

func GetAttendeesForEvent(eventId string) *models.AttendeesList {
	return models.NewAttendeeList(
		models.NewAttendee("John", "Smith", "john.smith@test.com"),
		models.NewAttendee("Jane", "Doe", "jane.doe@test.com"),
	)
}

func GetAttendeeById(id string) *models.Attendee {
	return models.NewAttendee("Bob", "Stevens", "bob@test.com")
}

func GetAttendeeByEmail(email string) *models.Attendee {
	return models.NewAttendee("Sam", "Wise", "same.wise@test.com")
}
