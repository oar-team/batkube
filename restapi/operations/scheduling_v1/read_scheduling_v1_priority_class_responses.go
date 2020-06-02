// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadSchedulingV1PriorityClassOKCode is the HTTP code returned for type ReadSchedulingV1PriorityClassOK
const ReadSchedulingV1PriorityClassOKCode int = 200

/*ReadSchedulingV1PriorityClassOK OK

swagger:response readSchedulingV1PriorityClassOK
*/
type ReadSchedulingV1PriorityClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPISchedulingV1PriorityClass `json:"body,omitempty"`
}

// NewReadSchedulingV1PriorityClassOK creates ReadSchedulingV1PriorityClassOK with default headers values
func NewReadSchedulingV1PriorityClassOK() *ReadSchedulingV1PriorityClassOK {

	return &ReadSchedulingV1PriorityClassOK{}
}

// WithPayload adds the payload to the read scheduling v1 priority class o k response
func (o *ReadSchedulingV1PriorityClassOK) WithPayload(payload *models.IoK8sAPISchedulingV1PriorityClass) *ReadSchedulingV1PriorityClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read scheduling v1 priority class o k response
func (o *ReadSchedulingV1PriorityClassOK) SetPayload(payload *models.IoK8sAPISchedulingV1PriorityClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadSchedulingV1PriorityClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadSchedulingV1PriorityClassUnauthorizedCode is the HTTP code returned for type ReadSchedulingV1PriorityClassUnauthorized
const ReadSchedulingV1PriorityClassUnauthorizedCode int = 401

/*ReadSchedulingV1PriorityClassUnauthorized Unauthorized

swagger:response readSchedulingV1PriorityClassUnauthorized
*/
type ReadSchedulingV1PriorityClassUnauthorized struct {
}

// NewReadSchedulingV1PriorityClassUnauthorized creates ReadSchedulingV1PriorityClassUnauthorized with default headers values
func NewReadSchedulingV1PriorityClassUnauthorized() *ReadSchedulingV1PriorityClassUnauthorized {

	return &ReadSchedulingV1PriorityClassUnauthorized{}
}

// WriteResponse to the client
func (o *ReadSchedulingV1PriorityClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
