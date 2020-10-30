package handlers

import (
	"github.com/sctskw/attend.io/restapi/operations"
)

func AttachHandlers(api *operations.AttendIoAPI) {
	AttachSystemHandlers(api)
	AttachTalksHandlers(api)
	AttachAttendeesHandlers(api)
	AttachEventsHandlers(api)
}
