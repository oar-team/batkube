// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOKCode is the HTTP code returned for type DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK
const DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOKCode int = 200

/*DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK OK

swagger:response deleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK
*/
type DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK creates DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK with default headers values
func NewDeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK() *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK {

	return &DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK{}
}

// WithPayload adds the payload to the delete autoscaling v2beta2 namespaced horizontal pod autoscaler o k response
func (o *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete autoscaling v2beta2 namespaced horizontal pod autoscaler o k response
func (o *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAcceptedCode is the HTTP code returned for type DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted
const DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAcceptedCode int = 202

/*DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted Accepted

swagger:response deleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted
*/
type DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted creates DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted with default headers values
func NewDeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted() *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted {

	return &DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted{}
}

// WithPayload adds the payload to the delete autoscaling v2beta2 namespaced horizontal pod autoscaler accepted response
func (o *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete autoscaling v2beta2 namespaced horizontal pod autoscaler accepted response
func (o *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorizedCode is the HTTP code returned for type DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized
const DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorizedCode int = 401

/*DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized Unauthorized

swagger:response deleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized
*/
type DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

// NewDeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized creates DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewDeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized() *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized {

	return &DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
