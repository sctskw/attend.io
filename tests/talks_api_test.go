package tests

import (
	"fmt"
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
		res, err := s.Update(fmt.Sprintf("/talks/%s/attendees", talk.ID), []string{"zszPSgO73zCcEUEjdRoq"}, &talk)
		s.Assert().Nil(err)
		s.Assert().Equal(200, res.StatusCode)
		s.Assert().Len(talk.RefAttendees, 1)

	})

	//s.Run("delete a Talk", func() {
	//	s.Delete("/talks/" + talk.ID)
	//})

}
