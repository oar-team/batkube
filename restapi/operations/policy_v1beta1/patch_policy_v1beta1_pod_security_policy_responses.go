// Code generated by go-swagger; DO NOT EDIT.

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchPolicyV1beta1PodSecurityPolicyOKCode is the HTTP code returned for type PatchPolicyV1beta1PodSecurityPolicyOK
const PatchPolicyV1beta1PodSecurityPolicyOKCode int = 200

/*PatchPolicyV1beta1PodSecurityPolicyOK OK

swagger:response patchPolicyV1beta1PodSecurityPolicyOK
*/
type PatchPolicyV1beta1PodSecurityPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy `json:"body,omitempty"`
}

// NewPatchPolicyV1beta1PodSecurityPolicyOK creates PatchPolicyV1beta1PodSecurityPolicyOK with default headers values
func NewPatchPolicyV1beta1PodSecurityPolicyOK() *PatchPolicyV1beta1PodSecurityPolicyOK {

	return &PatchPolicyV1beta1PodSecurityPolicyOK{}
}

// WithPayload adds the payload to the patch policy v1beta1 pod security policy o k response
func (o *PatchPolicyV1beta1PodSecurityPolicyOK) WithPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) *PatchPolicyV1beta1PodSecurityPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch policy v1beta1 pod security policy o k response
func (o *PatchPolicyV1beta1PodSecurityPolicyOK) SetPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchPolicyV1beta1PodSecurityPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchPolicyV1beta1PodSecurityPolicyUnauthorizedCode is the HTTP code returned for type PatchPolicyV1beta1PodSecurityPolicyUnauthorized
const PatchPolicyV1beta1PodSecurityPolicyUnauthorizedCode int = 401

/*PatchPolicyV1beta1PodSecurityPolicyUnauthorized Unauthorized

swagger:response patchPolicyV1beta1PodSecurityPolicyUnauthorized
*/
type PatchPolicyV1beta1PodSecurityPolicyUnauthorized struct {
}

// NewPatchPolicyV1beta1PodSecurityPolicyUnauthorized creates PatchPolicyV1beta1PodSecurityPolicyUnauthorized with default headers values
func NewPatchPolicyV1beta1PodSecurityPolicyUnauthorized() *PatchPolicyV1beta1PodSecurityPolicyUnauthorized {

	return &PatchPolicyV1beta1PodSecurityPolicyUnauthorized{}
}

// WriteResponse to the client
func (o *PatchPolicyV1beta1PodSecurityPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
