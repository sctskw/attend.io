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
}

func GetTalks(params talks.GetTalksParams) middleware.Responder {
	return talks.NewGetTalksOK().WithPayload(services.Talks().GetAll())
}

func GetTalkById(params talks.GetTalkByIDParams) middleware.Responder {
	return talks.NewGetTalkByIDOK().WithPayload(services.Talks().GetById(params.ID))
}

func GetTalkAttendees(params talks.GetTalkAttendeesParams) middleware.Responder {
	return talks.NewGetTalkAttendeesOK().WithPayload(services.Talks().GetAttendees(params.ID))
}

func CreateTalk(params talks.PostTalksParams) middleware.Responder {
	return talks.NewPostTalksOK().WithPayload(services.Talks().Create(params.Talk))
}
