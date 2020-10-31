package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sctskw/attend.io/restapi/operations"
	"github.com/sctskw/attend.io/restapi/operations/system"
	"github.com/sctskw/attend.io/services"
)

func AttachSystemHandlers(api *operations.AttendIoAPI) {
	api.SystemGetStatusHandler = system.GetStatusHandlerFunc(GetSystemStatus)
}

func GetSystemStatus(params system.GetStatusParams) middleware.Responder {
	return system.NewGetStatusOK().WithPayload(services.System().GetStatus())
}
