package tests

import (
	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestAttendeesAPI_GetAttendeeById() {

	attendee := models.Attendee{}

	s.Fetch("/attendees?id=4dd1fae8-93ed-4a0d-a7d4-518e42488633", &attendee)
	s.Assert().NotNil(attendee, "returns an attendee")
	s.Assert().Equal("bob@test.com", attendee.Email.String())

}

func (s *ApiTestSuite) TestAttendeesAPI_GetAttendeeByEmail() {

	attendee := models.Attendee{}

	s.Fetch("/attendees?email=sam.wise@test.com", &attendee)
	s.Assert().NotNil(attendee, "returns an attendee")
	s.Assert().Equal("sam.wise@test.com", attendee.Email.String())

}
