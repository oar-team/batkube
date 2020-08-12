// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOKCode is the HTTP code returned for type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK
const ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOKCode int = 200

/*ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK OK

swagger:response replaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK
*/
type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK creates ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK with default headers values
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK() *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK {

	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK{}
}

// WithPayload adds the payload to the replace autoscaling v1 namespaced horizontal pod autoscaler status o k response
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK) WithPayload(payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace autoscaling v1 namespaced horizontal pod autoscaler status o k response
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK) SetPayload(payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreatedCode is the HTTP code returned for type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated
const ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreatedCode int = 201

/*ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated Created

swagger:response replaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated
*/
type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated creates ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated with default headers values
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated() *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated {

	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated{}
}

// WithPayload adds the payload to the replace autoscaling v1 namespaced horizontal pod autoscaler status created response
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated) WithPayload(payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace autoscaling v1 namespaced horizontal pod autoscaler status created response
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated) SetPayload(payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusUnauthorizedCode is the HTTP code returned for type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusUnauthorized
const ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusUnauthorizedCode int = 401

/*ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusUnauthorized Unauthorized

swagger:response replaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusUnauthorized
*/
type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusUnauthorized struct {
}

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusUnauthorized creates ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusUnauthorized with default headers values
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusUnauthorized() *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusUnauthorized {

	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
