// Code generated by go-swagger; DO NOT EDIT.

package authentication_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// CreateAuthenticationV1TokenReviewOKCode is the HTTP code returned for type CreateAuthenticationV1TokenReviewOK
const CreateAuthenticationV1TokenReviewOKCode int = 200

/*CreateAuthenticationV1TokenReviewOK OK

swagger:response createAuthenticationV1TokenReviewOK
*/
type CreateAuthenticationV1TokenReviewOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAuthenticationV1TokenReview `json:"body,omitempty"`
}

// NewCreateAuthenticationV1TokenReviewOK creates CreateAuthenticationV1TokenReviewOK with default headers values
func NewCreateAuthenticationV1TokenReviewOK() *CreateAuthenticationV1TokenReviewOK {

	return &CreateAuthenticationV1TokenReviewOK{}
}

// WithPayload adds the payload to the create authentication v1 token review o k response
func (o *CreateAuthenticationV1TokenReviewOK) WithPayload(payload *models.IoK8sAPIAuthenticationV1TokenReview) *CreateAuthenticationV1TokenReviewOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authentication v1 token review o k response
func (o *CreateAuthenticationV1TokenReviewOK) SetPayload(payload *models.IoK8sAPIAuthenticationV1TokenReview) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthenticationV1TokenReviewOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthenticationV1TokenReviewCreatedCode is the HTTP code returned for type CreateAuthenticationV1TokenReviewCreated
const CreateAuthenticationV1TokenReviewCreatedCode int = 201

/*CreateAuthenticationV1TokenReviewCreated Created

swagger:response createAuthenticationV1TokenReviewCreated
*/
type CreateAuthenticationV1TokenReviewCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAuthenticationV1TokenReview `json:"body,omitempty"`
}

// NewCreateAuthenticationV1TokenReviewCreated creates CreateAuthenticationV1TokenReviewCreated with default headers values
func NewCreateAuthenticationV1TokenReviewCreated() *CreateAuthenticationV1TokenReviewCreated {

	return &CreateAuthenticationV1TokenReviewCreated{}
}

// WithPayload adds the payload to the create authentication v1 token review created response
func (o *CreateAuthenticationV1TokenReviewCreated) WithPayload(payload *models.IoK8sAPIAuthenticationV1TokenReview) *CreateAuthenticationV1TokenReviewCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authentication v1 token review created response
func (o *CreateAuthenticationV1TokenReviewCreated) SetPayload(payload *models.IoK8sAPIAuthenticationV1TokenReview) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthenticationV1TokenReviewCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthenticationV1TokenReviewAcceptedCode is the HTTP code returned for type CreateAuthenticationV1TokenReviewAccepted
const CreateAuthenticationV1TokenReviewAcceptedCode int = 202

/*CreateAuthenticationV1TokenReviewAccepted Accepted

swagger:response createAuthenticationV1TokenReviewAccepted
*/
type CreateAuthenticationV1TokenReviewAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAuthenticationV1TokenReview `json:"body,omitempty"`
}

// NewCreateAuthenticationV1TokenReviewAccepted creates CreateAuthenticationV1TokenReviewAccepted with default headers values
func NewCreateAuthenticationV1TokenReviewAccepted() *CreateAuthenticationV1TokenReviewAccepted {

	return &CreateAuthenticationV1TokenReviewAccepted{}
}

// WithPayload adds the payload to the create authentication v1 token review accepted response
func (o *CreateAuthenticationV1TokenReviewAccepted) WithPayload(payload *models.IoK8sAPIAuthenticationV1TokenReview) *CreateAuthenticationV1TokenReviewAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authentication v1 token review accepted response
func (o *CreateAuthenticationV1TokenReviewAccepted) SetPayload(payload *models.IoK8sAPIAuthenticationV1TokenReview) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthenticationV1TokenReviewAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthenticationV1TokenReviewUnauthorizedCode is the HTTP code returned for type CreateAuthenticationV1TokenReviewUnauthorized
const CreateAuthenticationV1TokenReviewUnauthorizedCode int = 401

/*CreateAuthenticationV1TokenReviewUnauthorized Unauthorized

swagger:response createAuthenticationV1TokenReviewUnauthorized
*/
type CreateAuthenticationV1TokenReviewUnauthorized struct {
}

// NewCreateAuthenticationV1TokenReviewUnauthorized creates CreateAuthenticationV1TokenReviewUnauthorized with default headers values
func NewCreateAuthenticationV1TokenReviewUnauthorized() *CreateAuthenticationV1TokenReviewUnauthorized {

	return &CreateAuthenticationV1TokenReviewUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAuthenticationV1TokenReviewUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
