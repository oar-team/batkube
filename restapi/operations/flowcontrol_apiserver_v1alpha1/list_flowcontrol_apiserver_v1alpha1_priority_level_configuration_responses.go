// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOKCode is the HTTP code returned for type ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK
const ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOKCode int = 200

/*ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK OK

swagger:response listFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK
*/
type ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfigurationList `json:"body,omitempty"`
}

// NewListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK creates ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK with default headers values
func NewListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK() *ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK {

	return &ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK{}
}

// WithPayload adds the payload to the list flowcontrol apiserver v1alpha1 priority level configuration o k response
func (o *ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK) WithPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfigurationList) *ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list flowcontrol apiserver v1alpha1 priority level configuration o k response
func (o *ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK) SetPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfigurationList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorizedCode is the HTTP code returned for type ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized
const ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorizedCode int = 401

/*ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized Unauthorized

swagger:response listFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized
*/
type ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized struct {
}

// NewListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized creates ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized with default headers values
func NewListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized() *ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized {

	return &ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}