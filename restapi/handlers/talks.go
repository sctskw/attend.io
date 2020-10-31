package handlers

import (
	"fmt"

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
	api.TalksDeleteAttendeeFromTalkHandler = talks.DeleteAttendeeFromTalkHandlerFunc(RemoveAttendeeFromTalk)
	api.TalksDeleteTalkByIDHandler = talks.DeleteTalkByIDHandlerFunc(DeleteTalk)
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

func DeleteTalk(params talks.DeleteTalkByIDParams) middleware.Responder {

	err := services.Talks().Delete(params.ID)

	if err != nil {
		return talks.NewGetTalkByIDDefault(500)
	}

	return talks.NewPostTalksDefault(204)
}

func AddAttendeeToTalk(params talks.AddAttendeeToTalkParams) middleware.Responder {

	fmt.Println(fmt.Sprintf("Talk: %s -> Attendees: %+v", params.ID, params.Attendees))

	result, err := services.Talks().AddAttendee(params.ID, params.Attendees)

	if err != nil {
		return talks.NewGetTalkByIDDefault(500)
	}

	return talks.NewPostTalksOK().WithPayload(result)
}

func RemoveAttendeeFromTalk(params talks.DeleteAttendeeFromTalkParams) middleware.Responder {

	err := services.Talks().RemoveAttendee(params.ID, params.AttendeeID)

	if err != nil {
		return talks.NewGetTalkByIDDefault(500)
	}

	return talks.NewPostTalksDefault(204)
}
