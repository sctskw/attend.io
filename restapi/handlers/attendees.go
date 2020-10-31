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
	api.AttendeesDeleteAttendeeByIDHandler = attendees.DeleteAttendeeByIDHandlerFunc(DeleteAttendee)
}

func GetAttendeeByField(params attendees.GetAttendeeByFieldParams) middleware.Responder {

	if params.Email != nil {

		result, err := services.Attendees().GetByEmail(*params.Email)

		if err != nil {
			return attendees.NewGetAttendeeByFieldDefault(500)
		}

		return attendees.NewGetAttendeeByFieldOK().WithPayload(result)
	}

	if params.ID != nil {
		result, err := services.Attendees().GetById(*params.ID)

		if err != nil {
			return attendees.NewGetAttendeeByFieldDefault(500)
		}

		return attendees.NewGetAttendeeByFieldOK().WithPayload(result)
	}

	return attendees.NewGetAttendeeByFieldDefault(400) //BAD REQUEST
}

func CreateAttendee(params attendees.PostAttendeesParams) middleware.Responder {
	result, err := services.Attendees().Create(params.Attendee)

	if err != nil {
		return attendees.NewGetAttendeeByFieldDefault(500)
	}

	return attendees.NewPostAttendeesOK().WithPayload(result)
}

func DeleteAttendee(params attendees.DeleteAttendeeByIDParams) middleware.Responder {
	err := services.Attendees().DeleteById(params.ID)

	if err != nil {
		return attendees.NewDeleteAttendeeByIDDefault(500)
	}

	return attendees.NewDeleteAttendeeByIDNoContent()
}
