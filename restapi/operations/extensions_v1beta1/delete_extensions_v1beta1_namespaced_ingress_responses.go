// Code generated by go-swagger; DO NOT EDIT.

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteExtensionsV1beta1NamespacedIngressOKCode is the HTTP code returned for type DeleteExtensionsV1beta1NamespacedIngressOK
const DeleteExtensionsV1beta1NamespacedIngressOKCode int = 200

/*DeleteExtensionsV1beta1NamespacedIngressOK OK

swagger:response deleteExtensionsV1beta1NamespacedIngressOK
*/
type DeleteExtensionsV1beta1NamespacedIngressOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteExtensionsV1beta1NamespacedIngressOK creates DeleteExtensionsV1beta1NamespacedIngressOK with default headers values
func NewDeleteExtensionsV1beta1NamespacedIngressOK() *DeleteExtensionsV1beta1NamespacedIngressOK {

	return &DeleteExtensionsV1beta1NamespacedIngressOK{}
}

// WithPayload adds the payload to the delete extensions v1beta1 namespaced ingress o k response
func (o *DeleteExtensionsV1beta1NamespacedIngressOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteExtensionsV1beta1NamespacedIngressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete extensions v1beta1 namespaced ingress o k response
func (o *DeleteExtensionsV1beta1NamespacedIngressOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteExtensionsV1beta1NamespacedIngressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteExtensionsV1beta1NamespacedIngressAcceptedCode is the HTTP code returned for type DeleteExtensionsV1beta1NamespacedIngressAccepted
const DeleteExtensionsV1beta1NamespacedIngressAcceptedCode int = 202

/*DeleteExtensionsV1beta1NamespacedIngressAccepted Accepted

swagger:response deleteExtensionsV1beta1NamespacedIngressAccepted
*/
type DeleteExtensionsV1beta1NamespacedIngressAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteExtensionsV1beta1NamespacedIngressAccepted creates DeleteExtensionsV1beta1NamespacedIngressAccepted with default headers values
func NewDeleteExtensionsV1beta1NamespacedIngressAccepted() *DeleteExtensionsV1beta1NamespacedIngressAccepted {

	return &DeleteExtensionsV1beta1NamespacedIngressAccepted{}
}

// WithPayload adds the payload to the delete extensions v1beta1 namespaced ingress accepted response
func (o *DeleteExtensionsV1beta1NamespacedIngressAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteExtensionsV1beta1NamespacedIngressAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete extensions v1beta1 namespaced ingress accepted response
func (o *DeleteExtensionsV1beta1NamespacedIngressAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteExtensionsV1beta1NamespacedIngressAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteExtensionsV1beta1NamespacedIngressUnauthorizedCode is the HTTP code returned for type DeleteExtensionsV1beta1NamespacedIngressUnauthorized
const DeleteExtensionsV1beta1NamespacedIngressUnauthorizedCode int = 401

/*DeleteExtensionsV1beta1NamespacedIngressUnauthorized Unauthorized

swagger:response deleteExtensionsV1beta1NamespacedIngressUnauthorized
*/
type DeleteExtensionsV1beta1NamespacedIngressUnauthorized struct {
}

// NewDeleteExtensionsV1beta1NamespacedIngressUnauthorized creates DeleteExtensionsV1beta1NamespacedIngressUnauthorized with default headers values
func NewDeleteExtensionsV1beta1NamespacedIngressUnauthorized() *DeleteExtensionsV1beta1NamespacedIngressUnauthorized {

	return &DeleteExtensionsV1beta1NamespacedIngressUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteExtensionsV1beta1NamespacedIngressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
