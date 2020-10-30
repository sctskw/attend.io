// Code generated by go-swagger; DO NOT EDIT.

package talks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sctskw/attend.io/models"
)

// GetTalkAttendeesOKCode is the HTTP code returned for type GetTalkAttendeesOK
const GetTalkAttendeesOKCode int = 200

/*GetTalkAttendeesOK the list of attendees from the event

swagger:response getTalkAttendeesOK
*/
type GetTalkAttendeesOK struct {

	/*
	  In: Body
	*/
	Payload models.AttendeesList `json:"body,omitempty"`
}

// NewGetTalkAttendeesOK creates GetTalkAttendeesOK with default headers values
func NewGetTalkAttendeesOK() *GetTalkAttendeesOK {

	return &GetTalkAttendeesOK{}
}

// WithPayload adds the payload to the get talk attendees o k response
func (o *GetTalkAttendeesOK) WithPayload(payload models.AttendeesList) *GetTalkAttendeesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get talk attendees o k response
func (o *GetTalkAttendeesOK) SetPayload(payload models.AttendeesList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTalkAttendeesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.AttendeesList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetTalkAttendeesDefault generic error response

swagger:response getTalkAttendeesDefault
*/
type GetTalkAttendeesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTalkAttendeesDefault creates GetTalkAttendeesDefault with default headers values
func NewGetTalkAttendeesDefault(code int) *GetTalkAttendeesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTalkAttendeesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get talk attendees default response
func (o *GetTalkAttendeesDefault) WithStatusCode(code int) *GetTalkAttendeesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get talk attendees default response
func (o *GetTalkAttendeesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get talk attendees default response
func (o *GetTalkAttendeesDefault) WithPayload(payload *models.Error) *GetTalkAttendeesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get talk attendees default response
func (o *GetTalkAttendeesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTalkAttendeesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}