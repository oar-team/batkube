// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOKCode is the HTTP code returned for type ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK
const ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOKCode int = 200

/*ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK OK

swagger:response replaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK
*/
type ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK creates ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK with default headers values
func NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK() *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK {

	return &ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK{}
}

// WithPayload adds the payload to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler o k response
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler o k response
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreatedCode is the HTTP code returned for type ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated
const ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreatedCode int = 201

/*ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated Created

swagger:response replaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated
*/
type ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated creates ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated with default headers values
func NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated() *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated {

	return &ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated{}
}

// WithPayload adds the payload to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler created response
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler created response
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorizedCode is the HTTP code returned for type ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized
const ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorizedCode int = 401

/*ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized Unauthorized

swagger:response replaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized
*/
type ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

// NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized creates ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized() *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized {

	return &ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
