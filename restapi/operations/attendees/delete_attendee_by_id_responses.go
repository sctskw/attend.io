// Code generated by go-swagger; DO NOT EDIT.

package attendees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sctskw/attend.io/models"
)

// DeleteAttendeeByIDNoContentCode is the HTTP code returned for type DeleteAttendeeByIDNoContent
const DeleteAttendeeByIDNoContentCode int = 204

/*DeleteAttendeeByIDNoContent Attendee was deleted.

swagger:response deleteAttendeeByIdNoContent
*/
type DeleteAttendeeByIDNoContent struct {
}

// NewDeleteAttendeeByIDNoContent creates DeleteAttendeeByIDNoContent with default headers values
func NewDeleteAttendeeByIDNoContent() *DeleteAttendeeByIDNoContent {

	return &DeleteAttendeeByIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteAttendeeByIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteAttendeeByIDDefault generic error response

swagger:response deleteAttendeeByIdDefault
*/
type DeleteAttendeeByIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteAttendeeByIDDefault creates DeleteAttendeeByIDDefault with default headers values
func NewDeleteAttendeeByIDDefault(code int) *DeleteAttendeeByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteAttendeeByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete attendee by Id default response
func (o *DeleteAttendeeByIDDefault) WithStatusCode(code int) *DeleteAttendeeByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete attendee by Id default response
func (o *DeleteAttendeeByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete attendee by Id default response
func (o *DeleteAttendeeByIDDefault) WithPayload(payload *models.Error) *DeleteAttendeeByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete attendee by Id default response
func (o *DeleteAttendeeByIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAttendeeByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}