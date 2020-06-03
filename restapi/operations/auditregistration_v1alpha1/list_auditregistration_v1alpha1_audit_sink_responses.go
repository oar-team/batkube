// Code generated by go-swagger; DO NOT EDIT.

package auditregistration_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListAuditregistrationV1alpha1AuditSinkOKCode is the HTTP code returned for type ListAuditregistrationV1alpha1AuditSinkOK
const ListAuditregistrationV1alpha1AuditSinkOKCode int = 200

/*ListAuditregistrationV1alpha1AuditSinkOK OK

swagger:response listAuditregistrationV1alpha1AuditSinkOK
*/
type ListAuditregistrationV1alpha1AuditSinkOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAuditregistrationV1alpha1AuditSinkList `json:"body,omitempty"`
}

// NewListAuditregistrationV1alpha1AuditSinkOK creates ListAuditregistrationV1alpha1AuditSinkOK with default headers values
func NewListAuditregistrationV1alpha1AuditSinkOK() *ListAuditregistrationV1alpha1AuditSinkOK {

	return &ListAuditregistrationV1alpha1AuditSinkOK{}
}

// WithPayload adds the payload to the list auditregistration v1alpha1 audit sink o k response
func (o *ListAuditregistrationV1alpha1AuditSinkOK) WithPayload(payload *models.IoK8sAPIAuditregistrationV1alpha1AuditSinkList) *ListAuditregistrationV1alpha1AuditSinkOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list auditregistration v1alpha1 audit sink o k response
func (o *ListAuditregistrationV1alpha1AuditSinkOK) SetPayload(payload *models.IoK8sAPIAuditregistrationV1alpha1AuditSinkList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAuditregistrationV1alpha1AuditSinkOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAuditregistrationV1alpha1AuditSinkUnauthorizedCode is the HTTP code returned for type ListAuditregistrationV1alpha1AuditSinkUnauthorized
const ListAuditregistrationV1alpha1AuditSinkUnauthorizedCode int = 401

/*ListAuditregistrationV1alpha1AuditSinkUnauthorized Unauthorized

swagger:response listAuditregistrationV1alpha1AuditSinkUnauthorized
*/
type ListAuditregistrationV1alpha1AuditSinkUnauthorized struct {
}

// NewListAuditregistrationV1alpha1AuditSinkUnauthorized creates ListAuditregistrationV1alpha1AuditSinkUnauthorized with default headers values
func NewListAuditregistrationV1alpha1AuditSinkUnauthorized() *ListAuditregistrationV1alpha1AuditSinkUnauthorized {

	return &ListAuditregistrationV1alpha1AuditSinkUnauthorized{}
}

// WriteResponse to the client
func (o *ListAuditregistrationV1alpha1AuditSinkUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}