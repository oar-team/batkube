// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// PatchCoreV1NamespaceOKCode is the HTTP code returned for type PatchCoreV1NamespaceOK
const PatchCoreV1NamespaceOKCode int = 200

/*PatchCoreV1NamespaceOK OK

swagger:response patchCoreV1NamespaceOK
*/
type PatchCoreV1NamespaceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Namespace `json:"body,omitempty"`
}

// NewPatchCoreV1NamespaceOK creates PatchCoreV1NamespaceOK with default headers values
func NewPatchCoreV1NamespaceOK() *PatchCoreV1NamespaceOK {

	return &PatchCoreV1NamespaceOK{}
}

// WithPayload adds the payload to the patch core v1 namespace o k response
func (o *PatchCoreV1NamespaceOK) WithPayload(payload *models.IoK8sAPICoreV1Namespace) *PatchCoreV1NamespaceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch core v1 namespace o k response
func (o *PatchCoreV1NamespaceOK) SetPayload(payload *models.IoK8sAPICoreV1Namespace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoreV1NamespaceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoreV1NamespaceUnauthorizedCode is the HTTP code returned for type PatchCoreV1NamespaceUnauthorized
const PatchCoreV1NamespaceUnauthorizedCode int = 401

/*PatchCoreV1NamespaceUnauthorized Unauthorized

swagger:response patchCoreV1NamespaceUnauthorized
*/
type PatchCoreV1NamespaceUnauthorized struct {
}

// NewPatchCoreV1NamespaceUnauthorized creates PatchCoreV1NamespaceUnauthorized with default headers values
func NewPatchCoreV1NamespaceUnauthorized() *PatchCoreV1NamespaceUnauthorized {

	return &PatchCoreV1NamespaceUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoreV1NamespaceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
