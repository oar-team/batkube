// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOKCode is the HTTP code returned for type ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK
const ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOKCode int = 200

/*ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK OK

swagger:response replaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK
*/
type ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK creates ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK with default headers values
func NewReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK() *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK {

	return &ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK{}
}

// WithPayload adds the payload to the replace autoscaling v2beta2 namespaced horizontal pod autoscaler o k response
func (o *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler) *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace autoscaling v2beta2 namespaced horizontal pod autoscaler o k response
func (o *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreatedCode is the HTTP code returned for type ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated
const ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreatedCode int = 201

/*ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated Created

swagger:response replaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated
*/
type ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated creates ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated with default headers values
func NewReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated() *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated {

	return &ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated{}
}

// WithPayload adds the payload to the replace autoscaling v2beta2 namespaced horizontal pod autoscaler created response
func (o *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler) *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace autoscaling v2beta2 namespaced horizontal pod autoscaler created response
func (o *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorizedCode is the HTTP code returned for type ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized
const ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorizedCode int = 401

/*ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized Unauthorized

swagger:response replaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized
*/
type ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

// NewReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized creates ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized() *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized {

	return &ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
