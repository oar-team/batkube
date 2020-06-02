// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// PatchAppsV1NamespacedStatefulSetOKCode is the HTTP code returned for type PatchAppsV1NamespacedStatefulSetOK
const PatchAppsV1NamespacedStatefulSetOKCode int = 200

/*PatchAppsV1NamespacedStatefulSetOK OK

swagger:response patchAppsV1NamespacedStatefulSetOK
*/
type PatchAppsV1NamespacedStatefulSetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1StatefulSet `json:"body,omitempty"`
}

// NewPatchAppsV1NamespacedStatefulSetOK creates PatchAppsV1NamespacedStatefulSetOK with default headers values
func NewPatchAppsV1NamespacedStatefulSetOK() *PatchAppsV1NamespacedStatefulSetOK {

	return &PatchAppsV1NamespacedStatefulSetOK{}
}

// WithPayload adds the payload to the patch apps v1 namespaced stateful set o k response
func (o *PatchAppsV1NamespacedStatefulSetOK) WithPayload(payload *models.IoK8sAPIAppsV1StatefulSet) *PatchAppsV1NamespacedStatefulSetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apps v1 namespaced stateful set o k response
func (o *PatchAppsV1NamespacedStatefulSetOK) SetPayload(payload *models.IoK8sAPIAppsV1StatefulSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAppsV1NamespacedStatefulSetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchAppsV1NamespacedStatefulSetUnauthorizedCode is the HTTP code returned for type PatchAppsV1NamespacedStatefulSetUnauthorized
const PatchAppsV1NamespacedStatefulSetUnauthorizedCode int = 401

/*PatchAppsV1NamespacedStatefulSetUnauthorized Unauthorized

swagger:response patchAppsV1NamespacedStatefulSetUnauthorized
*/
type PatchAppsV1NamespacedStatefulSetUnauthorized struct {
}

// NewPatchAppsV1NamespacedStatefulSetUnauthorized creates PatchAppsV1NamespacedStatefulSetUnauthorized with default headers values
func NewPatchAppsV1NamespacedStatefulSetUnauthorized() *PatchAppsV1NamespacedStatefulSetUnauthorized {

	return &PatchAppsV1NamespacedStatefulSetUnauthorized{}
}

// WriteResponse to the client
func (o *PatchAppsV1NamespacedStatefulSetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
