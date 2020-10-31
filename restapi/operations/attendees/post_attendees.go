// Code generated by go-swagger; DO NOT EDIT.

package attendees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostAttendeesHandlerFunc turns a function with the right signature into a post attendees handler
type PostAttendeesHandlerFunc func(PostAttendeesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAttendeesHandlerFunc) Handle(params PostAttendeesParams) middleware.Responder {
	return fn(params)
}

// PostAttendeesHandler interface for that can handle valid post attendees params
type PostAttendeesHandler interface {
	Handle(PostAttendeesParams) middleware.Responder
}

// NewPostAttendees creates a new http.Handler for the post attendees operation
func NewPostAttendees(ctx *middleware.Context, handler PostAttendeesHandler) *PostAttendees {
	return &PostAttendees{Context: ctx, Handler: handler}
}

/*PostAttendees swagger:route POST /attendees attendees postAttendees

Create an Attendee

*/
type PostAttendees struct {
	Context *middleware.Context
	Handler PostAttendeesHandler
}

func (o *PostAttendees) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostAttendeesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
