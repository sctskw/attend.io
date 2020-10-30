// Code generated by go-swagger; DO NOT EDIT.

package attendees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAttendeesByIDHandlerFunc turns a function with the right signature into a get attendees by Id handler
type GetAttendeesByIDHandlerFunc func(GetAttendeesByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAttendeesByIDHandlerFunc) Handle(params GetAttendeesByIDParams) middleware.Responder {
	return fn(params)
}

// GetAttendeesByIDHandler interface for that can handle valid get attendees by Id params
type GetAttendeesByIDHandler interface {
	Handle(GetAttendeesByIDParams) middleware.Responder
}

// NewGetAttendeesByID creates a new http.Handler for the get attendees by Id operation
func NewGetAttendeesByID(ctx *middleware.Context, handler GetAttendeesByIDHandler) *GetAttendeesByID {
	return &GetAttendeesByID{Context: ctx, Handler: handler}
}

/*GetAttendeesByID swagger:route GET /attendee/{id} attendees getAttendeesById

GetAttendeesByID get attendees by Id API

*/
type GetAttendeesByID struct {
	Context *middleware.Context
	Handler GetAttendeesByIDHandler
}

func (o *GetAttendeesByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAttendeesByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
