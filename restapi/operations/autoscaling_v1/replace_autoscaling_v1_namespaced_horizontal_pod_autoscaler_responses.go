// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOKCode is the HTTP code returned for type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK
const ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOKCode int = 200

/*ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK OK

swagger:response replaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK
*/
type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK creates ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK with default headers values
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK() *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK {

	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK{}
}

// WithPayload adds the payload to the replace autoscaling v1 namespaced horizontal pod autoscaler o k response
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK) WithPayload(payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace autoscaling v1 namespaced horizontal pod autoscaler o k response
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK) SetPayload(payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreatedCode is the HTTP code returned for type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated
const ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreatedCode int = 201

/*ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated Created

swagger:response replaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated
*/
type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated creates ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated with default headers values
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated() *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated {

	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated{}
}

// WithPayload adds the payload to the replace autoscaling v1 namespaced horizontal pod autoscaler created response
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated) WithPayload(payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace autoscaling v1 namespaced horizontal pod autoscaler created response
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated) SetPayload(payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorizedCode is the HTTP code returned for type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized
const ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorizedCode int = 401

/*ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized Unauthorized

swagger:response replaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized
*/
type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized creates ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized() *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized {

	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
