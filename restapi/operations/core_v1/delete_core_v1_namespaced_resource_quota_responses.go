// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteCoreV1NamespacedResourceQuotaOKCode is the HTTP code returned for type DeleteCoreV1NamespacedResourceQuotaOK
const DeleteCoreV1NamespacedResourceQuotaOKCode int = 200

/*DeleteCoreV1NamespacedResourceQuotaOK OK

swagger:response deleteCoreV1NamespacedResourceQuotaOK
*/
type DeleteCoreV1NamespacedResourceQuotaOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ResourceQuota `json:"body,omitempty"`
}

// NewDeleteCoreV1NamespacedResourceQuotaOK creates DeleteCoreV1NamespacedResourceQuotaOK with default headers values
func NewDeleteCoreV1NamespacedResourceQuotaOK() *DeleteCoreV1NamespacedResourceQuotaOK {

	return &DeleteCoreV1NamespacedResourceQuotaOK{}
}

// WithPayload adds the payload to the delete core v1 namespaced resource quota o k response
func (o *DeleteCoreV1NamespacedResourceQuotaOK) WithPayload(payload *models.IoK8sAPICoreV1ResourceQuota) *DeleteCoreV1NamespacedResourceQuotaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 namespaced resource quota o k response
func (o *DeleteCoreV1NamespacedResourceQuotaOK) SetPayload(payload *models.IoK8sAPICoreV1ResourceQuota) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedResourceQuotaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NamespacedResourceQuotaAcceptedCode is the HTTP code returned for type DeleteCoreV1NamespacedResourceQuotaAccepted
const DeleteCoreV1NamespacedResourceQuotaAcceptedCode int = 202

/*DeleteCoreV1NamespacedResourceQuotaAccepted Accepted

swagger:response deleteCoreV1NamespacedResourceQuotaAccepted
*/
type DeleteCoreV1NamespacedResourceQuotaAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ResourceQuota `json:"body,omitempty"`
}

// NewDeleteCoreV1NamespacedResourceQuotaAccepted creates DeleteCoreV1NamespacedResourceQuotaAccepted with default headers values
func NewDeleteCoreV1NamespacedResourceQuotaAccepted() *DeleteCoreV1NamespacedResourceQuotaAccepted {

	return &DeleteCoreV1NamespacedResourceQuotaAccepted{}
}

// WithPayload adds the payload to the delete core v1 namespaced resource quota accepted response
func (o *DeleteCoreV1NamespacedResourceQuotaAccepted) WithPayload(payload *models.IoK8sAPICoreV1ResourceQuota) *DeleteCoreV1NamespacedResourceQuotaAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 namespaced resource quota accepted response
func (o *DeleteCoreV1NamespacedResourceQuotaAccepted) SetPayload(payload *models.IoK8sAPICoreV1ResourceQuota) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedResourceQuotaAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NamespacedResourceQuotaUnauthorizedCode is the HTTP code returned for type DeleteCoreV1NamespacedResourceQuotaUnauthorized
const DeleteCoreV1NamespacedResourceQuotaUnauthorizedCode int = 401

/*DeleteCoreV1NamespacedResourceQuotaUnauthorized Unauthorized

swagger:response deleteCoreV1NamespacedResourceQuotaUnauthorized
*/
type DeleteCoreV1NamespacedResourceQuotaUnauthorized struct {
}

// NewDeleteCoreV1NamespacedResourceQuotaUnauthorized creates DeleteCoreV1NamespacedResourceQuotaUnauthorized with default headers values
func NewDeleteCoreV1NamespacedResourceQuotaUnauthorized() *DeleteCoreV1NamespacedResourceQuotaUnauthorized {

	return &DeleteCoreV1NamespacedResourceQuotaUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedResourceQuotaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
