package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sctskw/attend.io/restapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func GetHttpServer() *httptest.Server {
	handler, err := restapi.GetAPIHandler()
	if err != nil {
		panic(err)
	}
	return httptest.NewServer(handler)
}

type ApiTestSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ApiTestSuite))
}

func (s *ApiTestSuite) SetupSuite() {
	s.server = GetHttpServer()
}

func (s *ApiTestSuite) TearDownSuite() {
	s.server.Close()
}

func (s *ApiTestSuite) Fetch(path string, d interface{}) {

	res, err := http.Get(s.server.URL + path)

	if err != nil {
		assert.FailNow(s.T(), err.Error())
	}

	assert.Nil(s.T(), err, fmt.Sprintf("no error on %s", path))

	if res == nil || res.Body == nil {
		assert.FailNow(s.T(), "invalid response body")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	assert.Nil(s.T(), err, "no error parsing response")

	err = json.Unmarshal(body, &d)

	if err != nil {
		assert.FailNow(s.T(), "cannot parse response body")
	}

}
