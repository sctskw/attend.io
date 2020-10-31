package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sctskw/attend.io/restapi/operations"
	"github.com/sctskw/attend.io/restapi/operations/talks"
	"github.com/sctskw/attend.io/services"
)

func AttachTalksHandlers(api *operations.AttendIoAPI) {
	api.TalksGetTalksHandler = talks.GetTalksHandlerFunc(GetTalks)
	api.TalksGetTalkByIDHandler = talks.GetTalkByIDHandlerFunc(GetTalkById)
	api.TalksGetTalkAttendeesHandler = talks.GetTalkAttendeesHandlerFunc(GetTalkAttendees)
	api.TalksPostTalksHandler = talks.PostTalksHandlerFunc(CreateTalk)
	api.TalksAddAttendeeToTalkHandler = talks.AddAttendeeToTalkHandlerFunc(AddAttendeeToTalk)
}

func GetTalks(params talks.GetTalksParams) middleware.Responder {
	return talks.NewGetTalksOK().WithPayload(services.Talks().GetAll())
}

func GetTalkById(params talks.GetTalkByIDParams) middleware.Responder {
	result, err := services.Talks().GetById(params.ID)

	if err != nil {
		return talks.NewGetTalkByIDDefault(500)
	}

	return talks.NewGetTalkByIDOK().WithPayload(result)
}

func GetTalkAttendees(params talks.GetTalkAttendeesParams) middleware.Responder {
	result, err := services.Talks().GetAttendees(params.ID)

	if err != nil {
		return talks.NewGetTalkByIDDefault(500)
	}

	return talks.NewGetTalkAttendeesOK().WithPayload(result)
}

func CreateTalk(params talks.PostTalksParams) middleware.Responder {

	result, err := services.Talks().Create(params.Talk)

	if err != nil {
		return talks.NewGetTalkByIDDefault(500)
	}

	return talks.NewPostTalksOK().WithPayload(result)
}

func AddAttendeeToTalk(params talks.AddAttendeeToTalkParams) middleware.Responder {

	result, err := services.Talks().AddAttendee(params.ID, params.Attendees)

	if err != nil {
		return talks.NewGetTalkByIDDefault(500)
	}

	return talks.NewPostTalksOK().WithPayload(result)
}
