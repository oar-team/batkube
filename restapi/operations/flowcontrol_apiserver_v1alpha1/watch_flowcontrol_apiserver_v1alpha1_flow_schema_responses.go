// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchFlowcontrolApiserverV1alpha1FlowSchemaOKCode is the HTTP code returned for type WatchFlowcontrolApiserverV1alpha1FlowSchemaOK
const WatchFlowcontrolApiserverV1alpha1FlowSchemaOKCode int = 200

/*WatchFlowcontrolApiserverV1alpha1FlowSchemaOK OK

swagger:response watchFlowcontrolApiserverV1alpha1FlowSchemaOK
*/
type WatchFlowcontrolApiserverV1alpha1FlowSchemaOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchFlowcontrolApiserverV1alpha1FlowSchemaOK creates WatchFlowcontrolApiserverV1alpha1FlowSchemaOK with default headers values
func NewWatchFlowcontrolApiserverV1alpha1FlowSchemaOK() *WatchFlowcontrolApiserverV1alpha1FlowSchemaOK {

	return &WatchFlowcontrolApiserverV1alpha1FlowSchemaOK{}
}

// WithPayload adds the payload to the watch flowcontrol apiserver v1alpha1 flow schema o k response
func (o *WatchFlowcontrolApiserverV1alpha1FlowSchemaOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchFlowcontrolApiserverV1alpha1FlowSchemaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch flowcontrol apiserver v1alpha1 flow schema o k response
func (o *WatchFlowcontrolApiserverV1alpha1FlowSchemaOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchFlowcontrolApiserverV1alpha1FlowSchemaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorizedCode is the HTTP code returned for type WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized
const WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorizedCode int = 401

/*WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized Unauthorized

swagger:response watchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized
*/
type WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized struct {
}

// NewWatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized creates WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized with default headers values
func NewWatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized() *WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized {

	return &WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized{}
}

// WriteResponse to the client
func (o *WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
