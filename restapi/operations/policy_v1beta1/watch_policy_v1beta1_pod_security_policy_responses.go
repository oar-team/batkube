// Code generated by go-swagger; DO NOT EDIT.

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// WatchPolicyV1beta1PodSecurityPolicyOKCode is the HTTP code returned for type WatchPolicyV1beta1PodSecurityPolicyOK
const WatchPolicyV1beta1PodSecurityPolicyOKCode int = 200

/*WatchPolicyV1beta1PodSecurityPolicyOK OK

swagger:response watchPolicyV1beta1PodSecurityPolicyOK
*/
type WatchPolicyV1beta1PodSecurityPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchPolicyV1beta1PodSecurityPolicyOK creates WatchPolicyV1beta1PodSecurityPolicyOK with default headers values
func NewWatchPolicyV1beta1PodSecurityPolicyOK() *WatchPolicyV1beta1PodSecurityPolicyOK {

	return &WatchPolicyV1beta1PodSecurityPolicyOK{}
}

// WithPayload adds the payload to the watch policy v1beta1 pod security policy o k response
func (o *WatchPolicyV1beta1PodSecurityPolicyOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchPolicyV1beta1PodSecurityPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch policy v1beta1 pod security policy o k response
func (o *WatchPolicyV1beta1PodSecurityPolicyOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchPolicyV1beta1PodSecurityPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchPolicyV1beta1PodSecurityPolicyUnauthorizedCode is the HTTP code returned for type WatchPolicyV1beta1PodSecurityPolicyUnauthorized
const WatchPolicyV1beta1PodSecurityPolicyUnauthorizedCode int = 401

/*WatchPolicyV1beta1PodSecurityPolicyUnauthorized Unauthorized

swagger:response watchPolicyV1beta1PodSecurityPolicyUnauthorized
*/
type WatchPolicyV1beta1PodSecurityPolicyUnauthorized struct {
}

// NewWatchPolicyV1beta1PodSecurityPolicyUnauthorized creates WatchPolicyV1beta1PodSecurityPolicyUnauthorized with default headers values
func NewWatchPolicyV1beta1PodSecurityPolicyUnauthorized() *WatchPolicyV1beta1PodSecurityPolicyUnauthorized {

	return &WatchPolicyV1beta1PodSecurityPolicyUnauthorized{}
}

// WriteResponse to the client
func (o *WatchPolicyV1beta1PodSecurityPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}