package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sctskw/attend.io/restapi/operations"
	"github.com/sctskw/attend.io/restapi/operations/system"
	"github.com/sctskw/attend.io/services"
)

func AttachSystemHandlers(api *operations.AttendIoAPI) {
	api.SystemGetHandler = system.GetHandlerFunc(GetSystemStatus)
}

func GetSystemStatus(params system.GetParams) middleware.Responder {
	return system.NewGetOK().WithPayload(services.GetSystemStatus())
}
