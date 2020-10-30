package tests

import (
	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestSystemAPI_GetStatus() {

	status := &models.SystemStatus{}

	s.Fetch("/", &status)

	s.Assert().EqualValues(200, status.Code, "api is healthy")

}
