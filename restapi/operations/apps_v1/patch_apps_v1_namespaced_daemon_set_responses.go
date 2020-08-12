// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchAppsV1NamespacedDaemonSetOKCode is the HTTP code returned for type PatchAppsV1NamespacedDaemonSetOK
const PatchAppsV1NamespacedDaemonSetOKCode int = 200

/*PatchAppsV1NamespacedDaemonSetOK OK

swagger:response patchAppsV1NamespacedDaemonSetOK
*/
type PatchAppsV1NamespacedDaemonSetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1DaemonSet `json:"body,omitempty"`
}

// NewPatchAppsV1NamespacedDaemonSetOK creates PatchAppsV1NamespacedDaemonSetOK with default headers values
func NewPatchAppsV1NamespacedDaemonSetOK() *PatchAppsV1NamespacedDaemonSetOK {

	return &PatchAppsV1NamespacedDaemonSetOK{}
}

// WithPayload adds the payload to the patch apps v1 namespaced daemon set o k response
func (o *PatchAppsV1NamespacedDaemonSetOK) WithPayload(payload *models.IoK8sAPIAppsV1DaemonSet) *PatchAppsV1NamespacedDaemonSetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apps v1 namespaced daemon set o k response
func (o *PatchAppsV1NamespacedDaemonSetOK) SetPayload(payload *models.IoK8sAPIAppsV1DaemonSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAppsV1NamespacedDaemonSetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchAppsV1NamespacedDaemonSetUnauthorizedCode is the HTTP code returned for type PatchAppsV1NamespacedDaemonSetUnauthorized
const PatchAppsV1NamespacedDaemonSetUnauthorizedCode int = 401

/*PatchAppsV1NamespacedDaemonSetUnauthorized Unauthorized

swagger:response patchAppsV1NamespacedDaemonSetUnauthorized
*/
type PatchAppsV1NamespacedDaemonSetUnauthorized struct {
}

// NewPatchAppsV1NamespacedDaemonSetUnauthorized creates PatchAppsV1NamespacedDaemonSetUnauthorized with default headers values
func NewPatchAppsV1NamespacedDaemonSetUnauthorized() *PatchAppsV1NamespacedDaemonSetUnauthorized {

	return &PatchAppsV1NamespacedDaemonSetUnauthorized{}
}

// WriteResponse to the client
func (o *PatchAppsV1NamespacedDaemonSetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
