// Code generated by go-swagger; DO NOT EDIT.

package attendees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sctskw/attend.io/models"
)

// GetAttendeeByFieldOKCode is the HTTP code returned for type GetAttendeeByFieldOK
const GetAttendeeByFieldOKCode int = 200

/*GetAttendeeByFieldOK find an attendee by id

swagger:response getAttendeeByFieldOK
*/
type GetAttendeeByFieldOK struct {

	/*
	  In: Body
	*/
	Payload *models.Attendee `json:"body,omitempty"`
}

// NewGetAttendeeByFieldOK creates GetAttendeeByFieldOK with default headers values
func NewGetAttendeeByFieldOK() *GetAttendeeByFieldOK {

	return &GetAttendeeByFieldOK{}
}

// WithPayload adds the payload to the get attendee by field o k response
func (o *GetAttendeeByFieldOK) WithPayload(payload *models.Attendee) *GetAttendeeByFieldOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get attendee by field o k response
func (o *GetAttendeeByFieldOK) SetPayload(payload *models.Attendee) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAttendeeByFieldOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAttendeeByFieldDefault generic error response

swagger:response getAttendeeByFieldDefault
*/
type GetAttendeeByFieldDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAttendeeByFieldDefault creates GetAttendeeByFieldDefault with default headers values
func NewGetAttendeeByFieldDefault(code int) *GetAttendeeByFieldDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAttendeeByFieldDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get attendee by field default response
func (o *GetAttendeeByFieldDefault) WithStatusCode(code int) *GetAttendeeByFieldDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get attendee by field default response
func (o *GetAttendeeByFieldDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get attendee by field default response
func (o *GetAttendeeByFieldDefault) WithPayload(payload *models.Error) *GetAttendeeByFieldDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get attendee by field default response
func (o *GetAttendeeByFieldDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAttendeeByFieldDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
