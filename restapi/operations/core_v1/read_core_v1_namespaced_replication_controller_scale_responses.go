// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadCoreV1NamespacedReplicationControllerScaleOKCode is the HTTP code returned for type ReadCoreV1NamespacedReplicationControllerScaleOK
const ReadCoreV1NamespacedReplicationControllerScaleOKCode int = 200

/*ReadCoreV1NamespacedReplicationControllerScaleOK OK

swagger:response readCoreV1NamespacedReplicationControllerScaleOK
*/
type ReadCoreV1NamespacedReplicationControllerScaleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1Scale `json:"body,omitempty"`
}

// NewReadCoreV1NamespacedReplicationControllerScaleOK creates ReadCoreV1NamespacedReplicationControllerScaleOK with default headers values
func NewReadCoreV1NamespacedReplicationControllerScaleOK() *ReadCoreV1NamespacedReplicationControllerScaleOK {

	return &ReadCoreV1NamespacedReplicationControllerScaleOK{}
}

// WithPayload adds the payload to the read core v1 namespaced replication controller scale o k response
func (o *ReadCoreV1NamespacedReplicationControllerScaleOK) WithPayload(payload *models.IoK8sAPIAutoscalingV1Scale) *ReadCoreV1NamespacedReplicationControllerScaleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read core v1 namespaced replication controller scale o k response
func (o *ReadCoreV1NamespacedReplicationControllerScaleOK) SetPayload(payload *models.IoK8sAPIAutoscalingV1Scale) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedReplicationControllerScaleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadCoreV1NamespacedReplicationControllerScaleUnauthorizedCode is the HTTP code returned for type ReadCoreV1NamespacedReplicationControllerScaleUnauthorized
const ReadCoreV1NamespacedReplicationControllerScaleUnauthorizedCode int = 401

/*ReadCoreV1NamespacedReplicationControllerScaleUnauthorized Unauthorized

swagger:response readCoreV1NamespacedReplicationControllerScaleUnauthorized
*/
type ReadCoreV1NamespacedReplicationControllerScaleUnauthorized struct {
}

// NewReadCoreV1NamespacedReplicationControllerScaleUnauthorized creates ReadCoreV1NamespacedReplicationControllerScaleUnauthorized with default headers values
func NewReadCoreV1NamespacedReplicationControllerScaleUnauthorized() *ReadCoreV1NamespacedReplicationControllerScaleUnauthorized {

	return &ReadCoreV1NamespacedReplicationControllerScaleUnauthorized{}
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedReplicationControllerScaleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
