// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteCoreV1NamespacedReplicationControllerOKCode is the HTTP code returned for type DeleteCoreV1NamespacedReplicationControllerOK
const DeleteCoreV1NamespacedReplicationControllerOKCode int = 200

/*DeleteCoreV1NamespacedReplicationControllerOK OK

swagger:response deleteCoreV1NamespacedReplicationControllerOK
*/
type DeleteCoreV1NamespacedReplicationControllerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1NamespacedReplicationControllerOK creates DeleteCoreV1NamespacedReplicationControllerOK with default headers values
func NewDeleteCoreV1NamespacedReplicationControllerOK() *DeleteCoreV1NamespacedReplicationControllerOK {

	return &DeleteCoreV1NamespacedReplicationControllerOK{}
}

// WithPayload adds the payload to the delete core v1 namespaced replication controller o k response
func (o *DeleteCoreV1NamespacedReplicationControllerOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1NamespacedReplicationControllerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 namespaced replication controller o k response
func (o *DeleteCoreV1NamespacedReplicationControllerOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedReplicationControllerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NamespacedReplicationControllerAcceptedCode is the HTTP code returned for type DeleteCoreV1NamespacedReplicationControllerAccepted
const DeleteCoreV1NamespacedReplicationControllerAcceptedCode int = 202

/*DeleteCoreV1NamespacedReplicationControllerAccepted Accepted

swagger:response deleteCoreV1NamespacedReplicationControllerAccepted
*/
type DeleteCoreV1NamespacedReplicationControllerAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1NamespacedReplicationControllerAccepted creates DeleteCoreV1NamespacedReplicationControllerAccepted with default headers values
func NewDeleteCoreV1NamespacedReplicationControllerAccepted() *DeleteCoreV1NamespacedReplicationControllerAccepted {

	return &DeleteCoreV1NamespacedReplicationControllerAccepted{}
}

// WithPayload adds the payload to the delete core v1 namespaced replication controller accepted response
func (o *DeleteCoreV1NamespacedReplicationControllerAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1NamespacedReplicationControllerAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 namespaced replication controller accepted response
func (o *DeleteCoreV1NamespacedReplicationControllerAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedReplicationControllerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NamespacedReplicationControllerUnauthorizedCode is the HTTP code returned for type DeleteCoreV1NamespacedReplicationControllerUnauthorized
const DeleteCoreV1NamespacedReplicationControllerUnauthorizedCode int = 401

/*DeleteCoreV1NamespacedReplicationControllerUnauthorized Unauthorized

swagger:response deleteCoreV1NamespacedReplicationControllerUnauthorized
*/
type DeleteCoreV1NamespacedReplicationControllerUnauthorized struct {
}

// NewDeleteCoreV1NamespacedReplicationControllerUnauthorized creates DeleteCoreV1NamespacedReplicationControllerUnauthorized with default headers values
func NewDeleteCoreV1NamespacedReplicationControllerUnauthorized() *DeleteCoreV1NamespacedReplicationControllerUnauthorized {

	return &DeleteCoreV1NamespacedReplicationControllerUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedReplicationControllerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
