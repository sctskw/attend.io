package tests

import (
	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestTalksAPI_GetTalks() {

	talks := make(models.TalkList, 0)

	s.Fetch("/talks", &talks)
	s.Assert().Len(talks, 2, "returns total number of talks")

	for _, talk := range talks {
		s.Assert().NotNil(talk.ID, "has an ID")
		s.Assert().NotNil(talk.Name, "has a name")
		s.Assert().Contains(talk.Presenter, "Presenter")
		s.Assert().NotNil(talk.Description, "has a description")
	}
}

func (s *ApiTestSuite) TestTalksAPI_GetTalkByID() {

	talk := &models.Talk{}
	s.Fetch("/talks/OAUthvWOMLX1uOIvA41c", talk)
	s.Assert().Equal("Talk 1", *talk.Name)
}
