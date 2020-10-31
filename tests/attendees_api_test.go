package tests

import (
	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestAttendeesAPI_GetAttendeeById() {

	attendee := models.Attendee{}

	s.Fetch("/attendees?id=azyFLMKTudmRfoBJjH5c", &attendee)
	s.Assert().NotNil(attendee, "returns an attendee")
	s.Assert().Equal("bob.smith@test.com", attendee.Email.String())
}

func (s *ApiTestSuite) TestAttendeesAPI_GetAttendeeByEmail() {

	attendee := models.Attendee{}

	s.Fetch("/attendees?email=sam.wise@test.com", &attendee)

	s.Assert().NotNil(attendee, "returns an attendee")
	s.Assert().Equal("sam.wise@test.com", attendee.Email.String())
}

//func (s *ApiTestSuite) TestAttendeesAPI_CreateAttendee() {
//
//	a := models.NewAttendee(
//		"Bob",
//		"Saget",
//		"bob.saget@test.com",
//	)
//
//	attendee := models.Attendee{}
//	res := s.Create("/attendees", a, &attendee)
//
//	s.Assert().Equal(200, res.StatusCode)
//	s.Assert().Empty(a.ID)
//	s.Assert().NotEmpty(attendee.ID)
//	s.Assert().Equal(a.NameFirst, attendee.NameFirst)
//	s.Assert().Equal(a.NameLast, attendee.NameLast)
//	s.Assert().Equal(a.Email, attendee.Email)
//
//}
