// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReadAutoscalingV1NamespacedHorizontalPodAutoscalerOKCode is the HTTP code returned for type ReadAutoscalingV1NamespacedHorizontalPodAutoscalerOK
const ReadAutoscalingV1NamespacedHorizontalPodAutoscalerOKCode int = 200

/*ReadAutoscalingV1NamespacedHorizontalPodAutoscalerOK OK

swagger:response readAutoscalingV1NamespacedHorizontalPodAutoscalerOK
*/
type ReadAutoscalingV1NamespacedHorizontalPodAutoscalerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerOK creates ReadAutoscalingV1NamespacedHorizontalPodAutoscalerOK with default headers values
func NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerOK() *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerOK {

	return &ReadAutoscalingV1NamespacedHorizontalPodAutoscalerOK{}
}

// WithPayload adds the payload to the read autoscaling v1 namespaced horizontal pod autoscaler o k response
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerOK) WithPayload(payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler) *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read autoscaling v1 namespaced horizontal pod autoscaler o k response
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerOK) SetPayload(payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorizedCode is the HTTP code returned for type ReadAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized
const ReadAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorizedCode int = 401

/*ReadAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized Unauthorized

swagger:response readAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized
*/
type ReadAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

// NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized creates ReadAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized() *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized {

	return &ReadAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized{}
}

// WriteResponse to the client
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
