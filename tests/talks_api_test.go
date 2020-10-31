package tests

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestTalksAPI() {

	t := models.NewTalk(
		"Storm Chasers",
		"Helen Hunt",
		"its a twister",
		time.Now(),
		time.Now(),
	)
	talk := models.Talk{}
	talks := models.TalkList{}

	a := models.NewAttendee(
		"John",
		"Henry",
		"john.henry@test.com",
	)
	attendee := models.Attendee{}

	s.Run("create a Talk", func() {
		res := s.Create("/talks", t, &talk)
		s.Assert().Equal(200, res.StatusCode)
		s.Assert().Empty(t.ID)
		s.Assert().NotEmpty(talk.ID)
		s.Assert().Equal(t.Name, talk.Name)
		s.Assert().Equal(t.Presenter, talk.Presenter)
		s.Assert().Equal(t.Description, talk.Description)
	})

	s.Run("fetch all Talks", func() {
		res := s.Fetch("/talks", &talks)
		s.Assert().Equal(200, res.StatusCode)
		s.Assert().GreaterOrEqual(len(talks), 1)
	})

	s.Run("fetch Talk by ID", func() {
		res := s.Fetch("/talks/"+talk.ID, &talk)
		s.Assert().Equal(200, res.StatusCode)
	})

	s.Run("add Attendee to Talk", func() {

		//need an Attendee first
		res := s.Create("/attendees", a, &attendee)
		s.Assert().Equal(200, res.StatusCode)
		s.Assert().NotEmpty(attendee.ID)

		//update the TalkList
		res, err := s.Update(fmt.Sprintf("/talks/%s/attendees", talk.ID), []string{attendee.ID}, &talk)
		s.Assert().Nil(err)
		s.Assert().Equal(200, res.StatusCode)
		s.Assert().Len(talk.RefAttendees, 1)
	})

	s.Run("remove an Attendee from a Talk", func() {

		var res *http.Response

		//update the TalkList
		res = s.Delete(fmt.Sprintf("/talks/%s/attendees/%s", talk.ID, attendee.ID))
		s.Assert().Equal(204, res.StatusCode)

		t := models.Talk{}
		res = s.Fetch("/talks/"+talk.ID, &t)
		s.Assert().Equal(200, res.StatusCode)
		s.Assert().Len(t.RefAttendees, 0)

	})

	s.Run("delete a Talk", func() {
		s.Delete("/talks/" + talk.ID)
		s.Delete("/attendees/" + attendee.ID)
	})

}
