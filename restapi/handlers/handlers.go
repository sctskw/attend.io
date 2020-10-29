package handlers

import (
	"github.com/sctskw/attend.io/restapi/operations"
)

func AttachHandlers(api *operations.AttendIoAPI) {
	AttachTalksHandlers(api)
	AttachSystemHandlers(api)
}
