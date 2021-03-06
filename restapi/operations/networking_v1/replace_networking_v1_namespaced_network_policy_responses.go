// Code generated by go-swagger; DO NOT EDIT.

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceNetworkingV1NamespacedNetworkPolicyOKCode is the HTTP code returned for type ReplaceNetworkingV1NamespacedNetworkPolicyOK
const ReplaceNetworkingV1NamespacedNetworkPolicyOKCode int = 200

/*ReplaceNetworkingV1NamespacedNetworkPolicyOK OK

swagger:response replaceNetworkingV1NamespacedNetworkPolicyOK
*/
type ReplaceNetworkingV1NamespacedNetworkPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINetworkingV1NetworkPolicy `json:"body,omitempty"`
}

// NewReplaceNetworkingV1NamespacedNetworkPolicyOK creates ReplaceNetworkingV1NamespacedNetworkPolicyOK with default headers values
func NewReplaceNetworkingV1NamespacedNetworkPolicyOK() *ReplaceNetworkingV1NamespacedNetworkPolicyOK {

	return &ReplaceNetworkingV1NamespacedNetworkPolicyOK{}
}

// WithPayload adds the payload to the replace networking v1 namespaced network policy o k response
func (o *ReplaceNetworkingV1NamespacedNetworkPolicyOK) WithPayload(payload *models.IoK8sAPINetworkingV1NetworkPolicy) *ReplaceNetworkingV1NamespacedNetworkPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace networking v1 namespaced network policy o k response
func (o *ReplaceNetworkingV1NamespacedNetworkPolicyOK) SetPayload(payload *models.IoK8sAPINetworkingV1NetworkPolicy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceNetworkingV1NamespacedNetworkPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceNetworkingV1NamespacedNetworkPolicyCreatedCode is the HTTP code returned for type ReplaceNetworkingV1NamespacedNetworkPolicyCreated
const ReplaceNetworkingV1NamespacedNetworkPolicyCreatedCode int = 201

/*ReplaceNetworkingV1NamespacedNetworkPolicyCreated Created

swagger:response replaceNetworkingV1NamespacedNetworkPolicyCreated
*/
type ReplaceNetworkingV1NamespacedNetworkPolicyCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINetworkingV1NetworkPolicy `json:"body,omitempty"`
}

// NewReplaceNetworkingV1NamespacedNetworkPolicyCreated creates ReplaceNetworkingV1NamespacedNetworkPolicyCreated with default headers values
func NewReplaceNetworkingV1NamespacedNetworkPolicyCreated() *ReplaceNetworkingV1NamespacedNetworkPolicyCreated {

	return &ReplaceNetworkingV1NamespacedNetworkPolicyCreated{}
}

// WithPayload adds the payload to the replace networking v1 namespaced network policy created response
func (o *ReplaceNetworkingV1NamespacedNetworkPolicyCreated) WithPayload(payload *models.IoK8sAPINetworkingV1NetworkPolicy) *ReplaceNetworkingV1NamespacedNetworkPolicyCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace networking v1 namespaced network policy created response
func (o *ReplaceNetworkingV1NamespacedNetworkPolicyCreated) SetPayload(payload *models.IoK8sAPINetworkingV1NetworkPolicy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceNetworkingV1NamespacedNetworkPolicyCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceNetworkingV1NamespacedNetworkPolicyUnauthorizedCode is the HTTP code returned for type ReplaceNetworkingV1NamespacedNetworkPolicyUnauthorized
const ReplaceNetworkingV1NamespacedNetworkPolicyUnauthorizedCode int = 401

/*ReplaceNetworkingV1NamespacedNetworkPolicyUnauthorized Unauthorized

swagger:response replaceNetworkingV1NamespacedNetworkPolicyUnauthorized
*/
type ReplaceNetworkingV1NamespacedNetworkPolicyUnauthorized struct {
}

// NewReplaceNetworkingV1NamespacedNetworkPolicyUnauthorized creates ReplaceNetworkingV1NamespacedNetworkPolicyUnauthorized with default headers values
func NewReplaceNetworkingV1NamespacedNetworkPolicyUnauthorized() *ReplaceNetworkingV1NamespacedNetworkPolicyUnauthorized {

	return &ReplaceNetworkingV1NamespacedNetworkPolicyUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceNetworkingV1NamespacedNetworkPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
