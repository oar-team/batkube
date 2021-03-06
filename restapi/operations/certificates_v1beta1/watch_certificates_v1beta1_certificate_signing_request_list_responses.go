// Code generated by go-swagger; DO NOT EDIT.

package certificates_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchCertificatesV1beta1CertificateSigningRequestListOKCode is the HTTP code returned for type WatchCertificatesV1beta1CertificateSigningRequestListOK
const WatchCertificatesV1beta1CertificateSigningRequestListOKCode int = 200

/*WatchCertificatesV1beta1CertificateSigningRequestListOK OK

swagger:response watchCertificatesV1beta1CertificateSigningRequestListOK
*/
type WatchCertificatesV1beta1CertificateSigningRequestListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCertificatesV1beta1CertificateSigningRequestListOK creates WatchCertificatesV1beta1CertificateSigningRequestListOK with default headers values
func NewWatchCertificatesV1beta1CertificateSigningRequestListOK() *WatchCertificatesV1beta1CertificateSigningRequestListOK {

	return &WatchCertificatesV1beta1CertificateSigningRequestListOK{}
}

// WithPayload adds the payload to the watch certificates v1beta1 certificate signing request list o k response
func (o *WatchCertificatesV1beta1CertificateSigningRequestListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCertificatesV1beta1CertificateSigningRequestListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch certificates v1beta1 certificate signing request list o k response
func (o *WatchCertificatesV1beta1CertificateSigningRequestListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCertificatesV1beta1CertificateSigningRequestListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCertificatesV1beta1CertificateSigningRequestListUnauthorizedCode is the HTTP code returned for type WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized
const WatchCertificatesV1beta1CertificateSigningRequestListUnauthorizedCode int = 401

/*WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized Unauthorized

swagger:response watchCertificatesV1beta1CertificateSigningRequestListUnauthorized
*/
type WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized struct {
}

// NewWatchCertificatesV1beta1CertificateSigningRequestListUnauthorized creates WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized with default headers values
func NewWatchCertificatesV1beta1CertificateSigningRequestListUnauthorized() *WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized {

	return &WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
