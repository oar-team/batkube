// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceCoreV1NodeStatusOKCode is the HTTP code returned for type ReplaceCoreV1NodeStatusOK
const ReplaceCoreV1NodeStatusOKCode int = 200

/*ReplaceCoreV1NodeStatusOK OK

swagger:response replaceCoreV1NodeStatusOK
*/
type ReplaceCoreV1NodeStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Node `json:"body,omitempty"`
}

// NewReplaceCoreV1NodeStatusOK creates ReplaceCoreV1NodeStatusOK with default headers values
func NewReplaceCoreV1NodeStatusOK() *ReplaceCoreV1NodeStatusOK {

	return &ReplaceCoreV1NodeStatusOK{}
}

// WithPayload adds the payload to the replace core v1 node status o k response
func (o *ReplaceCoreV1NodeStatusOK) WithPayload(payload *models.IoK8sAPICoreV1Node) *ReplaceCoreV1NodeStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 node status o k response
func (o *ReplaceCoreV1NodeStatusOK) SetPayload(payload *models.IoK8sAPICoreV1Node) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NodeStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NodeStatusCreatedCode is the HTTP code returned for type ReplaceCoreV1NodeStatusCreated
const ReplaceCoreV1NodeStatusCreatedCode int = 201

/*ReplaceCoreV1NodeStatusCreated Created

swagger:response replaceCoreV1NodeStatusCreated
*/
type ReplaceCoreV1NodeStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Node `json:"body,omitempty"`
}

// NewReplaceCoreV1NodeStatusCreated creates ReplaceCoreV1NodeStatusCreated with default headers values
func NewReplaceCoreV1NodeStatusCreated() *ReplaceCoreV1NodeStatusCreated {

	return &ReplaceCoreV1NodeStatusCreated{}
}

// WithPayload adds the payload to the replace core v1 node status created response
func (o *ReplaceCoreV1NodeStatusCreated) WithPayload(payload *models.IoK8sAPICoreV1Node) *ReplaceCoreV1NodeStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 node status created response
func (o *ReplaceCoreV1NodeStatusCreated) SetPayload(payload *models.IoK8sAPICoreV1Node) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NodeStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NodeStatusUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NodeStatusUnauthorized
const ReplaceCoreV1NodeStatusUnauthorizedCode int = 401

/*ReplaceCoreV1NodeStatusUnauthorized Unauthorized

swagger:response replaceCoreV1NodeStatusUnauthorized
*/
type ReplaceCoreV1NodeStatusUnauthorized struct {
}

// NewReplaceCoreV1NodeStatusUnauthorized creates ReplaceCoreV1NodeStatusUnauthorized with default headers values
func NewReplaceCoreV1NodeStatusUnauthorized() *ReplaceCoreV1NodeStatusUnauthorized {

	return &ReplaceCoreV1NodeStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NodeStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
