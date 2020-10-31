package tests

import (
	"time"

	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestTalksAPI_GetTalks() {

	talks := make(models.TalkList, 0)

	s.Fetch("/talks", &talks)
	//s.Assert().Len(talks, 2, "returns total number of talks")

	for _, talk := range talks {
		s.Assert().NotNil(talk.ID, "has an ID")
		s.Assert().NotEqual("", talk.ID, "has an valid ID")
		s.Assert().NotNil(talk.Name, "has a name")
		s.Assert().NotNil(talk.Presenter, "Presenter")
		s.Assert().NotNil(talk.Description, "has a description")
	}
}

func (s *ApiTestSuite) TestTalksAPI_GetTalkByID() {

	talk := &models.Talk{}
	s.Fetch("/talks/HclaPWsc4TNfmTbVYNRy", talk)
	s.Assert().NotEqual("", talk.ID, "has an valid ID")
	s.Assert().Equal("Talk 2", *talk.Name)
}

func (s *ApiTestSuite) TestTalksAPI_GetTalkAttendees() {

	attendees := models.NewAttendeeList()
	s.Fetch("/talks/HclaPWsc4TNfmTbVYNRy/attendees", &attendees)
	s.Assert().Len(attendees, 3, "returns total number of attendees")

	for _, a := range attendees {
		s.Assert().NotEqual("", a.ID, "has an valid ID")
		s.Assert().NotEmpty(a.Email)
		s.Assert().NotEmpty(a.NameFirst)
		s.Assert().NotEmpty(a.NameLast)
	}
}

func (s *ApiTestSuite) TestTalksAPI_CreateTalk() {

	t := models.NewTalk(
		"TestTalkX",
		"Helen Hunt",
		"its a twister",
		time.Now(),
		time.Now(),
	)

	talk := models.Talk{}
	res := s.Create("/talks", t, &talk)

	s.Assert().Equal(200, res.StatusCode)
	s.Assert().Empty(t.ID)
	s.Assert().NotEmpty(talk.ID)
	s.Assert().Equal(t.Name, talk.Name)
	s.Assert().Equal(t.Presenter, talk.Presenter)
	s.Assert().Equal(t.Description, talk.Description)

}
