// Code generated by go-swagger; DO NOT EDIT.

package auditregistration_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// DeleteAuditregistrationV1alpha1AuditSinkOKCode is the HTTP code returned for type DeleteAuditregistrationV1alpha1AuditSinkOK
const DeleteAuditregistrationV1alpha1AuditSinkOKCode int = 200

/*DeleteAuditregistrationV1alpha1AuditSinkOK OK

swagger:response deleteAuditregistrationV1alpha1AuditSinkOK
*/
type DeleteAuditregistrationV1alpha1AuditSinkOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAuditregistrationV1alpha1AuditSinkOK creates DeleteAuditregistrationV1alpha1AuditSinkOK with default headers values
func NewDeleteAuditregistrationV1alpha1AuditSinkOK() *DeleteAuditregistrationV1alpha1AuditSinkOK {

	return &DeleteAuditregistrationV1alpha1AuditSinkOK{}
}

// WithPayload adds the payload to the delete auditregistration v1alpha1 audit sink o k response
func (o *DeleteAuditregistrationV1alpha1AuditSinkOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAuditregistrationV1alpha1AuditSinkOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete auditregistration v1alpha1 audit sink o k response
func (o *DeleteAuditregistrationV1alpha1AuditSinkOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAuditregistrationV1alpha1AuditSinkOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAuditregistrationV1alpha1AuditSinkAcceptedCode is the HTTP code returned for type DeleteAuditregistrationV1alpha1AuditSinkAccepted
const DeleteAuditregistrationV1alpha1AuditSinkAcceptedCode int = 202

/*DeleteAuditregistrationV1alpha1AuditSinkAccepted Accepted

swagger:response deleteAuditregistrationV1alpha1AuditSinkAccepted
*/
type DeleteAuditregistrationV1alpha1AuditSinkAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAuditregistrationV1alpha1AuditSinkAccepted creates DeleteAuditregistrationV1alpha1AuditSinkAccepted with default headers values
func NewDeleteAuditregistrationV1alpha1AuditSinkAccepted() *DeleteAuditregistrationV1alpha1AuditSinkAccepted {

	return &DeleteAuditregistrationV1alpha1AuditSinkAccepted{}
}

// WithPayload adds the payload to the delete auditregistration v1alpha1 audit sink accepted response
func (o *DeleteAuditregistrationV1alpha1AuditSinkAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAuditregistrationV1alpha1AuditSinkAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete auditregistration v1alpha1 audit sink accepted response
func (o *DeleteAuditregistrationV1alpha1AuditSinkAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAuditregistrationV1alpha1AuditSinkAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAuditregistrationV1alpha1AuditSinkUnauthorizedCode is the HTTP code returned for type DeleteAuditregistrationV1alpha1AuditSinkUnauthorized
const DeleteAuditregistrationV1alpha1AuditSinkUnauthorizedCode int = 401

/*DeleteAuditregistrationV1alpha1AuditSinkUnauthorized Unauthorized

swagger:response deleteAuditregistrationV1alpha1AuditSinkUnauthorized
*/
type DeleteAuditregistrationV1alpha1AuditSinkUnauthorized struct {
}

// NewDeleteAuditregistrationV1alpha1AuditSinkUnauthorized creates DeleteAuditregistrationV1alpha1AuditSinkUnauthorized with default headers values
func NewDeleteAuditregistrationV1alpha1AuditSinkUnauthorized() *DeleteAuditregistrationV1alpha1AuditSinkUnauthorized {

	return &DeleteAuditregistrationV1alpha1AuditSinkUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteAuditregistrationV1alpha1AuditSinkUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}