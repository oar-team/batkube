// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListAutoscalingV1NamespacedHorizontalPodAutoscalerOKCode is the HTTP code returned for type ListAutoscalingV1NamespacedHorizontalPodAutoscalerOK
const ListAutoscalingV1NamespacedHorizontalPodAutoscalerOKCode int = 200

/*ListAutoscalingV1NamespacedHorizontalPodAutoscalerOK OK

swagger:response listAutoscalingV1NamespacedHorizontalPodAutoscalerOK
*/
type ListAutoscalingV1NamespacedHorizontalPodAutoscalerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscalerList `json:"body,omitempty"`
}

// NewListAutoscalingV1NamespacedHorizontalPodAutoscalerOK creates ListAutoscalingV1NamespacedHorizontalPodAutoscalerOK with default headers values
func NewListAutoscalingV1NamespacedHorizontalPodAutoscalerOK() *ListAutoscalingV1NamespacedHorizontalPodAutoscalerOK {

	return &ListAutoscalingV1NamespacedHorizontalPodAutoscalerOK{}
}

// WithPayload adds the payload to the list autoscaling v1 namespaced horizontal pod autoscaler o k response
func (o *ListAutoscalingV1NamespacedHorizontalPodAutoscalerOK) WithPayload(payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscalerList) *ListAutoscalingV1NamespacedHorizontalPodAutoscalerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list autoscaling v1 namespaced horizontal pod autoscaler o k response
func (o *ListAutoscalingV1NamespacedHorizontalPodAutoscalerOK) SetPayload(payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscalerList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAutoscalingV1NamespacedHorizontalPodAutoscalerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorizedCode is the HTTP code returned for type ListAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized
const ListAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorizedCode int = 401

/*ListAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized Unauthorized

swagger:response listAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized
*/
type ListAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

// NewListAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized creates ListAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewListAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized() *ListAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized {

	return &ListAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized{}
}

// WriteResponse to the client
func (o *ListAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
