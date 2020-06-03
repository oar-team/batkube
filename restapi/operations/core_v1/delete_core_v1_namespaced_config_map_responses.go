// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// DeleteCoreV1NamespacedConfigMapOKCode is the HTTP code returned for type DeleteCoreV1NamespacedConfigMapOK
const DeleteCoreV1NamespacedConfigMapOKCode int = 200

/*DeleteCoreV1NamespacedConfigMapOK OK

swagger:response deleteCoreV1NamespacedConfigMapOK
*/
type DeleteCoreV1NamespacedConfigMapOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1NamespacedConfigMapOK creates DeleteCoreV1NamespacedConfigMapOK with default headers values
func NewDeleteCoreV1NamespacedConfigMapOK() *DeleteCoreV1NamespacedConfigMapOK {

	return &DeleteCoreV1NamespacedConfigMapOK{}
}

// WithPayload adds the payload to the delete core v1 namespaced config map o k response
func (o *DeleteCoreV1NamespacedConfigMapOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1NamespacedConfigMapOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 namespaced config map o k response
func (o *DeleteCoreV1NamespacedConfigMapOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedConfigMapOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NamespacedConfigMapAcceptedCode is the HTTP code returned for type DeleteCoreV1NamespacedConfigMapAccepted
const DeleteCoreV1NamespacedConfigMapAcceptedCode int = 202

/*DeleteCoreV1NamespacedConfigMapAccepted Accepted

swagger:response deleteCoreV1NamespacedConfigMapAccepted
*/
type DeleteCoreV1NamespacedConfigMapAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1NamespacedConfigMapAccepted creates DeleteCoreV1NamespacedConfigMapAccepted with default headers values
func NewDeleteCoreV1NamespacedConfigMapAccepted() *DeleteCoreV1NamespacedConfigMapAccepted {

	return &DeleteCoreV1NamespacedConfigMapAccepted{}
}

// WithPayload adds the payload to the delete core v1 namespaced config map accepted response
func (o *DeleteCoreV1NamespacedConfigMapAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1NamespacedConfigMapAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 namespaced config map accepted response
func (o *DeleteCoreV1NamespacedConfigMapAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedConfigMapAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NamespacedConfigMapUnauthorizedCode is the HTTP code returned for type DeleteCoreV1NamespacedConfigMapUnauthorized
const DeleteCoreV1NamespacedConfigMapUnauthorizedCode int = 401

/*DeleteCoreV1NamespacedConfigMapUnauthorized Unauthorized

swagger:response deleteCoreV1NamespacedConfigMapUnauthorized
*/
type DeleteCoreV1NamespacedConfigMapUnauthorized struct {
}

// NewDeleteCoreV1NamespacedConfigMapUnauthorized creates DeleteCoreV1NamespacedConfigMapUnauthorized with default headers values
func NewDeleteCoreV1NamespacedConfigMapUnauthorized() *DeleteCoreV1NamespacedConfigMapUnauthorized {

	return &DeleteCoreV1NamespacedConfigMapUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedConfigMapUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}