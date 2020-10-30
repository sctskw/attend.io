package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sctskw/attend.io/restapi/operations"
	"github.com/sctskw/attend.io/restapi/operations/events"
	"github.com/sctskw/attend.io/services"
)

func AttachEventsHandlers(api *operations.AttendIoAPI) {
	api.EventsGetEventsHandler = events.GetEventsHandlerFunc(GetEvents)
	api.EventsGetEventByIDHandler = events.GetEventByIDHandlerFunc(GetEventById)
	api.EventsGetEventAttendeesHandler = events.GetEventAttendeesHandlerFunc(GetEventAttendees)
}

func GetEvents(params events.GetEventsParams) middleware.Responder {
	return events.NewGetEventsOK().WithPayload(services.Events().GetAll())
}

func GetEventById(params events.GetEventByIDParams) middleware.Responder {
	return events.NewGetEventByIDOK().WithPayload(services.Events().GetById(params.ID))
}

func GetEventAttendees(params events.GetEventAttendeesParams) middleware.Responder {
	return events.NewGetEventAttendeesOK().WithPayload(services.Events().GetAttendees(params.ID))
}
