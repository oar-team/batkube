// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// PatchCoreV1NamespacedSecretOKCode is the HTTP code returned for type PatchCoreV1NamespacedSecretOK
const PatchCoreV1NamespacedSecretOKCode int = 200

/*PatchCoreV1NamespacedSecretOK OK

swagger:response patchCoreV1NamespacedSecretOK
*/
type PatchCoreV1NamespacedSecretOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Secret `json:"body,omitempty"`
}

// NewPatchCoreV1NamespacedSecretOK creates PatchCoreV1NamespacedSecretOK with default headers values
func NewPatchCoreV1NamespacedSecretOK() *PatchCoreV1NamespacedSecretOK {

	return &PatchCoreV1NamespacedSecretOK{}
}

// WithPayload adds the payload to the patch core v1 namespaced secret o k response
func (o *PatchCoreV1NamespacedSecretOK) WithPayload(payload *models.IoK8sAPICoreV1Secret) *PatchCoreV1NamespacedSecretOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch core v1 namespaced secret o k response
func (o *PatchCoreV1NamespacedSecretOK) SetPayload(payload *models.IoK8sAPICoreV1Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedSecretOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoreV1NamespacedSecretUnauthorizedCode is the HTTP code returned for type PatchCoreV1NamespacedSecretUnauthorized
const PatchCoreV1NamespacedSecretUnauthorizedCode int = 401

/*PatchCoreV1NamespacedSecretUnauthorized Unauthorized

swagger:response patchCoreV1NamespacedSecretUnauthorized
*/
type PatchCoreV1NamespacedSecretUnauthorized struct {
}

// NewPatchCoreV1NamespacedSecretUnauthorized creates PatchCoreV1NamespacedSecretUnauthorized with default headers values
func NewPatchCoreV1NamespacedSecretUnauthorized() *PatchCoreV1NamespacedSecretUnauthorized {

	return &PatchCoreV1NamespacedSecretUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedSecretUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
