// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// CreateCoreV1NamespacedConfigMapOKCode is the HTTP code returned for type CreateCoreV1NamespacedConfigMapOK
const CreateCoreV1NamespacedConfigMapOKCode int = 200

/*CreateCoreV1NamespacedConfigMapOK OK

swagger:response createCoreV1NamespacedConfigMapOK
*/
type CreateCoreV1NamespacedConfigMapOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ConfigMap `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedConfigMapOK creates CreateCoreV1NamespacedConfigMapOK with default headers values
func NewCreateCoreV1NamespacedConfigMapOK() *CreateCoreV1NamespacedConfigMapOK {

	return &CreateCoreV1NamespacedConfigMapOK{}
}

// WithPayload adds the payload to the create core v1 namespaced config map o k response
func (o *CreateCoreV1NamespacedConfigMapOK) WithPayload(payload *models.IoK8sAPICoreV1ConfigMap) *CreateCoreV1NamespacedConfigMapOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced config map o k response
func (o *CreateCoreV1NamespacedConfigMapOK) SetPayload(payload *models.IoK8sAPICoreV1ConfigMap) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedConfigMapOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedConfigMapCreatedCode is the HTTP code returned for type CreateCoreV1NamespacedConfigMapCreated
const CreateCoreV1NamespacedConfigMapCreatedCode int = 201

/*CreateCoreV1NamespacedConfigMapCreated Created

swagger:response createCoreV1NamespacedConfigMapCreated
*/
type CreateCoreV1NamespacedConfigMapCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ConfigMap `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedConfigMapCreated creates CreateCoreV1NamespacedConfigMapCreated with default headers values
func NewCreateCoreV1NamespacedConfigMapCreated() *CreateCoreV1NamespacedConfigMapCreated {

	return &CreateCoreV1NamespacedConfigMapCreated{}
}

// WithPayload adds the payload to the create core v1 namespaced config map created response
func (o *CreateCoreV1NamespacedConfigMapCreated) WithPayload(payload *models.IoK8sAPICoreV1ConfigMap) *CreateCoreV1NamespacedConfigMapCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced config map created response
func (o *CreateCoreV1NamespacedConfigMapCreated) SetPayload(payload *models.IoK8sAPICoreV1ConfigMap) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedConfigMapCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedConfigMapAcceptedCode is the HTTP code returned for type CreateCoreV1NamespacedConfigMapAccepted
const CreateCoreV1NamespacedConfigMapAcceptedCode int = 202

/*CreateCoreV1NamespacedConfigMapAccepted Accepted

swagger:response createCoreV1NamespacedConfigMapAccepted
*/
type CreateCoreV1NamespacedConfigMapAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ConfigMap `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedConfigMapAccepted creates CreateCoreV1NamespacedConfigMapAccepted with default headers values
func NewCreateCoreV1NamespacedConfigMapAccepted() *CreateCoreV1NamespacedConfigMapAccepted {

	return &CreateCoreV1NamespacedConfigMapAccepted{}
}

// WithPayload adds the payload to the create core v1 namespaced config map accepted response
func (o *CreateCoreV1NamespacedConfigMapAccepted) WithPayload(payload *models.IoK8sAPICoreV1ConfigMap) *CreateCoreV1NamespacedConfigMapAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced config map accepted response
func (o *CreateCoreV1NamespacedConfigMapAccepted) SetPayload(payload *models.IoK8sAPICoreV1ConfigMap) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedConfigMapAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedConfigMapUnauthorizedCode is the HTTP code returned for type CreateCoreV1NamespacedConfigMapUnauthorized
const CreateCoreV1NamespacedConfigMapUnauthorizedCode int = 401

/*CreateCoreV1NamespacedConfigMapUnauthorized Unauthorized

swagger:response createCoreV1NamespacedConfigMapUnauthorized
*/
type CreateCoreV1NamespacedConfigMapUnauthorized struct {
}

// NewCreateCoreV1NamespacedConfigMapUnauthorized creates CreateCoreV1NamespacedConfigMapUnauthorized with default headers values
func NewCreateCoreV1NamespacedConfigMapUnauthorized() *CreateCoreV1NamespacedConfigMapUnauthorized {

	return &CreateCoreV1NamespacedConfigMapUnauthorized{}
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedConfigMapUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
