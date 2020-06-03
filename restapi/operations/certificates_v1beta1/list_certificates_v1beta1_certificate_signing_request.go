// Code generated by go-swagger; DO NOT EDIT.

package certificates_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListCertificatesV1beta1CertificateSigningRequestHandlerFunc turns a function with the right signature into a list certificates v1beta1 certificate signing request handler
type ListCertificatesV1beta1CertificateSigningRequestHandlerFunc func(ListCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListCertificatesV1beta1CertificateSigningRequestHandlerFunc) Handle(params ListCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder {
	return fn(params)
}

// ListCertificatesV1beta1CertificateSigningRequestHandler interface for that can handle valid list certificates v1beta1 certificate signing request params
type ListCertificatesV1beta1CertificateSigningRequestHandler interface {
	Handle(ListCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder
}

// NewListCertificatesV1beta1CertificateSigningRequest creates a new http.Handler for the list certificates v1beta1 certificate signing request operation
func NewListCertificatesV1beta1CertificateSigningRequest(ctx *middleware.Context, handler ListCertificatesV1beta1CertificateSigningRequestHandler) *ListCertificatesV1beta1CertificateSigningRequest {
	return &ListCertificatesV1beta1CertificateSigningRequest{Context: ctx, Handler: handler}
}

/*ListCertificatesV1beta1CertificateSigningRequest swagger:route GET /apis/certificates.k8s.io/v1beta1/certificatesigningrequests certificates_v1beta1 listCertificatesV1beta1CertificateSigningRequest

list or watch objects of kind CertificateSigningRequest

*/
type ListCertificatesV1beta1CertificateSigningRequest struct {
	Context *middleware.Context
	Handler ListCertificatesV1beta1CertificateSigningRequestHandler
}

func (o *ListCertificatesV1beta1CertificateSigningRequest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListCertificatesV1beta1CertificateSigningRequestParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
