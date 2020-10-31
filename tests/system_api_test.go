package tests

import (
	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestSystemAPI_GetStatus() {

	status := &models.SystemStatus{}

	s.Fetch("/status", &status)

	s.Assert().EqualValues(200, status.Code, "attend.io is alive.")

}
