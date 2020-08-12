// Code generated by go-swagger; DO NOT EDIT.

package certificates_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceCertificatesV1beta1CertificateSigningRequestStatusOKCode is the HTTP code returned for type ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK
const ReplaceCertificatesV1beta1CertificateSigningRequestStatusOKCode int = 200

/*ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK OK

swagger:response replaceCertificatesV1beta1CertificateSigningRequestStatusOK
*/
type ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest `json:"body,omitempty"`
}

// NewReplaceCertificatesV1beta1CertificateSigningRequestStatusOK creates ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK with default headers values
func NewReplaceCertificatesV1beta1CertificateSigningRequestStatusOK() *ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK {

	return &ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK{}
}

// WithPayload adds the payload to the replace certificates v1beta1 certificate signing request status o k response
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK) WithPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest) *ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace certificates v1beta1 certificate signing request status o k response
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK) SetPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreatedCode is the HTTP code returned for type ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated
const ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreatedCode int = 201

/*ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated Created

swagger:response replaceCertificatesV1beta1CertificateSigningRequestStatusCreated
*/
type ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest `json:"body,omitempty"`
}

// NewReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated creates ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated with default headers values
func NewReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated() *ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated {

	return &ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated{}
}

// WithPayload adds the payload to the replace certificates v1beta1 certificate signing request status created response
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated) WithPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest) *ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace certificates v1beta1 certificate signing request status created response
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated) SetPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorizedCode is the HTTP code returned for type ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized
const ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorizedCode int = 401

/*ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized Unauthorized

swagger:response replaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized
*/
type ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized struct {
}

// NewReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized creates ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized with default headers values
func NewReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized() *ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized {

	return &ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
