// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOKCode is the HTTP code returned for type ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK
const ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOKCode int = 200

/*ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK OK

swagger:response replaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK
*/
type ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration `json:"body,omitempty"`
}

// NewReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK creates ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK with default headers values
func NewReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK() *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK {

	return &ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK{}
}

// WithPayload adds the payload to the replace flowcontrol apiserver v1alpha1 priority level configuration status o k response
func (o *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK) WithPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration) *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace flowcontrol apiserver v1alpha1 priority level configuration status o k response
func (o *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK) SetPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreatedCode is the HTTP code returned for type ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated
const ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreatedCode int = 201

/*ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated Created

swagger:response replaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated
*/
type ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration `json:"body,omitempty"`
}

// NewReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated creates ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated with default headers values
func NewReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated() *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated {

	return &ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated{}
}

// WithPayload adds the payload to the replace flowcontrol apiserver v1alpha1 priority level configuration status created response
func (o *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated) WithPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration) *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace flowcontrol apiserver v1alpha1 priority level configuration status created response
func (o *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated) SetPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusUnauthorizedCode is the HTTP code returned for type ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusUnauthorized
const ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusUnauthorizedCode int = 401

/*ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusUnauthorized Unauthorized

swagger:response replaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusUnauthorized
*/
type ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusUnauthorized struct {
}

// NewReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusUnauthorized creates ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusUnauthorized with default headers values
func NewReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusUnauthorized() *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusUnauthorized {

	return &ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
