// Code generated by go-swagger; DO NOT EDIT.

package authentication_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAuthenticationV1beta1TokenReviewHandlerFunc turns a function with the right signature into a create authentication v1beta1 token review handler
type CreateAuthenticationV1beta1TokenReviewHandlerFunc func(CreateAuthenticationV1beta1TokenReviewParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAuthenticationV1beta1TokenReviewHandlerFunc) Handle(params CreateAuthenticationV1beta1TokenReviewParams) middleware.Responder {
	return fn(params)
}

// CreateAuthenticationV1beta1TokenReviewHandler interface for that can handle valid create authentication v1beta1 token review params
type CreateAuthenticationV1beta1TokenReviewHandler interface {
	Handle(CreateAuthenticationV1beta1TokenReviewParams) middleware.Responder
}

// NewCreateAuthenticationV1beta1TokenReview creates a new http.Handler for the create authentication v1beta1 token review operation
func NewCreateAuthenticationV1beta1TokenReview(ctx *middleware.Context, handler CreateAuthenticationV1beta1TokenReviewHandler) *CreateAuthenticationV1beta1TokenReview {
	return &CreateAuthenticationV1beta1TokenReview{Context: ctx, Handler: handler}
}

/*CreateAuthenticationV1beta1TokenReview swagger:route POST /apis/authentication.k8s.io/v1beta1/tokenreviews authentication_v1beta1 createAuthenticationV1beta1TokenReview

create a TokenReview

*/
type CreateAuthenticationV1beta1TokenReview struct {
	Context *middleware.Context
	Handler CreateAuthenticationV1beta1TokenReviewHandler
}

func (o *CreateAuthenticationV1beta1TokenReview) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAuthenticationV1beta1TokenReviewParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
