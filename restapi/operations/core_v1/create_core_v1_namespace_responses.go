// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// CreateCoreV1NamespaceOKCode is the HTTP code returned for type CreateCoreV1NamespaceOK
const CreateCoreV1NamespaceOKCode int = 200

/*CreateCoreV1NamespaceOK OK

swagger:response createCoreV1NamespaceOK
*/
type CreateCoreV1NamespaceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Namespace `json:"body,omitempty"`
}

// NewCreateCoreV1NamespaceOK creates CreateCoreV1NamespaceOK with default headers values
func NewCreateCoreV1NamespaceOK() *CreateCoreV1NamespaceOK {

	return &CreateCoreV1NamespaceOK{}
}

// WithPayload adds the payload to the create core v1 namespace o k response
func (o *CreateCoreV1NamespaceOK) WithPayload(payload *models.IoK8sAPICoreV1Namespace) *CreateCoreV1NamespaceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespace o k response
func (o *CreateCoreV1NamespaceOK) SetPayload(payload *models.IoK8sAPICoreV1Namespace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespaceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespaceCreatedCode is the HTTP code returned for type CreateCoreV1NamespaceCreated
const CreateCoreV1NamespaceCreatedCode int = 201

/*CreateCoreV1NamespaceCreated Created

swagger:response createCoreV1NamespaceCreated
*/
type CreateCoreV1NamespaceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Namespace `json:"body,omitempty"`
}

// NewCreateCoreV1NamespaceCreated creates CreateCoreV1NamespaceCreated with default headers values
func NewCreateCoreV1NamespaceCreated() *CreateCoreV1NamespaceCreated {

	return &CreateCoreV1NamespaceCreated{}
}

// WithPayload adds the payload to the create core v1 namespace created response
func (o *CreateCoreV1NamespaceCreated) WithPayload(payload *models.IoK8sAPICoreV1Namespace) *CreateCoreV1NamespaceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespace created response
func (o *CreateCoreV1NamespaceCreated) SetPayload(payload *models.IoK8sAPICoreV1Namespace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespaceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespaceAcceptedCode is the HTTP code returned for type CreateCoreV1NamespaceAccepted
const CreateCoreV1NamespaceAcceptedCode int = 202

/*CreateCoreV1NamespaceAccepted Accepted

swagger:response createCoreV1NamespaceAccepted
*/
type CreateCoreV1NamespaceAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Namespace `json:"body,omitempty"`
}

// NewCreateCoreV1NamespaceAccepted creates CreateCoreV1NamespaceAccepted with default headers values
func NewCreateCoreV1NamespaceAccepted() *CreateCoreV1NamespaceAccepted {

	return &CreateCoreV1NamespaceAccepted{}
}

// WithPayload adds the payload to the create core v1 namespace accepted response
func (o *CreateCoreV1NamespaceAccepted) WithPayload(payload *models.IoK8sAPICoreV1Namespace) *CreateCoreV1NamespaceAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespace accepted response
func (o *CreateCoreV1NamespaceAccepted) SetPayload(payload *models.IoK8sAPICoreV1Namespace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespaceAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespaceUnauthorizedCode is the HTTP code returned for type CreateCoreV1NamespaceUnauthorized
const CreateCoreV1NamespaceUnauthorizedCode int = 401

/*CreateCoreV1NamespaceUnauthorized Unauthorized

swagger:response createCoreV1NamespaceUnauthorized
*/
type CreateCoreV1NamespaceUnauthorized struct {
}

// NewCreateCoreV1NamespaceUnauthorized creates CreateCoreV1NamespaceUnauthorized with default headers values
func NewCreateCoreV1NamespaceUnauthorized() *CreateCoreV1NamespaceUnauthorized {

	return &CreateCoreV1NamespaceUnauthorized{}
}

// WriteResponse to the client
func (o *CreateCoreV1NamespaceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
