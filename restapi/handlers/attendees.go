package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sctskw/attend.io/restapi/operations"
	"github.com/sctskw/attend.io/restapi/operations/attendees"
	"github.com/sctskw/attend.io/services"
)

func AttachAttendeesHandlers(api *operations.AttendIoAPI) {
	api.AttendeesGetAttendeeByFieldHandler = attendees.GetAttendeeByFieldHandlerFunc(GetAttendeeByField)
	api.AttendeesPostAttendeesHandler = attendees.PostAttendeesHandlerFunc(CreateAttendee)
}

func GetAttendeeByField(params attendees.GetAttendeeByFieldParams) middleware.Responder {
	if params.Email != nil {
		return attendees.NewGetAttendeeByFieldOK().WithPayload(services.Attendees().GetByEmail(*params.Email))
	}

	if params.ID != nil {
		return attendees.NewGetAttendeeByFieldOK().WithPayload(services.Attendees().GetById(*params.ID))
	}

	return attendees.NewGetAttendeeByFieldDefault(404) //NOT FOUND
}

func CreateAttendee(params attendees.PostAttendeesParams) middleware.Responder {
	return attendees.NewPostAttendeesOK().WithPayload(services.Attendees().Create(params.Attendee))
}
