package tests

import (
	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestEventsAPI_GetAllEvents() {

	events := models.NewEventList()

	s.Fetch("/events", &events)
	s.Assert().Len(events, 1, "returns total number of events")
}

func (s *ApiTestSuite) TestEventsAPI_GetEventById() {

	event := &models.Event{}

	s.Fetch("/events/tQ3Z0tRsz1oNo1lrNVBd", event)
	s.Assert().NotNil(event, "returns an event")
	s.Assert().Equal("Event 1", *event.Name)

}

func (s *ApiTestSuite) TestEventsAPI_GetEventAttendees() {

	attendees := models.NewAttendeeList()
	s.Fetch("/events/tQ3Z0tRsz1oNo1lrNVBd/attendees", &attendees)
	s.Assert().Len(attendees, 3, "returns total number of attendees")

	for _, a := range attendees {
		s.Assert().NotEmpty(a.Email)
		s.Assert().NotEmpty(a.NameFirst)
		s.Assert().NotEmpty(a.NameLast)
	}
}
