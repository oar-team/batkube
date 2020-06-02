// Code generated by go-swagger; DO NOT EDIT.

package certificates_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// CreateCertificatesV1beta1CertificateSigningRequestOKCode is the HTTP code returned for type CreateCertificatesV1beta1CertificateSigningRequestOK
const CreateCertificatesV1beta1CertificateSigningRequestOKCode int = 200

/*CreateCertificatesV1beta1CertificateSigningRequestOK OK

swagger:response createCertificatesV1beta1CertificateSigningRequestOK
*/
type CreateCertificatesV1beta1CertificateSigningRequestOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest `json:"body,omitempty"`
}

// NewCreateCertificatesV1beta1CertificateSigningRequestOK creates CreateCertificatesV1beta1CertificateSigningRequestOK with default headers values
func NewCreateCertificatesV1beta1CertificateSigningRequestOK() *CreateCertificatesV1beta1CertificateSigningRequestOK {

	return &CreateCertificatesV1beta1CertificateSigningRequestOK{}
}

// WithPayload adds the payload to the create certificates v1beta1 certificate signing request o k response
func (o *CreateCertificatesV1beta1CertificateSigningRequestOK) WithPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest) *CreateCertificatesV1beta1CertificateSigningRequestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create certificates v1beta1 certificate signing request o k response
func (o *CreateCertificatesV1beta1CertificateSigningRequestOK) SetPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCertificatesV1beta1CertificateSigningRequestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCertificatesV1beta1CertificateSigningRequestCreatedCode is the HTTP code returned for type CreateCertificatesV1beta1CertificateSigningRequestCreated
const CreateCertificatesV1beta1CertificateSigningRequestCreatedCode int = 201

/*CreateCertificatesV1beta1CertificateSigningRequestCreated Created

swagger:response createCertificatesV1beta1CertificateSigningRequestCreated
*/
type CreateCertificatesV1beta1CertificateSigningRequestCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest `json:"body,omitempty"`
}

// NewCreateCertificatesV1beta1CertificateSigningRequestCreated creates CreateCertificatesV1beta1CertificateSigningRequestCreated with default headers values
func NewCreateCertificatesV1beta1CertificateSigningRequestCreated() *CreateCertificatesV1beta1CertificateSigningRequestCreated {

	return &CreateCertificatesV1beta1CertificateSigningRequestCreated{}
}

// WithPayload adds the payload to the create certificates v1beta1 certificate signing request created response
func (o *CreateCertificatesV1beta1CertificateSigningRequestCreated) WithPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest) *CreateCertificatesV1beta1CertificateSigningRequestCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create certificates v1beta1 certificate signing request created response
func (o *CreateCertificatesV1beta1CertificateSigningRequestCreated) SetPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCertificatesV1beta1CertificateSigningRequestCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCertificatesV1beta1CertificateSigningRequestAcceptedCode is the HTTP code returned for type CreateCertificatesV1beta1CertificateSigningRequestAccepted
const CreateCertificatesV1beta1CertificateSigningRequestAcceptedCode int = 202

/*CreateCertificatesV1beta1CertificateSigningRequestAccepted Accepted

swagger:response createCertificatesV1beta1CertificateSigningRequestAccepted
*/
type CreateCertificatesV1beta1CertificateSigningRequestAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest `json:"body,omitempty"`
}

// NewCreateCertificatesV1beta1CertificateSigningRequestAccepted creates CreateCertificatesV1beta1CertificateSigningRequestAccepted with default headers values
func NewCreateCertificatesV1beta1CertificateSigningRequestAccepted() *CreateCertificatesV1beta1CertificateSigningRequestAccepted {

	return &CreateCertificatesV1beta1CertificateSigningRequestAccepted{}
}

// WithPayload adds the payload to the create certificates v1beta1 certificate signing request accepted response
func (o *CreateCertificatesV1beta1CertificateSigningRequestAccepted) WithPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest) *CreateCertificatesV1beta1CertificateSigningRequestAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create certificates v1beta1 certificate signing request accepted response
func (o *CreateCertificatesV1beta1CertificateSigningRequestAccepted) SetPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCertificatesV1beta1CertificateSigningRequestAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCertificatesV1beta1CertificateSigningRequestUnauthorizedCode is the HTTP code returned for type CreateCertificatesV1beta1CertificateSigningRequestUnauthorized
const CreateCertificatesV1beta1CertificateSigningRequestUnauthorizedCode int = 401

/*CreateCertificatesV1beta1CertificateSigningRequestUnauthorized Unauthorized

swagger:response createCertificatesV1beta1CertificateSigningRequestUnauthorized
*/
type CreateCertificatesV1beta1CertificateSigningRequestUnauthorized struct {
}

// NewCreateCertificatesV1beta1CertificateSigningRequestUnauthorized creates CreateCertificatesV1beta1CertificateSigningRequestUnauthorized with default headers values
func NewCreateCertificatesV1beta1CertificateSigningRequestUnauthorized() *CreateCertificatesV1beta1CertificateSigningRequestUnauthorized {

	return &CreateCertificatesV1beta1CertificateSigningRequestUnauthorized{}
}

// WriteResponse to the client
func (o *CreateCertificatesV1beta1CertificateSigningRequestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
