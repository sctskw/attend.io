package tests

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestAttendeesAPI() {

	a := models.NewAttendee(
		"Bob",
		"Smith",
		"bob.smith@test.com",
	)

	t := models.NewTalk(
		"Stunt Doubling",
		"Bruce Lee",
		"hollywood knockout",
		time.Now(),
		time.Now(),
	)
	talk := models.Talk{}

	attendee := models.Attendee{}

	s.Run("create an Attendee", func() {
		res := s.Create("/attendees", a, &attendee)
		s.Assert().Equal(200, res.StatusCode)
		s.Assert().Empty(a.RefID)
		s.Assert().NotEmpty(attendee.RefID)
		s.Assert().Equal(a.NameFirst, attendee.NameFirst)
		s.Assert().Equal(a.NameLast, attendee.NameLast)
		s.Assert().Equal(a.Email, attendee.Email)
	})

	s.Run("get Attendee by ID", func() {
		res := s.Fetch("/attendees?id="+attendee.RefID, &attendee)
		s.Assert().Equal(200, res.StatusCode)
	})

	s.Run("get Attendee by Email", func() {
		res := s.Fetch("/attendees?email="+attendee.Email.String(), &attendee)
		s.Assert().Equal(200, res.StatusCode)
	})

	s.Run("join a Talk", func() {
		res := s.Create("/talks", t, &talk)
		s.Assert().Equal(200, res.StatusCode)
		s.Assert().NotEmpty(talk.RefID)

		//update the TalkList
		res, err := s.Update(fmt.Sprintf("/talks/%s/attendees", talk.RefID), []string{attendee.RefID}, &talk)
		s.Assert().Nil(err)
		s.Assert().Equal(200, res.StatusCode)
		s.Assert().Len(talk.RefAttendees, 1)
		s.Assert().Equal(talk.RefAttendees[0].RefID, attendee.RefID)
	})

	s.Run("delete an Attendee", func() {

		var res *http.Response

		t := models.Talk{}

		res = s.Delete("/attendees/" + attendee.RefID)
		s.Assert().Equal(204, res.StatusCode)

		res = s.Fetch("/talks/"+talk.RefID, &t)
		s.Assert().Equal(200, res.StatusCode)
		s.Assert().Len(t.RefAttendees, 0)

		res = s.Delete("/talks/" + t.RefID)
		s.Assert().Equal(204, res.StatusCode)

	})

}
