// Code generated by go-swagger; DO NOT EDIT.

package auditregistration_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchAuditregistrationV1alpha1AuditSinkOKCode is the HTTP code returned for type PatchAuditregistrationV1alpha1AuditSinkOK
const PatchAuditregistrationV1alpha1AuditSinkOKCode int = 200

/*PatchAuditregistrationV1alpha1AuditSinkOK OK

swagger:response patchAuditregistrationV1alpha1AuditSinkOK
*/
type PatchAuditregistrationV1alpha1AuditSinkOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAuditregistrationV1alpha1AuditSink `json:"body,omitempty"`
}

// NewPatchAuditregistrationV1alpha1AuditSinkOK creates PatchAuditregistrationV1alpha1AuditSinkOK with default headers values
func NewPatchAuditregistrationV1alpha1AuditSinkOK() *PatchAuditregistrationV1alpha1AuditSinkOK {

	return &PatchAuditregistrationV1alpha1AuditSinkOK{}
}

// WithPayload adds the payload to the patch auditregistration v1alpha1 audit sink o k response
func (o *PatchAuditregistrationV1alpha1AuditSinkOK) WithPayload(payload *models.IoK8sAPIAuditregistrationV1alpha1AuditSink) *PatchAuditregistrationV1alpha1AuditSinkOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch auditregistration v1alpha1 audit sink o k response
func (o *PatchAuditregistrationV1alpha1AuditSinkOK) SetPayload(payload *models.IoK8sAPIAuditregistrationV1alpha1AuditSink) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAuditregistrationV1alpha1AuditSinkOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchAuditregistrationV1alpha1AuditSinkUnauthorizedCode is the HTTP code returned for type PatchAuditregistrationV1alpha1AuditSinkUnauthorized
const PatchAuditregistrationV1alpha1AuditSinkUnauthorizedCode int = 401

/*PatchAuditregistrationV1alpha1AuditSinkUnauthorized Unauthorized

swagger:response patchAuditregistrationV1alpha1AuditSinkUnauthorized
*/
type PatchAuditregistrationV1alpha1AuditSinkUnauthorized struct {
}

// NewPatchAuditregistrationV1alpha1AuditSinkUnauthorized creates PatchAuditregistrationV1alpha1AuditSinkUnauthorized with default headers values
func NewPatchAuditregistrationV1alpha1AuditSinkUnauthorized() *PatchAuditregistrationV1alpha1AuditSinkUnauthorized {

	return &PatchAuditregistrationV1alpha1AuditSinkUnauthorized{}
}

// WriteResponse to the client
func (o *PatchAuditregistrationV1alpha1AuditSinkUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
