// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOKCode is the HTTP code returned for type DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK
const DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOKCode int = 200

/*DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK OK

swagger:response deleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK
*/
type DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK creates DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK with default headers values
func NewDeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK() *DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK {

	return &DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK{}
}

// WithPayload adds the payload to the delete autoscaling v2beta1 namespaced horizontal pod autoscaler o k response
func (o *DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete autoscaling v2beta1 namespaced horizontal pod autoscaler o k response
func (o *DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAcceptedCode is the HTTP code returned for type DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted
const DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAcceptedCode int = 202

/*DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted Accepted

swagger:response deleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted
*/
type DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted creates DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted with default headers values
func NewDeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted() *DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted {

	return &DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted{}
}

// WithPayload adds the payload to the delete autoscaling v2beta1 namespaced horizontal pod autoscaler accepted response
func (o *DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete autoscaling v2beta1 namespaced horizontal pod autoscaler accepted response
func (o *DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorizedCode is the HTTP code returned for type DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized
const DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorizedCode int = 401

/*DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized Unauthorized

swagger:response deleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized
*/
type DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

// NewDeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized creates DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewDeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized() *DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized {

	return &DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
