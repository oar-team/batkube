// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListCoreV1NamespacedSecretOKCode is the HTTP code returned for type ListCoreV1NamespacedSecretOK
const ListCoreV1NamespacedSecretOKCode int = 200

/*ListCoreV1NamespacedSecretOK OK

swagger:response listCoreV1NamespacedSecretOK
*/
type ListCoreV1NamespacedSecretOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1SecretList `json:"body,omitempty"`
}

// NewListCoreV1NamespacedSecretOK creates ListCoreV1NamespacedSecretOK with default headers values
func NewListCoreV1NamespacedSecretOK() *ListCoreV1NamespacedSecretOK {

	return &ListCoreV1NamespacedSecretOK{}
}

// WithPayload adds the payload to the list core v1 namespaced secret o k response
func (o *ListCoreV1NamespacedSecretOK) WithPayload(payload *models.IoK8sAPICoreV1SecretList) *ListCoreV1NamespacedSecretOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list core v1 namespaced secret o k response
func (o *ListCoreV1NamespacedSecretOK) SetPayload(payload *models.IoK8sAPICoreV1SecretList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCoreV1NamespacedSecretOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCoreV1NamespacedSecretUnauthorizedCode is the HTTP code returned for type ListCoreV1NamespacedSecretUnauthorized
const ListCoreV1NamespacedSecretUnauthorizedCode int = 401

/*ListCoreV1NamespacedSecretUnauthorized Unauthorized

swagger:response listCoreV1NamespacedSecretUnauthorized
*/
type ListCoreV1NamespacedSecretUnauthorized struct {
}

// NewListCoreV1NamespacedSecretUnauthorized creates ListCoreV1NamespacedSecretUnauthorized with default headers values
func NewListCoreV1NamespacedSecretUnauthorized() *ListCoreV1NamespacedSecretUnauthorized {

	return &ListCoreV1NamespacedSecretUnauthorized{}
}

// WriteResponse to the client
func (o *ListCoreV1NamespacedSecretUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
