// Code generated by go-swagger; DO NOT EDIT.

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListPolicyV1beta1PodSecurityPolicyOKCode is the HTTP code returned for type ListPolicyV1beta1PodSecurityPolicyOK
const ListPolicyV1beta1PodSecurityPolicyOKCode int = 200

/*ListPolicyV1beta1PodSecurityPolicyOK OK

swagger:response listPolicyV1beta1PodSecurityPolicyOK
*/
type ListPolicyV1beta1PodSecurityPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicyList `json:"body,omitempty"`
}

// NewListPolicyV1beta1PodSecurityPolicyOK creates ListPolicyV1beta1PodSecurityPolicyOK with default headers values
func NewListPolicyV1beta1PodSecurityPolicyOK() *ListPolicyV1beta1PodSecurityPolicyOK {

	return &ListPolicyV1beta1PodSecurityPolicyOK{}
}

// WithPayload adds the payload to the list policy v1beta1 pod security policy o k response
func (o *ListPolicyV1beta1PodSecurityPolicyOK) WithPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicyList) *ListPolicyV1beta1PodSecurityPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list policy v1beta1 pod security policy o k response
func (o *ListPolicyV1beta1PodSecurityPolicyOK) SetPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicyList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPolicyV1beta1PodSecurityPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListPolicyV1beta1PodSecurityPolicyUnauthorizedCode is the HTTP code returned for type ListPolicyV1beta1PodSecurityPolicyUnauthorized
const ListPolicyV1beta1PodSecurityPolicyUnauthorizedCode int = 401

/*ListPolicyV1beta1PodSecurityPolicyUnauthorized Unauthorized

swagger:response listPolicyV1beta1PodSecurityPolicyUnauthorized
*/
type ListPolicyV1beta1PodSecurityPolicyUnauthorized struct {
}

// NewListPolicyV1beta1PodSecurityPolicyUnauthorized creates ListPolicyV1beta1PodSecurityPolicyUnauthorized with default headers values
func NewListPolicyV1beta1PodSecurityPolicyUnauthorized() *ListPolicyV1beta1PodSecurityPolicyUnauthorized {

	return &ListPolicyV1beta1PodSecurityPolicyUnauthorized{}
}

// WriteResponse to the client
func (o *ListPolicyV1beta1PodSecurityPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
