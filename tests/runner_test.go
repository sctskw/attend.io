package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sctskw/attend.io/services"

	"github.com/sctskw/attend.io/db"
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
	server   *httptest.Server
	db       db.DatabaseClient
	services *services.ServiceRegistry
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ApiTestSuite))
}

func (s *ApiTestSuite) SetupSuite() {

	db.SetCredentialsPath("../.gcloud/attend-io.creds.json")
	s.server = GetHttpServer()
	s.db = db.NewClient()
	s.services = services.NewWithClient(s.db)
}

func (s *ApiTestSuite) TearDownSuite() {
	s.server.Close()
}

func (s *ApiTestSuite) Fetch(path string, d interface{}) (resp *http.Response) {

	res, err := http.Get(s.server.URL + path)

	if err != nil {
		assert.FailNow(s.T(), err.Error())
	}

	assert.Nil(s.T(), err, fmt.Sprintf("no error on %s", path))

	if res == nil || res.Body == nil {
		assert.FailNow(s.T(), "invalid response body")
	}

	if res.StatusCode != 200 {
		assert.FailNow(s.T(), fmt.Sprintf("response status %d: %+v", res.StatusCode, res.Body))
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	assert.Nil(s.T(), err, "no error parsing response")

	err = json.Unmarshal(body, &d)

	if err != nil {
		assert.FailNow(s.T(), "cannot parse response body")
	}

	return res

}

func (s *ApiTestSuite) Create(path string, d interface{}, r interface{}) (resp *http.Response) {

	b, _ := json.Marshal(d)
	res, err := http.Post(s.server.URL+path, "application/json", bytes.NewBuffer(b))

	if err != nil {
		assert.FailNow(s.T(), err.Error())
	}

	if res == nil || res.Body == nil {
		assert.FailNow(s.T(), "invalid response body")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		assert.FailNow(s.T(), string(body))
	}

	if err != nil {
		assert.FailNow(s.T(), err.Error())
	}

	err = json.Unmarshal(body, &r)

	if err != nil {
		assert.FailNow(s.T(), fmt.Sprintf("cannot parse response body: %s", err))
	}

	return res
}

func (s *ApiTestSuite) Update(path string, d interface{}, r interface{}) (resp *http.Response, err error) {

	b, _ := json.Marshal(d)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, s.server.URL+path, bytes.NewBuffer(b))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res == nil || res.Body == nil {
		assert.FailNow(s.T(), "invalid response body")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		assert.FailNow(s.T(), string(body))
	}

	if err != nil {
		assert.FailNow(s.T(), err.Error())
	}

	err = json.Unmarshal(body, &r)

	if err != nil {
		assert.FailNow(s.T(), fmt.Sprintf("cannot parse response body: %s", err))
	}

	return res, nil
}

func (s *ApiTestSuite) Delete(path string) (resp *http.Response) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, s.server.URL+path, nil)

	if err != nil {
		assert.FailNow(s.T(), err.Error())
	}

	res, err := client.Do(req)

	if res == nil || res.Body == nil {
		assert.FailNow(s.T(), "invalid response body")
	}

	if res.StatusCode != 200 && res.StatusCode != 204 {
		assert.FailNow(s.T(), fmt.Sprintf("response status: %d", res.StatusCode))
	}

	return res
}
