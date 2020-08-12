// Code generated by go-swagger; DO NOT EDIT.

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// CreatePolicyV1beta1PodSecurityPolicyOKCode is the HTTP code returned for type CreatePolicyV1beta1PodSecurityPolicyOK
const CreatePolicyV1beta1PodSecurityPolicyOKCode int = 200

/*CreatePolicyV1beta1PodSecurityPolicyOK OK

swagger:response createPolicyV1beta1PodSecurityPolicyOK
*/
type CreatePolicyV1beta1PodSecurityPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy `json:"body,omitempty"`
}

// NewCreatePolicyV1beta1PodSecurityPolicyOK creates CreatePolicyV1beta1PodSecurityPolicyOK with default headers values
func NewCreatePolicyV1beta1PodSecurityPolicyOK() *CreatePolicyV1beta1PodSecurityPolicyOK {

	return &CreatePolicyV1beta1PodSecurityPolicyOK{}
}

// WithPayload adds the payload to the create policy v1beta1 pod security policy o k response
func (o *CreatePolicyV1beta1PodSecurityPolicyOK) WithPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) *CreatePolicyV1beta1PodSecurityPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create policy v1beta1 pod security policy o k response
func (o *CreatePolicyV1beta1PodSecurityPolicyOK) SetPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePolicyV1beta1PodSecurityPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePolicyV1beta1PodSecurityPolicyCreatedCode is the HTTP code returned for type CreatePolicyV1beta1PodSecurityPolicyCreated
const CreatePolicyV1beta1PodSecurityPolicyCreatedCode int = 201

/*CreatePolicyV1beta1PodSecurityPolicyCreated Created

swagger:response createPolicyV1beta1PodSecurityPolicyCreated
*/
type CreatePolicyV1beta1PodSecurityPolicyCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy `json:"body,omitempty"`
}

// NewCreatePolicyV1beta1PodSecurityPolicyCreated creates CreatePolicyV1beta1PodSecurityPolicyCreated with default headers values
func NewCreatePolicyV1beta1PodSecurityPolicyCreated() *CreatePolicyV1beta1PodSecurityPolicyCreated {

	return &CreatePolicyV1beta1PodSecurityPolicyCreated{}
}

// WithPayload adds the payload to the create policy v1beta1 pod security policy created response
func (o *CreatePolicyV1beta1PodSecurityPolicyCreated) WithPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) *CreatePolicyV1beta1PodSecurityPolicyCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create policy v1beta1 pod security policy created response
func (o *CreatePolicyV1beta1PodSecurityPolicyCreated) SetPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePolicyV1beta1PodSecurityPolicyCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePolicyV1beta1PodSecurityPolicyAcceptedCode is the HTTP code returned for type CreatePolicyV1beta1PodSecurityPolicyAccepted
const CreatePolicyV1beta1PodSecurityPolicyAcceptedCode int = 202

/*CreatePolicyV1beta1PodSecurityPolicyAccepted Accepted

swagger:response createPolicyV1beta1PodSecurityPolicyAccepted
*/
type CreatePolicyV1beta1PodSecurityPolicyAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy `json:"body,omitempty"`
}

// NewCreatePolicyV1beta1PodSecurityPolicyAccepted creates CreatePolicyV1beta1PodSecurityPolicyAccepted with default headers values
func NewCreatePolicyV1beta1PodSecurityPolicyAccepted() *CreatePolicyV1beta1PodSecurityPolicyAccepted {

	return &CreatePolicyV1beta1PodSecurityPolicyAccepted{}
}

// WithPayload adds the payload to the create policy v1beta1 pod security policy accepted response
func (o *CreatePolicyV1beta1PodSecurityPolicyAccepted) WithPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) *CreatePolicyV1beta1PodSecurityPolicyAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create policy v1beta1 pod security policy accepted response
func (o *CreatePolicyV1beta1PodSecurityPolicyAccepted) SetPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePolicyV1beta1PodSecurityPolicyAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePolicyV1beta1PodSecurityPolicyUnauthorizedCode is the HTTP code returned for type CreatePolicyV1beta1PodSecurityPolicyUnauthorized
const CreatePolicyV1beta1PodSecurityPolicyUnauthorizedCode int = 401

/*CreatePolicyV1beta1PodSecurityPolicyUnauthorized Unauthorized

swagger:response createPolicyV1beta1PodSecurityPolicyUnauthorized
*/
type CreatePolicyV1beta1PodSecurityPolicyUnauthorized struct {
}

// NewCreatePolicyV1beta1PodSecurityPolicyUnauthorized creates CreatePolicyV1beta1PodSecurityPolicyUnauthorized with default headers values
func NewCreatePolicyV1beta1PodSecurityPolicyUnauthorized() *CreatePolicyV1beta1PodSecurityPolicyUnauthorized {

	return &CreatePolicyV1beta1PodSecurityPolicyUnauthorized{}
}

// WriteResponse to the client
func (o *CreatePolicyV1beta1PodSecurityPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
