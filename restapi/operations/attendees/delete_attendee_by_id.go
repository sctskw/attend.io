// Code generated by go-swagger; DO NOT EDIT.

package attendees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteAttendeeByIDHandlerFunc turns a function with the right signature into a delete attendee by Id handler
type DeleteAttendeeByIDHandlerFunc func(DeleteAttendeeByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAttendeeByIDHandlerFunc) Handle(params DeleteAttendeeByIDParams) middleware.Responder {
	return fn(params)
}

// DeleteAttendeeByIDHandler interface for that can handle valid delete attendee by Id params
type DeleteAttendeeByIDHandler interface {
	Handle(DeleteAttendeeByIDParams) middleware.Responder
}

// NewDeleteAttendeeByID creates a new http.Handler for the delete attendee by Id operation
func NewDeleteAttendeeByID(ctx *middleware.Context, handler DeleteAttendeeByIDHandler) *DeleteAttendeeByID {
	return &DeleteAttendeeByID{Context: ctx, Handler: handler}
}

/*DeleteAttendeeByID swagger:route DELETE /attendees/{id} attendees deleteAttendeeById

Deletes an Attendee by ID

*/
type DeleteAttendeeByID struct {
	Context *middleware.Context
	Handler DeleteAttendeeByIDHandler
}

func (o *DeleteAttendeeByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAttendeeByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
