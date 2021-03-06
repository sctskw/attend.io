// Code generated by go-swagger; DO NOT EDIT.

package talks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sctskw/attend.io/models"
)

// GetTalksOKCode is the HTTP code returned for type GetTalksOK
const GetTalksOKCode int = 200

/*GetTalksOK list the talks

swagger:response getTalksOK
*/
type GetTalksOK struct {

	/*
	  In: Body
	*/
	Payload models.TalkList `json:"body,omitempty"`
}

// NewGetTalksOK creates GetTalksOK with default headers values
func NewGetTalksOK() *GetTalksOK {

	return &GetTalksOK{}
}

// WithPayload adds the payload to the get talks o k response
func (o *GetTalksOK) WithPayload(payload models.TalkList) *GetTalksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get talks o k response
func (o *GetTalksOK) SetPayload(payload models.TalkList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTalksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.TalkList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetTalksDefault generic error response

swagger:response getTalksDefault
*/
type GetTalksDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTalksDefault creates GetTalksDefault with default headers values
func NewGetTalksDefault(code int) *GetTalksDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTalksDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get talks default response
func (o *GetTalksDefault) WithStatusCode(code int) *GetTalksDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get talks default response
func (o *GetTalksDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get talks default response
func (o *GetTalksDefault) WithPayload(payload *models.Error) *GetTalksDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get talks default response
func (o *GetTalksDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTalksDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
