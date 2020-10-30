// Code generated by go-swagger; DO NOT EDIT.

package attendees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAttendeeByFieldHandlerFunc turns a function with the right signature into a get attendee by field handler
type GetAttendeeByFieldHandlerFunc func(GetAttendeeByFieldParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAttendeeByFieldHandlerFunc) Handle(params GetAttendeeByFieldParams) middleware.Responder {
	return fn(params)
}

// GetAttendeeByFieldHandler interface for that can handle valid get attendee by field params
type GetAttendeeByFieldHandler interface {
	Handle(GetAttendeeByFieldParams) middleware.Responder
}

// NewGetAttendeeByField creates a new http.Handler for the get attendee by field operation
func NewGetAttendeeByField(ctx *middleware.Context, handler GetAttendeeByFieldHandler) *GetAttendeeByField {
	return &GetAttendeeByField{Context: ctx, Handler: handler}
}

/*GetAttendeeByField swagger:route GET /attendees attendees getAttendeeByField

GetAttendeeByField get attendee by field API

*/
type GetAttendeeByField struct {
	Context *middleware.Context
	Handler GetAttendeeByFieldHandler
}

func (o *GetAttendeeByField) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAttendeeByFieldParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
