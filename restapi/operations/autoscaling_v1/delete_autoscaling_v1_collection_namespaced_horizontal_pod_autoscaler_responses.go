// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOKCode is the HTTP code returned for type DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK
const DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOKCode int = 200

/*DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK OK

swagger:response deleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK
*/
type DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK creates DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK with default headers values
func NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK() *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK {

	return &DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK{}
}

// WithPayload adds the payload to the delete autoscaling v1 collection namespaced horizontal pod autoscaler o k response
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete autoscaling v1 collection namespaced horizontal pod autoscaler o k response
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorizedCode is the HTTP code returned for type DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized
const DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorizedCode int = 401

/*DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized Unauthorized

swagger:response deleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized
*/
type DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized struct {
}

// NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized creates DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized() *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized {

	return &DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
