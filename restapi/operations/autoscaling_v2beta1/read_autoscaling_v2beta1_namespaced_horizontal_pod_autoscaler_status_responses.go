// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOKCode is the HTTP code returned for type ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK
const ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOKCode int = 200

/*ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK OK

swagger:response readAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK
*/
type ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK creates ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK with default headers values
func NewReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK() *ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK {

	return &ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK{}
}

// WithPayload adds the payload to the read autoscaling v2beta1 namespaced horizontal pod autoscaler status o k response
func (o *ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) *ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read autoscaling v2beta1 namespaced horizontal pod autoscaler status o k response
func (o *ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorizedCode is the HTTP code returned for type ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized
const ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorizedCode int = 401

/*ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized Unauthorized

swagger:response readAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized
*/
type ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized struct {
}

// NewReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized creates ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized with default headers values
func NewReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized() *ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized {

	return &ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
