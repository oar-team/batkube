// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOKCode is the HTTP code returned for type PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK
const PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOKCode int = 200

/*PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK OK

swagger:response patchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK
*/
type PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration `json:"body,omitempty"`
}

// NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK creates PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK with default headers values
func NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK() *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK {

	return &PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK{}
}

// WithPayload adds the payload to the patch flowcontrol apiserver v1alpha1 priority level configuration o k response
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK) WithPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch flowcontrol apiserver v1alpha1 priority level configuration o k response
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK) SetPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorizedCode is the HTTP code returned for type PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized
const PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorizedCode int = 401

/*PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized Unauthorized

swagger:response patchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized
*/
type PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized struct {
}

// NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized creates PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized with default headers values
func NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized() *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized {

	return &PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
