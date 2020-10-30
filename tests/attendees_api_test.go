package tests

import (
	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestAttendeesAPI_GetAttendeeById() {

	attendee := models.Attendee{}

	s.Fetch("/attendees?id=0YB7Cy6RC9fY5FMohK0p", &attendee)
	s.Assert().NotNil(attendee, "returns an attendee")
	s.Assert().Equal("bob.smith@test.com", attendee.Email.String())

}

func (s *ApiTestSuite) TestAttendeesAPI_GetAttendeeByEmail() {

	attendee := models.Attendee{}

	s.Fetch("/attendees?email=sam.wise@test.com", &attendee)
	s.Assert().NotNil(attendee, "returns an attendee")
	s.Assert().Equal("sam.wise@test.com", attendee.Email.String())

}
