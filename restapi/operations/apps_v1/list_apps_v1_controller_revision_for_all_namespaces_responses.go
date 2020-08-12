// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListAppsV1ControllerRevisionForAllNamespacesOKCode is the HTTP code returned for type ListAppsV1ControllerRevisionForAllNamespacesOK
const ListAppsV1ControllerRevisionForAllNamespacesOKCode int = 200

/*ListAppsV1ControllerRevisionForAllNamespacesOK OK

swagger:response listAppsV1ControllerRevisionForAllNamespacesOK
*/
type ListAppsV1ControllerRevisionForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1ControllerRevisionList `json:"body,omitempty"`
}

// NewListAppsV1ControllerRevisionForAllNamespacesOK creates ListAppsV1ControllerRevisionForAllNamespacesOK with default headers values
func NewListAppsV1ControllerRevisionForAllNamespacesOK() *ListAppsV1ControllerRevisionForAllNamespacesOK {

	return &ListAppsV1ControllerRevisionForAllNamespacesOK{}
}

// WithPayload adds the payload to the list apps v1 controller revision for all namespaces o k response
func (o *ListAppsV1ControllerRevisionForAllNamespacesOK) WithPayload(payload *models.IoK8sAPIAppsV1ControllerRevisionList) *ListAppsV1ControllerRevisionForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list apps v1 controller revision for all namespaces o k response
func (o *ListAppsV1ControllerRevisionForAllNamespacesOK) SetPayload(payload *models.IoK8sAPIAppsV1ControllerRevisionList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAppsV1ControllerRevisionForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAppsV1ControllerRevisionForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListAppsV1ControllerRevisionForAllNamespacesUnauthorized
const ListAppsV1ControllerRevisionForAllNamespacesUnauthorizedCode int = 401

/*ListAppsV1ControllerRevisionForAllNamespacesUnauthorized Unauthorized

swagger:response listAppsV1ControllerRevisionForAllNamespacesUnauthorized
*/
type ListAppsV1ControllerRevisionForAllNamespacesUnauthorized struct {
}

// NewListAppsV1ControllerRevisionForAllNamespacesUnauthorized creates ListAppsV1ControllerRevisionForAllNamespacesUnauthorized with default headers values
func NewListAppsV1ControllerRevisionForAllNamespacesUnauthorized() *ListAppsV1ControllerRevisionForAllNamespacesUnauthorized {

	return &ListAppsV1ControllerRevisionForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListAppsV1ControllerRevisionForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
