// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchStorageV1beta1StorageClassListOKCode is the HTTP code returned for type WatchStorageV1beta1StorageClassListOK
const WatchStorageV1beta1StorageClassListOKCode int = 200

/*WatchStorageV1beta1StorageClassListOK OK

swagger:response watchStorageV1beta1StorageClassListOK
*/
type WatchStorageV1beta1StorageClassListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchStorageV1beta1StorageClassListOK creates WatchStorageV1beta1StorageClassListOK with default headers values
func NewWatchStorageV1beta1StorageClassListOK() *WatchStorageV1beta1StorageClassListOK {

	return &WatchStorageV1beta1StorageClassListOK{}
}

// WithPayload adds the payload to the watch storage v1beta1 storage class list o k response
func (o *WatchStorageV1beta1StorageClassListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchStorageV1beta1StorageClassListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch storage v1beta1 storage class list o k response
func (o *WatchStorageV1beta1StorageClassListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchStorageV1beta1StorageClassListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchStorageV1beta1StorageClassListUnauthorizedCode is the HTTP code returned for type WatchStorageV1beta1StorageClassListUnauthorized
const WatchStorageV1beta1StorageClassListUnauthorizedCode int = 401

/*WatchStorageV1beta1StorageClassListUnauthorized Unauthorized

swagger:response watchStorageV1beta1StorageClassListUnauthorized
*/
type WatchStorageV1beta1StorageClassListUnauthorized struct {
}

// NewWatchStorageV1beta1StorageClassListUnauthorized creates WatchStorageV1beta1StorageClassListUnauthorized with default headers values
func NewWatchStorageV1beta1StorageClassListUnauthorized() *WatchStorageV1beta1StorageClassListUnauthorized {

	return &WatchStorageV1beta1StorageClassListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchStorageV1beta1StorageClassListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
