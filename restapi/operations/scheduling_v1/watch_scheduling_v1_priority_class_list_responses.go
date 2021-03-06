// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchSchedulingV1PriorityClassListOKCode is the HTTP code returned for type WatchSchedulingV1PriorityClassListOK
const WatchSchedulingV1PriorityClassListOKCode int = 200

/*WatchSchedulingV1PriorityClassListOK OK

swagger:response watchSchedulingV1PriorityClassListOK
*/
type WatchSchedulingV1PriorityClassListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchSchedulingV1PriorityClassListOK creates WatchSchedulingV1PriorityClassListOK with default headers values
func NewWatchSchedulingV1PriorityClassListOK() *WatchSchedulingV1PriorityClassListOK {

	return &WatchSchedulingV1PriorityClassListOK{}
}

// WithPayload adds the payload to the watch scheduling v1 priority class list o k response
func (o *WatchSchedulingV1PriorityClassListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchSchedulingV1PriorityClassListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch scheduling v1 priority class list o k response
func (o *WatchSchedulingV1PriorityClassListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchSchedulingV1PriorityClassListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchSchedulingV1PriorityClassListUnauthorizedCode is the HTTP code returned for type WatchSchedulingV1PriorityClassListUnauthorized
const WatchSchedulingV1PriorityClassListUnauthorizedCode int = 401

/*WatchSchedulingV1PriorityClassListUnauthorized Unauthorized

swagger:response watchSchedulingV1PriorityClassListUnauthorized
*/
type WatchSchedulingV1PriorityClassListUnauthorized struct {
}

// NewWatchSchedulingV1PriorityClassListUnauthorized creates WatchSchedulingV1PriorityClassListUnauthorized with default headers values
func NewWatchSchedulingV1PriorityClassListUnauthorized() *WatchSchedulingV1PriorityClassListUnauthorized {

	return &WatchSchedulingV1PriorityClassListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchSchedulingV1PriorityClassListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
