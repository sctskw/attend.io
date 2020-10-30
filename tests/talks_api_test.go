package tests

import (
	"github.com/sctskw/attend.io/models"

	"github.com/stretchr/testify/assert"
)

func (s *ApiTestSuite) TestTalksAPI_GetTalks() {

	data := make([]*models.Talk, 2)

	s.Fetch("/talks", &data)

	assert.Len(s.T(), data, 2)

	for _, m := range data {
		assert.NotNil(s.T(), m.ID, "has an ID")
	}
}
