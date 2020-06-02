// Code generated by go-swagger; DO NOT EDIT.

package settings_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListSettingsV1alpha1NamespacedPodPresetOKCode is the HTTP code returned for type ListSettingsV1alpha1NamespacedPodPresetOK
const ListSettingsV1alpha1NamespacedPodPresetOKCode int = 200

/*ListSettingsV1alpha1NamespacedPodPresetOK OK

swagger:response listSettingsV1alpha1NamespacedPodPresetOK
*/
type ListSettingsV1alpha1NamespacedPodPresetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPISettingsV1alpha1PodPresetList `json:"body,omitempty"`
}

// NewListSettingsV1alpha1NamespacedPodPresetOK creates ListSettingsV1alpha1NamespacedPodPresetOK with default headers values
func NewListSettingsV1alpha1NamespacedPodPresetOK() *ListSettingsV1alpha1NamespacedPodPresetOK {

	return &ListSettingsV1alpha1NamespacedPodPresetOK{}
}

// WithPayload adds the payload to the list settings v1alpha1 namespaced pod preset o k response
func (o *ListSettingsV1alpha1NamespacedPodPresetOK) WithPayload(payload *models.IoK8sAPISettingsV1alpha1PodPresetList) *ListSettingsV1alpha1NamespacedPodPresetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list settings v1alpha1 namespaced pod preset o k response
func (o *ListSettingsV1alpha1NamespacedPodPresetOK) SetPayload(payload *models.IoK8sAPISettingsV1alpha1PodPresetList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListSettingsV1alpha1NamespacedPodPresetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListSettingsV1alpha1NamespacedPodPresetUnauthorizedCode is the HTTP code returned for type ListSettingsV1alpha1NamespacedPodPresetUnauthorized
const ListSettingsV1alpha1NamespacedPodPresetUnauthorizedCode int = 401

/*ListSettingsV1alpha1NamespacedPodPresetUnauthorized Unauthorized

swagger:response listSettingsV1alpha1NamespacedPodPresetUnauthorized
*/
type ListSettingsV1alpha1NamespacedPodPresetUnauthorized struct {
}

// NewListSettingsV1alpha1NamespacedPodPresetUnauthorized creates ListSettingsV1alpha1NamespacedPodPresetUnauthorized with default headers values
func NewListSettingsV1alpha1NamespacedPodPresetUnauthorized() *ListSettingsV1alpha1NamespacedPodPresetUnauthorized {

	return &ListSettingsV1alpha1NamespacedPodPresetUnauthorized{}
}

// WriteResponse to the client
func (o *ListSettingsV1alpha1NamespacedPodPresetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
