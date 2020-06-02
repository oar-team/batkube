// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// WatchStorageV1StorageClassOKCode is the HTTP code returned for type WatchStorageV1StorageClassOK
const WatchStorageV1StorageClassOKCode int = 200

/*WatchStorageV1StorageClassOK OK

swagger:response watchStorageV1StorageClassOK
*/
type WatchStorageV1StorageClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchStorageV1StorageClassOK creates WatchStorageV1StorageClassOK with default headers values
func NewWatchStorageV1StorageClassOK() *WatchStorageV1StorageClassOK {

	return &WatchStorageV1StorageClassOK{}
}

// WithPayload adds the payload to the watch storage v1 storage class o k response
func (o *WatchStorageV1StorageClassOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchStorageV1StorageClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch storage v1 storage class o k response
func (o *WatchStorageV1StorageClassOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchStorageV1StorageClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchStorageV1StorageClassUnauthorizedCode is the HTTP code returned for type WatchStorageV1StorageClassUnauthorized
const WatchStorageV1StorageClassUnauthorizedCode int = 401

/*WatchStorageV1StorageClassUnauthorized Unauthorized

swagger:response watchStorageV1StorageClassUnauthorized
*/
type WatchStorageV1StorageClassUnauthorized struct {
}

// NewWatchStorageV1StorageClassUnauthorized creates WatchStorageV1StorageClassUnauthorized with default headers values
func NewWatchStorageV1StorageClassUnauthorized() *WatchStorageV1StorageClassUnauthorized {

	return &WatchStorageV1StorageClassUnauthorized{}
}

// WriteResponse to the client
func (o *WatchStorageV1StorageClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
