// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// CreateAppsV1NamespacedStatefulSetOKCode is the HTTP code returned for type CreateAppsV1NamespacedStatefulSetOK
const CreateAppsV1NamespacedStatefulSetOKCode int = 200

/*CreateAppsV1NamespacedStatefulSetOK OK

swagger:response createAppsV1NamespacedStatefulSetOK
*/
type CreateAppsV1NamespacedStatefulSetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1StatefulSet `json:"body,omitempty"`
}

// NewCreateAppsV1NamespacedStatefulSetOK creates CreateAppsV1NamespacedStatefulSetOK with default headers values
func NewCreateAppsV1NamespacedStatefulSetOK() *CreateAppsV1NamespacedStatefulSetOK {

	return &CreateAppsV1NamespacedStatefulSetOK{}
}

// WithPayload adds the payload to the create apps v1 namespaced stateful set o k response
func (o *CreateAppsV1NamespacedStatefulSetOK) WithPayload(payload *models.IoK8sAPIAppsV1StatefulSet) *CreateAppsV1NamespacedStatefulSetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apps v1 namespaced stateful set o k response
func (o *CreateAppsV1NamespacedStatefulSetOK) SetPayload(payload *models.IoK8sAPIAppsV1StatefulSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedStatefulSetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAppsV1NamespacedStatefulSetCreatedCode is the HTTP code returned for type CreateAppsV1NamespacedStatefulSetCreated
const CreateAppsV1NamespacedStatefulSetCreatedCode int = 201

/*CreateAppsV1NamespacedStatefulSetCreated Created

swagger:response createAppsV1NamespacedStatefulSetCreated
*/
type CreateAppsV1NamespacedStatefulSetCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1StatefulSet `json:"body,omitempty"`
}

// NewCreateAppsV1NamespacedStatefulSetCreated creates CreateAppsV1NamespacedStatefulSetCreated with default headers values
func NewCreateAppsV1NamespacedStatefulSetCreated() *CreateAppsV1NamespacedStatefulSetCreated {

	return &CreateAppsV1NamespacedStatefulSetCreated{}
}

// WithPayload adds the payload to the create apps v1 namespaced stateful set created response
func (o *CreateAppsV1NamespacedStatefulSetCreated) WithPayload(payload *models.IoK8sAPIAppsV1StatefulSet) *CreateAppsV1NamespacedStatefulSetCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apps v1 namespaced stateful set created response
func (o *CreateAppsV1NamespacedStatefulSetCreated) SetPayload(payload *models.IoK8sAPIAppsV1StatefulSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedStatefulSetCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAppsV1NamespacedStatefulSetAcceptedCode is the HTTP code returned for type CreateAppsV1NamespacedStatefulSetAccepted
const CreateAppsV1NamespacedStatefulSetAcceptedCode int = 202

/*CreateAppsV1NamespacedStatefulSetAccepted Accepted

swagger:response createAppsV1NamespacedStatefulSetAccepted
*/
type CreateAppsV1NamespacedStatefulSetAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1StatefulSet `json:"body,omitempty"`
}

// NewCreateAppsV1NamespacedStatefulSetAccepted creates CreateAppsV1NamespacedStatefulSetAccepted with default headers values
func NewCreateAppsV1NamespacedStatefulSetAccepted() *CreateAppsV1NamespacedStatefulSetAccepted {

	return &CreateAppsV1NamespacedStatefulSetAccepted{}
}

// WithPayload adds the payload to the create apps v1 namespaced stateful set accepted response
func (o *CreateAppsV1NamespacedStatefulSetAccepted) WithPayload(payload *models.IoK8sAPIAppsV1StatefulSet) *CreateAppsV1NamespacedStatefulSetAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apps v1 namespaced stateful set accepted response
func (o *CreateAppsV1NamespacedStatefulSetAccepted) SetPayload(payload *models.IoK8sAPIAppsV1StatefulSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedStatefulSetAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAppsV1NamespacedStatefulSetUnauthorizedCode is the HTTP code returned for type CreateAppsV1NamespacedStatefulSetUnauthorized
const CreateAppsV1NamespacedStatefulSetUnauthorizedCode int = 401

/*CreateAppsV1NamespacedStatefulSetUnauthorized Unauthorized

swagger:response createAppsV1NamespacedStatefulSetUnauthorized
*/
type CreateAppsV1NamespacedStatefulSetUnauthorized struct {
}

// NewCreateAppsV1NamespacedStatefulSetUnauthorized creates CreateAppsV1NamespacedStatefulSetUnauthorized with default headers values
func NewCreateAppsV1NamespacedStatefulSetUnauthorized() *CreateAppsV1NamespacedStatefulSetUnauthorized {

	return &CreateAppsV1NamespacedStatefulSetUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedStatefulSetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
