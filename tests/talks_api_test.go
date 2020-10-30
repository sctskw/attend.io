package tests

import (
	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestTalksAPI_GetTalks() {

	talks := make(models.TalkList, 2)

	s.Fetch("/talks", &talks)
	s.Assert().Len(talks, 2, "returns total number of talks")

	for _, talk := range talks {
		s.Assert().NotNil(talk.ID, "has an ID")
	}
}

func (s *ApiTestSuite) TestTalksAPI_GetTalkByID() {
	talk := models.Talk{}

	s.Fetch("/talks/d0e7a628-1a7b-11eb-adc1-0242ac120002", &talk)
	s.Assert().Equal("Talk 3", *talk.Name)
}
