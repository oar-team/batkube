// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchAppsV1NamespacedDaemonSetStatusOKCode is the HTTP code returned for type PatchAppsV1NamespacedDaemonSetStatusOK
const PatchAppsV1NamespacedDaemonSetStatusOKCode int = 200

/*PatchAppsV1NamespacedDaemonSetStatusOK OK

swagger:response patchAppsV1NamespacedDaemonSetStatusOK
*/
type PatchAppsV1NamespacedDaemonSetStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1DaemonSet `json:"body,omitempty"`
}

// NewPatchAppsV1NamespacedDaemonSetStatusOK creates PatchAppsV1NamespacedDaemonSetStatusOK with default headers values
func NewPatchAppsV1NamespacedDaemonSetStatusOK() *PatchAppsV1NamespacedDaemonSetStatusOK {

	return &PatchAppsV1NamespacedDaemonSetStatusOK{}
}

// WithPayload adds the payload to the patch apps v1 namespaced daemon set status o k response
func (o *PatchAppsV1NamespacedDaemonSetStatusOK) WithPayload(payload *models.IoK8sAPIAppsV1DaemonSet) *PatchAppsV1NamespacedDaemonSetStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apps v1 namespaced daemon set status o k response
func (o *PatchAppsV1NamespacedDaemonSetStatusOK) SetPayload(payload *models.IoK8sAPIAppsV1DaemonSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAppsV1NamespacedDaemonSetStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchAppsV1NamespacedDaemonSetStatusUnauthorizedCode is the HTTP code returned for type PatchAppsV1NamespacedDaemonSetStatusUnauthorized
const PatchAppsV1NamespacedDaemonSetStatusUnauthorizedCode int = 401

/*PatchAppsV1NamespacedDaemonSetStatusUnauthorized Unauthorized

swagger:response patchAppsV1NamespacedDaemonSetStatusUnauthorized
*/
type PatchAppsV1NamespacedDaemonSetStatusUnauthorized struct {
}

// NewPatchAppsV1NamespacedDaemonSetStatusUnauthorized creates PatchAppsV1NamespacedDaemonSetStatusUnauthorized with default headers values
func NewPatchAppsV1NamespacedDaemonSetStatusUnauthorized() *PatchAppsV1NamespacedDaemonSetStatusUnauthorized {

	return &PatchAppsV1NamespacedDaemonSetStatusUnauthorized{}
}

// WriteResponse to the client
func (o *PatchAppsV1NamespacedDaemonSetStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
