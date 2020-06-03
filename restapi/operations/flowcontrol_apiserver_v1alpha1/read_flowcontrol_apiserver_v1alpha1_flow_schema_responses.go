// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadFlowcontrolApiserverV1alpha1FlowSchemaOKCode is the HTTP code returned for type ReadFlowcontrolApiserverV1alpha1FlowSchemaOK
const ReadFlowcontrolApiserverV1alpha1FlowSchemaOKCode int = 200

/*ReadFlowcontrolApiserverV1alpha1FlowSchemaOK OK

swagger:response readFlowcontrolApiserverV1alpha1FlowSchemaOK
*/
type ReadFlowcontrolApiserverV1alpha1FlowSchemaOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchema `json:"body,omitempty"`
}

// NewReadFlowcontrolApiserverV1alpha1FlowSchemaOK creates ReadFlowcontrolApiserverV1alpha1FlowSchemaOK with default headers values
func NewReadFlowcontrolApiserverV1alpha1FlowSchemaOK() *ReadFlowcontrolApiserverV1alpha1FlowSchemaOK {

	return &ReadFlowcontrolApiserverV1alpha1FlowSchemaOK{}
}

// WithPayload adds the payload to the read flowcontrol apiserver v1alpha1 flow schema o k response
func (o *ReadFlowcontrolApiserverV1alpha1FlowSchemaOK) WithPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchema) *ReadFlowcontrolApiserverV1alpha1FlowSchemaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read flowcontrol apiserver v1alpha1 flow schema o k response
func (o *ReadFlowcontrolApiserverV1alpha1FlowSchemaOK) SetPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchema) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadFlowcontrolApiserverV1alpha1FlowSchemaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadFlowcontrolApiserverV1alpha1FlowSchemaUnauthorizedCode is the HTTP code returned for type ReadFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized
const ReadFlowcontrolApiserverV1alpha1FlowSchemaUnauthorizedCode int = 401

/*ReadFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized Unauthorized

swagger:response readFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized
*/
type ReadFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized struct {
}

// NewReadFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized creates ReadFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized with default headers values
func NewReadFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized() *ReadFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized {

	return &ReadFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized{}
}

// WriteResponse to the client
func (o *ReadFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}