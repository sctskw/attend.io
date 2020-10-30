package handlers

import (
	"github.com/sctskw/attend.io/restapi/operations"
	"github.com/sctskw/attend.io/services"
)

func AttachHandlers(api *operations.AttendIoAPI) {

	//init the service registry
	services.New()

	AttachSystemHandlers(api)
	AttachTalksHandlers(api)
	AttachAttendeesHandlers(api)
}
