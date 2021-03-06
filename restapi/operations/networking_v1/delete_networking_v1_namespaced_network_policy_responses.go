// Code generated by go-swagger; DO NOT EDIT.

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteNetworkingV1NamespacedNetworkPolicyOKCode is the HTTP code returned for type DeleteNetworkingV1NamespacedNetworkPolicyOK
const DeleteNetworkingV1NamespacedNetworkPolicyOKCode int = 200

/*DeleteNetworkingV1NamespacedNetworkPolicyOK OK

swagger:response deleteNetworkingV1NamespacedNetworkPolicyOK
*/
type DeleteNetworkingV1NamespacedNetworkPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteNetworkingV1NamespacedNetworkPolicyOK creates DeleteNetworkingV1NamespacedNetworkPolicyOK with default headers values
func NewDeleteNetworkingV1NamespacedNetworkPolicyOK() *DeleteNetworkingV1NamespacedNetworkPolicyOK {

	return &DeleteNetworkingV1NamespacedNetworkPolicyOK{}
}

// WithPayload adds the payload to the delete networking v1 namespaced network policy o k response
func (o *DeleteNetworkingV1NamespacedNetworkPolicyOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteNetworkingV1NamespacedNetworkPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete networking v1 namespaced network policy o k response
func (o *DeleteNetworkingV1NamespacedNetworkPolicyOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNetworkingV1NamespacedNetworkPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteNetworkingV1NamespacedNetworkPolicyAcceptedCode is the HTTP code returned for type DeleteNetworkingV1NamespacedNetworkPolicyAccepted
const DeleteNetworkingV1NamespacedNetworkPolicyAcceptedCode int = 202

/*DeleteNetworkingV1NamespacedNetworkPolicyAccepted Accepted

swagger:response deleteNetworkingV1NamespacedNetworkPolicyAccepted
*/
type DeleteNetworkingV1NamespacedNetworkPolicyAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteNetworkingV1NamespacedNetworkPolicyAccepted creates DeleteNetworkingV1NamespacedNetworkPolicyAccepted with default headers values
func NewDeleteNetworkingV1NamespacedNetworkPolicyAccepted() *DeleteNetworkingV1NamespacedNetworkPolicyAccepted {

	return &DeleteNetworkingV1NamespacedNetworkPolicyAccepted{}
}

// WithPayload adds the payload to the delete networking v1 namespaced network policy accepted response
func (o *DeleteNetworkingV1NamespacedNetworkPolicyAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteNetworkingV1NamespacedNetworkPolicyAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete networking v1 namespaced network policy accepted response
func (o *DeleteNetworkingV1NamespacedNetworkPolicyAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNetworkingV1NamespacedNetworkPolicyAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteNetworkingV1NamespacedNetworkPolicyUnauthorizedCode is the HTTP code returned for type DeleteNetworkingV1NamespacedNetworkPolicyUnauthorized
const DeleteNetworkingV1NamespacedNetworkPolicyUnauthorizedCode int = 401

/*DeleteNetworkingV1NamespacedNetworkPolicyUnauthorized Unauthorized

swagger:response deleteNetworkingV1NamespacedNetworkPolicyUnauthorized
*/
type DeleteNetworkingV1NamespacedNetworkPolicyUnauthorized struct {
}

// NewDeleteNetworkingV1NamespacedNetworkPolicyUnauthorized creates DeleteNetworkingV1NamespacedNetworkPolicyUnauthorized with default headers values
func NewDeleteNetworkingV1NamespacedNetworkPolicyUnauthorized() *DeleteNetworkingV1NamespacedNetworkPolicyUnauthorized {

	return &DeleteNetworkingV1NamespacedNetworkPolicyUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteNetworkingV1NamespacedNetworkPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
