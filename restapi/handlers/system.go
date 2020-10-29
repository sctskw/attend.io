package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/sctskw/attend.io/models"
	"github.com/sctskw/attend.io/restapi/operations"
	"github.com/sctskw/attend.io/restapi/operations/system"
)

func AttachSystemHandlers(api *operations.AttendIoAPI) {
	api.SystemGetHandler = system.GetHandlerFunc(GetSystemStatus)
}

func GetSystemStatus(params system.GetParams) middleware.Responder {
	return system.NewGetOK().WithPayload(&models.SystemStatus{
		Code:    200,
		Message: swag.String("attend.io is running"),
	})
}
