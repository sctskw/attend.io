package tests

import (
	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestTalksAPI_GetTalks() {

	talks := make([]*models.Talk, 2)

	s.Fetch("/talks", &talks)

	s.Assert().Len(talks, 2, "returns total number of talks")

	for _, talk := range talks {
		s.Assert().NotNil(talk.ID, "has an ID")
	}
}
