// Code generated by go-swagger; DO NOT EDIT.

package auditregistration_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// WatchAuditregistrationV1alpha1AuditSinkListOKCode is the HTTP code returned for type WatchAuditregistrationV1alpha1AuditSinkListOK
const WatchAuditregistrationV1alpha1AuditSinkListOKCode int = 200

/*WatchAuditregistrationV1alpha1AuditSinkListOK OK

swagger:response watchAuditregistrationV1alpha1AuditSinkListOK
*/
type WatchAuditregistrationV1alpha1AuditSinkListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchAuditregistrationV1alpha1AuditSinkListOK creates WatchAuditregistrationV1alpha1AuditSinkListOK with default headers values
func NewWatchAuditregistrationV1alpha1AuditSinkListOK() *WatchAuditregistrationV1alpha1AuditSinkListOK {

	return &WatchAuditregistrationV1alpha1AuditSinkListOK{}
}

// WithPayload adds the payload to the watch auditregistration v1alpha1 audit sink list o k response
func (o *WatchAuditregistrationV1alpha1AuditSinkListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchAuditregistrationV1alpha1AuditSinkListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch auditregistration v1alpha1 audit sink list o k response
func (o *WatchAuditregistrationV1alpha1AuditSinkListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchAuditregistrationV1alpha1AuditSinkListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchAuditregistrationV1alpha1AuditSinkListUnauthorizedCode is the HTTP code returned for type WatchAuditregistrationV1alpha1AuditSinkListUnauthorized
const WatchAuditregistrationV1alpha1AuditSinkListUnauthorizedCode int = 401

/*WatchAuditregistrationV1alpha1AuditSinkListUnauthorized Unauthorized

swagger:response watchAuditregistrationV1alpha1AuditSinkListUnauthorized
*/
type WatchAuditregistrationV1alpha1AuditSinkListUnauthorized struct {
}

// NewWatchAuditregistrationV1alpha1AuditSinkListUnauthorized creates WatchAuditregistrationV1alpha1AuditSinkListUnauthorized with default headers values
func NewWatchAuditregistrationV1alpha1AuditSinkListUnauthorized() *WatchAuditregistrationV1alpha1AuditSinkListUnauthorized {

	return &WatchAuditregistrationV1alpha1AuditSinkListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchAuditregistrationV1alpha1AuditSinkListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}