// Code generated by go-swagger; DO NOT EDIT.

package certificates_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCertificatesV1beta1CertificateSigningRequestListHandlerFunc turns a function with the right signature into a watch certificates v1beta1 certificate signing request list handler
type WatchCertificatesV1beta1CertificateSigningRequestListHandlerFunc func(WatchCertificatesV1beta1CertificateSigningRequestListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCertificatesV1beta1CertificateSigningRequestListHandlerFunc) Handle(params WatchCertificatesV1beta1CertificateSigningRequestListParams) middleware.Responder {
	return fn(params)
}

// WatchCertificatesV1beta1CertificateSigningRequestListHandler interface for that can handle valid watch certificates v1beta1 certificate signing request list params
type WatchCertificatesV1beta1CertificateSigningRequestListHandler interface {
	Handle(WatchCertificatesV1beta1CertificateSigningRequestListParams) middleware.Responder
}

// NewWatchCertificatesV1beta1CertificateSigningRequestList creates a new http.Handler for the watch certificates v1beta1 certificate signing request list operation
func NewWatchCertificatesV1beta1CertificateSigningRequestList(ctx *middleware.Context, handler WatchCertificatesV1beta1CertificateSigningRequestListHandler) *WatchCertificatesV1beta1CertificateSigningRequestList {
	return &WatchCertificatesV1beta1CertificateSigningRequestList{Context: ctx, Handler: handler}
}

/*WatchCertificatesV1beta1CertificateSigningRequestList swagger:route GET /apis/certificates.k8s.io/v1beta1/watch/certificatesigningrequests certificates_v1beta1 watchCertificatesV1beta1CertificateSigningRequestList

watch individual changes to a list of CertificateSigningRequest. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCertificatesV1beta1CertificateSigningRequestList struct {
	Context *middleware.Context
	Handler WatchCertificatesV1beta1CertificateSigningRequestListHandler
}

func (o *WatchCertificatesV1beta1CertificateSigningRequestList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCertificatesV1beta1CertificateSigningRequestListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
