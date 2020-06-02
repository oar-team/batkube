// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListAppsV1StatefulSetForAllNamespacesOKCode is the HTTP code returned for type ListAppsV1StatefulSetForAllNamespacesOK
const ListAppsV1StatefulSetForAllNamespacesOKCode int = 200

/*ListAppsV1StatefulSetForAllNamespacesOK OK

swagger:response listAppsV1StatefulSetForAllNamespacesOK
*/
type ListAppsV1StatefulSetForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1StatefulSetList `json:"body,omitempty"`
}

// NewListAppsV1StatefulSetForAllNamespacesOK creates ListAppsV1StatefulSetForAllNamespacesOK with default headers values
func NewListAppsV1StatefulSetForAllNamespacesOK() *ListAppsV1StatefulSetForAllNamespacesOK {

	return &ListAppsV1StatefulSetForAllNamespacesOK{}
}

// WithPayload adds the payload to the list apps v1 stateful set for all namespaces o k response
func (o *ListAppsV1StatefulSetForAllNamespacesOK) WithPayload(payload *models.IoK8sAPIAppsV1StatefulSetList) *ListAppsV1StatefulSetForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list apps v1 stateful set for all namespaces o k response
func (o *ListAppsV1StatefulSetForAllNamespacesOK) SetPayload(payload *models.IoK8sAPIAppsV1StatefulSetList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAppsV1StatefulSetForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAppsV1StatefulSetForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListAppsV1StatefulSetForAllNamespacesUnauthorized
const ListAppsV1StatefulSetForAllNamespacesUnauthorizedCode int = 401

/*ListAppsV1StatefulSetForAllNamespacesUnauthorized Unauthorized

swagger:response listAppsV1StatefulSetForAllNamespacesUnauthorized
*/
type ListAppsV1StatefulSetForAllNamespacesUnauthorized struct {
}

// NewListAppsV1StatefulSetForAllNamespacesUnauthorized creates ListAppsV1StatefulSetForAllNamespacesUnauthorized with default headers values
func NewListAppsV1StatefulSetForAllNamespacesUnauthorized() *ListAppsV1StatefulSetForAllNamespacesUnauthorized {

	return &ListAppsV1StatefulSetForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListAppsV1StatefulSetForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
