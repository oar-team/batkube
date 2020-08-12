// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchStorageV1beta1CSINodeOKCode is the HTTP code returned for type WatchStorageV1beta1CSINodeOK
const WatchStorageV1beta1CSINodeOKCode int = 200

/*WatchStorageV1beta1CSINodeOK OK

swagger:response watchStorageV1beta1CSINodeOK
*/
type WatchStorageV1beta1CSINodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchStorageV1beta1CSINodeOK creates WatchStorageV1beta1CSINodeOK with default headers values
func NewWatchStorageV1beta1CSINodeOK() *WatchStorageV1beta1CSINodeOK {

	return &WatchStorageV1beta1CSINodeOK{}
}

// WithPayload adds the payload to the watch storage v1beta1 c s i node o k response
func (o *WatchStorageV1beta1CSINodeOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchStorageV1beta1CSINodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch storage v1beta1 c s i node o k response
func (o *WatchStorageV1beta1CSINodeOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchStorageV1beta1CSINodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchStorageV1beta1CSINodeUnauthorizedCode is the HTTP code returned for type WatchStorageV1beta1CSINodeUnauthorized
const WatchStorageV1beta1CSINodeUnauthorizedCode int = 401

/*WatchStorageV1beta1CSINodeUnauthorized Unauthorized

swagger:response watchStorageV1beta1CSINodeUnauthorized
*/
type WatchStorageV1beta1CSINodeUnauthorized struct {
}

// NewWatchStorageV1beta1CSINodeUnauthorized creates WatchStorageV1beta1CSINodeUnauthorized with default headers values
func NewWatchStorageV1beta1CSINodeUnauthorized() *WatchStorageV1beta1CSINodeUnauthorized {

	return &WatchStorageV1beta1CSINodeUnauthorized{}
}

// WriteResponse to the client
func (o *WatchStorageV1beta1CSINodeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
