package restapi

import (
	"net/http"

	loads "github.com/go-openapi/loads"
	"github.com/sctskw/attend.io/restapi/operations"
)

func getAPI() (*operations.AttendIoAPI, error) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	api := operations.NewAttendIoAPI(swaggerSpec)
	return api, nil
}

func GetAPIHandler() (http.Handler, error) {
	api, err := getAPI()
	if err != nil {
		return nil, err
	}
	h := configureAPI(api)
	err = api.Validate()
	if err != nil {
		return nil, err
	}
	return h, nil
}
