// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchStorageV1beta1StorageClassOKCode is the HTTP code returned for type WatchStorageV1beta1StorageClassOK
const WatchStorageV1beta1StorageClassOKCode int = 200

/*WatchStorageV1beta1StorageClassOK OK

swagger:response watchStorageV1beta1StorageClassOK
*/
type WatchStorageV1beta1StorageClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchStorageV1beta1StorageClassOK creates WatchStorageV1beta1StorageClassOK with default headers values
func NewWatchStorageV1beta1StorageClassOK() *WatchStorageV1beta1StorageClassOK {

	return &WatchStorageV1beta1StorageClassOK{}
}

// WithPayload adds the payload to the watch storage v1beta1 storage class o k response
func (o *WatchStorageV1beta1StorageClassOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchStorageV1beta1StorageClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch storage v1beta1 storage class o k response
func (o *WatchStorageV1beta1StorageClassOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchStorageV1beta1StorageClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchStorageV1beta1StorageClassUnauthorizedCode is the HTTP code returned for type WatchStorageV1beta1StorageClassUnauthorized
const WatchStorageV1beta1StorageClassUnauthorizedCode int = 401

/*WatchStorageV1beta1StorageClassUnauthorized Unauthorized

swagger:response watchStorageV1beta1StorageClassUnauthorized
*/
type WatchStorageV1beta1StorageClassUnauthorized struct {
}

// NewWatchStorageV1beta1StorageClassUnauthorized creates WatchStorageV1beta1StorageClassUnauthorized with default headers values
func NewWatchStorageV1beta1StorageClassUnauthorized() *WatchStorageV1beta1StorageClassUnauthorized {

	return &WatchStorageV1beta1StorageClassUnauthorized{}
}

// WriteResponse to the client
func (o *WatchStorageV1beta1StorageClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
