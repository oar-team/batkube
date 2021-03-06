// Code generated by go-swagger; DO NOT EDIT.

package discovery_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOKCode is the HTTP code returned for type DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK
const DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOKCode int = 200

/*DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK OK

swagger:response deleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK
*/
type DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK creates DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK with default headers values
func NewDeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK() *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK {

	return &DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK{}
}

// WithPayload adds the payload to the delete discovery v1beta1 collection namespaced endpoint slice o k response
func (o *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete discovery v1beta1 collection namespaced endpoint slice o k response
func (o *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorizedCode is the HTTP code returned for type DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized
const DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorizedCode int = 401

/*DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized Unauthorized

swagger:response deleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized
*/
type DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized struct {
}

// NewDeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized creates DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized with default headers values
func NewDeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized() *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized {

	return &DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
