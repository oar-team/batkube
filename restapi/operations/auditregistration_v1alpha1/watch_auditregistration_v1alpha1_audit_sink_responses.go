// Code generated by go-swagger; DO NOT EDIT.

package auditregistration_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// WatchAuditregistrationV1alpha1AuditSinkOKCode is the HTTP code returned for type WatchAuditregistrationV1alpha1AuditSinkOK
const WatchAuditregistrationV1alpha1AuditSinkOKCode int = 200

/*WatchAuditregistrationV1alpha1AuditSinkOK OK

swagger:response watchAuditregistrationV1alpha1AuditSinkOK
*/
type WatchAuditregistrationV1alpha1AuditSinkOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchAuditregistrationV1alpha1AuditSinkOK creates WatchAuditregistrationV1alpha1AuditSinkOK with default headers values
func NewWatchAuditregistrationV1alpha1AuditSinkOK() *WatchAuditregistrationV1alpha1AuditSinkOK {

	return &WatchAuditregistrationV1alpha1AuditSinkOK{}
}

// WithPayload adds the payload to the watch auditregistration v1alpha1 audit sink o k response
func (o *WatchAuditregistrationV1alpha1AuditSinkOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchAuditregistrationV1alpha1AuditSinkOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch auditregistration v1alpha1 audit sink o k response
func (o *WatchAuditregistrationV1alpha1AuditSinkOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchAuditregistrationV1alpha1AuditSinkOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchAuditregistrationV1alpha1AuditSinkUnauthorizedCode is the HTTP code returned for type WatchAuditregistrationV1alpha1AuditSinkUnauthorized
const WatchAuditregistrationV1alpha1AuditSinkUnauthorizedCode int = 401

/*WatchAuditregistrationV1alpha1AuditSinkUnauthorized Unauthorized

swagger:response watchAuditregistrationV1alpha1AuditSinkUnauthorized
*/
type WatchAuditregistrationV1alpha1AuditSinkUnauthorized struct {
}

// NewWatchAuditregistrationV1alpha1AuditSinkUnauthorized creates WatchAuditregistrationV1alpha1AuditSinkUnauthorized with default headers values
func NewWatchAuditregistrationV1alpha1AuditSinkUnauthorized() *WatchAuditregistrationV1alpha1AuditSinkUnauthorized {

	return &WatchAuditregistrationV1alpha1AuditSinkUnauthorized{}
}

// WriteResponse to the client
func (o *WatchAuditregistrationV1alpha1AuditSinkUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}