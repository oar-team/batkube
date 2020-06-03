// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// DeleteCoreV1NamespacedSecretOKCode is the HTTP code returned for type DeleteCoreV1NamespacedSecretOK
const DeleteCoreV1NamespacedSecretOKCode int = 200

/*DeleteCoreV1NamespacedSecretOK OK

swagger:response deleteCoreV1NamespacedSecretOK
*/
type DeleteCoreV1NamespacedSecretOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1NamespacedSecretOK creates DeleteCoreV1NamespacedSecretOK with default headers values
func NewDeleteCoreV1NamespacedSecretOK() *DeleteCoreV1NamespacedSecretOK {

	return &DeleteCoreV1NamespacedSecretOK{}
}

// WithPayload adds the payload to the delete core v1 namespaced secret o k response
func (o *DeleteCoreV1NamespacedSecretOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1NamespacedSecretOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 namespaced secret o k response
func (o *DeleteCoreV1NamespacedSecretOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedSecretOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NamespacedSecretAcceptedCode is the HTTP code returned for type DeleteCoreV1NamespacedSecretAccepted
const DeleteCoreV1NamespacedSecretAcceptedCode int = 202

/*DeleteCoreV1NamespacedSecretAccepted Accepted

swagger:response deleteCoreV1NamespacedSecretAccepted
*/
type DeleteCoreV1NamespacedSecretAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1NamespacedSecretAccepted creates DeleteCoreV1NamespacedSecretAccepted with default headers values
func NewDeleteCoreV1NamespacedSecretAccepted() *DeleteCoreV1NamespacedSecretAccepted {

	return &DeleteCoreV1NamespacedSecretAccepted{}
}

// WithPayload adds the payload to the delete core v1 namespaced secret accepted response
func (o *DeleteCoreV1NamespacedSecretAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1NamespacedSecretAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 namespaced secret accepted response
func (o *DeleteCoreV1NamespacedSecretAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedSecretAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NamespacedSecretUnauthorizedCode is the HTTP code returned for type DeleteCoreV1NamespacedSecretUnauthorized
const DeleteCoreV1NamespacedSecretUnauthorizedCode int = 401

/*DeleteCoreV1NamespacedSecretUnauthorized Unauthorized

swagger:response deleteCoreV1NamespacedSecretUnauthorized
*/
type DeleteCoreV1NamespacedSecretUnauthorized struct {
}

// NewDeleteCoreV1NamespacedSecretUnauthorized creates DeleteCoreV1NamespacedSecretUnauthorized with default headers values
func NewDeleteCoreV1NamespacedSecretUnauthorized() *DeleteCoreV1NamespacedSecretUnauthorized {

	return &DeleteCoreV1NamespacedSecretUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedSecretUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}