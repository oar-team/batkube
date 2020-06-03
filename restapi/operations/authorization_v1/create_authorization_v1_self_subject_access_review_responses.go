// Code generated by go-swagger; DO NOT EDIT.

package authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// CreateAuthorizationV1SelfSubjectAccessReviewOKCode is the HTTP code returned for type CreateAuthorizationV1SelfSubjectAccessReviewOK
const CreateAuthorizationV1SelfSubjectAccessReviewOKCode int = 200

/*CreateAuthorizationV1SelfSubjectAccessReviewOK OK

swagger:response createAuthorizationV1SelfSubjectAccessReviewOK
*/
type CreateAuthorizationV1SelfSubjectAccessReviewOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAuthorizationV1SelfSubjectAccessReview `json:"body,omitempty"`
}

// NewCreateAuthorizationV1SelfSubjectAccessReviewOK creates CreateAuthorizationV1SelfSubjectAccessReviewOK with default headers values
func NewCreateAuthorizationV1SelfSubjectAccessReviewOK() *CreateAuthorizationV1SelfSubjectAccessReviewOK {

	return &CreateAuthorizationV1SelfSubjectAccessReviewOK{}
}

// WithPayload adds the payload to the create authorization v1 self subject access review o k response
func (o *CreateAuthorizationV1SelfSubjectAccessReviewOK) WithPayload(payload *models.IoK8sAPIAuthorizationV1SelfSubjectAccessReview) *CreateAuthorizationV1SelfSubjectAccessReviewOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authorization v1 self subject access review o k response
func (o *CreateAuthorizationV1SelfSubjectAccessReviewOK) SetPayload(payload *models.IoK8sAPIAuthorizationV1SelfSubjectAccessReview) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthorizationV1SelfSubjectAccessReviewOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthorizationV1SelfSubjectAccessReviewCreatedCode is the HTTP code returned for type CreateAuthorizationV1SelfSubjectAccessReviewCreated
const CreateAuthorizationV1SelfSubjectAccessReviewCreatedCode int = 201

/*CreateAuthorizationV1SelfSubjectAccessReviewCreated Created

swagger:response createAuthorizationV1SelfSubjectAccessReviewCreated
*/
type CreateAuthorizationV1SelfSubjectAccessReviewCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAuthorizationV1SelfSubjectAccessReview `json:"body,omitempty"`
}

// NewCreateAuthorizationV1SelfSubjectAccessReviewCreated creates CreateAuthorizationV1SelfSubjectAccessReviewCreated with default headers values
func NewCreateAuthorizationV1SelfSubjectAccessReviewCreated() *CreateAuthorizationV1SelfSubjectAccessReviewCreated {

	return &CreateAuthorizationV1SelfSubjectAccessReviewCreated{}
}

// WithPayload adds the payload to the create authorization v1 self subject access review created response
func (o *CreateAuthorizationV1SelfSubjectAccessReviewCreated) WithPayload(payload *models.IoK8sAPIAuthorizationV1SelfSubjectAccessReview) *CreateAuthorizationV1SelfSubjectAccessReviewCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authorization v1 self subject access review created response
func (o *CreateAuthorizationV1SelfSubjectAccessReviewCreated) SetPayload(payload *models.IoK8sAPIAuthorizationV1SelfSubjectAccessReview) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthorizationV1SelfSubjectAccessReviewCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthorizationV1SelfSubjectAccessReviewAcceptedCode is the HTTP code returned for type CreateAuthorizationV1SelfSubjectAccessReviewAccepted
const CreateAuthorizationV1SelfSubjectAccessReviewAcceptedCode int = 202

/*CreateAuthorizationV1SelfSubjectAccessReviewAccepted Accepted

swagger:response createAuthorizationV1SelfSubjectAccessReviewAccepted
*/
type CreateAuthorizationV1SelfSubjectAccessReviewAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAuthorizationV1SelfSubjectAccessReview `json:"body,omitempty"`
}

// NewCreateAuthorizationV1SelfSubjectAccessReviewAccepted creates CreateAuthorizationV1SelfSubjectAccessReviewAccepted with default headers values
func NewCreateAuthorizationV1SelfSubjectAccessReviewAccepted() *CreateAuthorizationV1SelfSubjectAccessReviewAccepted {

	return &CreateAuthorizationV1SelfSubjectAccessReviewAccepted{}
}

// WithPayload adds the payload to the create authorization v1 self subject access review accepted response
func (o *CreateAuthorizationV1SelfSubjectAccessReviewAccepted) WithPayload(payload *models.IoK8sAPIAuthorizationV1SelfSubjectAccessReview) *CreateAuthorizationV1SelfSubjectAccessReviewAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authorization v1 self subject access review accepted response
func (o *CreateAuthorizationV1SelfSubjectAccessReviewAccepted) SetPayload(payload *models.IoK8sAPIAuthorizationV1SelfSubjectAccessReview) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthorizationV1SelfSubjectAccessReviewAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthorizationV1SelfSubjectAccessReviewUnauthorizedCode is the HTTP code returned for type CreateAuthorizationV1SelfSubjectAccessReviewUnauthorized
const CreateAuthorizationV1SelfSubjectAccessReviewUnauthorizedCode int = 401

/*CreateAuthorizationV1SelfSubjectAccessReviewUnauthorized Unauthorized

swagger:response createAuthorizationV1SelfSubjectAccessReviewUnauthorized
*/
type CreateAuthorizationV1SelfSubjectAccessReviewUnauthorized struct {
}

// NewCreateAuthorizationV1SelfSubjectAccessReviewUnauthorized creates CreateAuthorizationV1SelfSubjectAccessReviewUnauthorized with default headers values
func NewCreateAuthorizationV1SelfSubjectAccessReviewUnauthorized() *CreateAuthorizationV1SelfSubjectAccessReviewUnauthorized {

	return &CreateAuthorizationV1SelfSubjectAccessReviewUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAuthorizationV1SelfSubjectAccessReviewUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}