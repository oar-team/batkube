// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOKCode is the HTTP code returned for type ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK
const ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOKCode int = 200

/*ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK OK

swagger:response listAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK
*/
type ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerList `json:"body,omitempty"`
}

// NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK creates ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK with default headers values
func NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK() *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK {

	return &ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK{}
}

// WithPayload adds the payload to the list autoscaling v2beta2 horizontal pod autoscaler for all namespaces o k response
func (o *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerList) *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list autoscaling v2beta2 horizontal pod autoscaler for all namespaces o k response
func (o *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized
const ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorizedCode int = 401

/*ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized Unauthorized

swagger:response listAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized
*/
type ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized struct {
}

// NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized creates ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized with default headers values
func NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized() *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized {

	return &ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
