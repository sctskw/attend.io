package tests

import (
	"github.com/sctskw/attend.io/models"
)

func (s *ApiTestSuite) TestEventsAPI_GetAllEvents() {

	events := models.NewEventList()

	s.Fetch("/events", &events)
	s.Assert().Len(events, 2, "returns total number of events")
}

func (s *ApiTestSuite) TestEventsAPI_GetEventById() {

	event := models.Event{}

	s.Fetch("/events/4dd1fae8-93ed-4a0d-a7d4-518e42488633", &event)
	s.Assert().NotNil(event, "returns an event")
	s.Assert().Equal("Event 3", *event.Name)

}
