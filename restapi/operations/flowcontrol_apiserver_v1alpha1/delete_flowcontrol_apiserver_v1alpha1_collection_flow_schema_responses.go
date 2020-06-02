// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOKCode is the HTTP code returned for type DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK
const DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOKCode int = 200

/*DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK OK

swagger:response deleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK
*/
type DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK creates DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK with default headers values
func NewDeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK() *DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK {

	return &DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK{}
}

// WithPayload adds the payload to the delete flowcontrol apiserver v1alpha1 collection flow schema o k response
func (o *DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete flowcontrol apiserver v1alpha1 collection flow schema o k response
func (o *DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaUnauthorizedCode is the HTTP code returned for type DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaUnauthorized
const DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaUnauthorizedCode int = 401

/*DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaUnauthorized Unauthorized

swagger:response deleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaUnauthorized
*/
type DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaUnauthorized struct {
}

// NewDeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaUnauthorized creates DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaUnauthorized with default headers values
func NewDeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaUnauthorized() *DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaUnauthorized {

	return &DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
