package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sctskw/attend.io/restapi/operations"
	"github.com/sctskw/attend.io/restapi/operations/attendees"
	"github.com/sctskw/attend.io/services"
)

func AttachAttendeesHandlers(api *operations.AttendIoAPI) {
	api.AttendeesGetAttendeesByEventIDHandler = attendees.GetAttendeesByEventIDHandlerFunc(GetAttendeesByEventId)
	api.AttendeesGetAttendeeByFieldHandler = attendees.GetAttendeeByFieldHandlerFunc(GetAttendeeByField)
}

func GetAttendeesByEventId(params attendees.GetAttendeesByEventIDParams) middleware.Responder {
	return attendees.NewGetAttendeesByEventIDOK().WithPayload(services.Get().Attendees().GetByEventId(params.EventID))
}

func GetAttendeeByField(params attendees.GetAttendeeByFieldParams) middleware.Responder {
	if params.Email != nil {
		return attendees.NewGetAttendeeByFieldOK().WithPayload(services.Get().Attendees().GetByEmail(*params.Email))
	}

	if params.ID != nil {
		return attendees.NewGetAttendeeByFieldOK().WithPayload(services.Get().Attendees().GetById(*params.ID))
	}

	return attendees.NewGetAttendeeByFieldDefault(404) //NOT FOUND
}
