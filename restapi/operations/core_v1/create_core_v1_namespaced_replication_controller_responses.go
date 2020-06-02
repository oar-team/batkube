// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// CreateCoreV1NamespacedReplicationControllerOKCode is the HTTP code returned for type CreateCoreV1NamespacedReplicationControllerOK
const CreateCoreV1NamespacedReplicationControllerOKCode int = 200

/*CreateCoreV1NamespacedReplicationControllerOK OK

swagger:response createCoreV1NamespacedReplicationControllerOK
*/
type CreateCoreV1NamespacedReplicationControllerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ReplicationController `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedReplicationControllerOK creates CreateCoreV1NamespacedReplicationControllerOK with default headers values
func NewCreateCoreV1NamespacedReplicationControllerOK() *CreateCoreV1NamespacedReplicationControllerOK {

	return &CreateCoreV1NamespacedReplicationControllerOK{}
}

// WithPayload adds the payload to the create core v1 namespaced replication controller o k response
func (o *CreateCoreV1NamespacedReplicationControllerOK) WithPayload(payload *models.IoK8sAPICoreV1ReplicationController) *CreateCoreV1NamespacedReplicationControllerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced replication controller o k response
func (o *CreateCoreV1NamespacedReplicationControllerOK) SetPayload(payload *models.IoK8sAPICoreV1ReplicationController) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedReplicationControllerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedReplicationControllerCreatedCode is the HTTP code returned for type CreateCoreV1NamespacedReplicationControllerCreated
const CreateCoreV1NamespacedReplicationControllerCreatedCode int = 201

/*CreateCoreV1NamespacedReplicationControllerCreated Created

swagger:response createCoreV1NamespacedReplicationControllerCreated
*/
type CreateCoreV1NamespacedReplicationControllerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ReplicationController `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedReplicationControllerCreated creates CreateCoreV1NamespacedReplicationControllerCreated with default headers values
func NewCreateCoreV1NamespacedReplicationControllerCreated() *CreateCoreV1NamespacedReplicationControllerCreated {

	return &CreateCoreV1NamespacedReplicationControllerCreated{}
}

// WithPayload adds the payload to the create core v1 namespaced replication controller created response
func (o *CreateCoreV1NamespacedReplicationControllerCreated) WithPayload(payload *models.IoK8sAPICoreV1ReplicationController) *CreateCoreV1NamespacedReplicationControllerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced replication controller created response
func (o *CreateCoreV1NamespacedReplicationControllerCreated) SetPayload(payload *models.IoK8sAPICoreV1ReplicationController) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedReplicationControllerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedReplicationControllerAcceptedCode is the HTTP code returned for type CreateCoreV1NamespacedReplicationControllerAccepted
const CreateCoreV1NamespacedReplicationControllerAcceptedCode int = 202

/*CreateCoreV1NamespacedReplicationControllerAccepted Accepted

swagger:response createCoreV1NamespacedReplicationControllerAccepted
*/
type CreateCoreV1NamespacedReplicationControllerAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ReplicationController `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedReplicationControllerAccepted creates CreateCoreV1NamespacedReplicationControllerAccepted with default headers values
func NewCreateCoreV1NamespacedReplicationControllerAccepted() *CreateCoreV1NamespacedReplicationControllerAccepted {

	return &CreateCoreV1NamespacedReplicationControllerAccepted{}
}

// WithPayload adds the payload to the create core v1 namespaced replication controller accepted response
func (o *CreateCoreV1NamespacedReplicationControllerAccepted) WithPayload(payload *models.IoK8sAPICoreV1ReplicationController) *CreateCoreV1NamespacedReplicationControllerAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced replication controller accepted response
func (o *CreateCoreV1NamespacedReplicationControllerAccepted) SetPayload(payload *models.IoK8sAPICoreV1ReplicationController) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedReplicationControllerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedReplicationControllerUnauthorizedCode is the HTTP code returned for type CreateCoreV1NamespacedReplicationControllerUnauthorized
const CreateCoreV1NamespacedReplicationControllerUnauthorizedCode int = 401

/*CreateCoreV1NamespacedReplicationControllerUnauthorized Unauthorized

swagger:response createCoreV1NamespacedReplicationControllerUnauthorized
*/
type CreateCoreV1NamespacedReplicationControllerUnauthorized struct {
}

// NewCreateCoreV1NamespacedReplicationControllerUnauthorized creates CreateCoreV1NamespacedReplicationControllerUnauthorized with default headers values
func NewCreateCoreV1NamespacedReplicationControllerUnauthorized() *CreateCoreV1NamespacedReplicationControllerUnauthorized {

	return &CreateCoreV1NamespacedReplicationControllerUnauthorized{}
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedReplicationControllerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
