package tests

import "github.com/sctskw/attend.io/models"

func (s *ApiTestSuite) TestAttendeesAPI() {

	a := models.NewAttendee(
		"Bob",
		"Smith",
		"bob.smith@test.com",
	)

	attendee := models.Attendee{}

	s.Run("create an Attendee", func() {
		res := s.Create("/attendees", a, &attendee)

		s.Assert().Equal(200, res.StatusCode)
		s.Assert().Empty(a.ID)
		s.Assert().NotEmpty(attendee.ID)
		s.Assert().Equal(a.NameFirst, attendee.NameFirst)
		s.Assert().Equal(a.NameLast, attendee.NameLast)
		s.Assert().Equal(a.Email, attendee.Email)
	})

	s.Run("get Attendee by ID", func() {
		res := s.Fetch("/attendees?id="+attendee.ID, &attendee)
		s.Assert().Equal(200, res.StatusCode)
	})

	s.Run("get Attendee by Email", func() {
		res := s.Fetch("/attendees?email="+attendee.Email.String(), &attendee)
		s.Assert().Equal(200, res.StatusCode)
	})

	s.Run("delete an Attendee", func() {
		res := s.Delete("/attendees/" + attendee.ID)
		s.Assert().Equal(204, res.StatusCode)
	})

}
