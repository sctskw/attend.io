// Code generated by go-swagger; DO NOT EDIT.

package talks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostTalksHandlerFunc turns a function with the right signature into a post talks handler
type PostTalksHandlerFunc func(PostTalksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostTalksHandlerFunc) Handle(params PostTalksParams) middleware.Responder {
	return fn(params)
}

// PostTalksHandler interface for that can handle valid post talks params
type PostTalksHandler interface {
	Handle(PostTalksParams) middleware.Responder
}

// NewPostTalks creates a new http.Handler for the post talks operation
func NewPostTalks(ctx *middleware.Context, handler PostTalksHandler) *PostTalks {
	return &PostTalks{Context: ctx, Handler: handler}
}

/*PostTalks swagger:route POST /talks talks postTalks

create a talk

*/
type PostTalks struct {
	Context *middleware.Context
	Handler PostTalksHandler
}

func (o *PostTalks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostTalksParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
