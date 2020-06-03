// Code generated by go-swagger; DO NOT EDIT.

package authentication_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAuthenticationV1TokenReviewHandlerFunc turns a function with the right signature into a create authentication v1 token review handler
type CreateAuthenticationV1TokenReviewHandlerFunc func(CreateAuthenticationV1TokenReviewParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAuthenticationV1TokenReviewHandlerFunc) Handle(params CreateAuthenticationV1TokenReviewParams) middleware.Responder {
	return fn(params)
}

// CreateAuthenticationV1TokenReviewHandler interface for that can handle valid create authentication v1 token review params
type CreateAuthenticationV1TokenReviewHandler interface {
	Handle(CreateAuthenticationV1TokenReviewParams) middleware.Responder
}

// NewCreateAuthenticationV1TokenReview creates a new http.Handler for the create authentication v1 token review operation
func NewCreateAuthenticationV1TokenReview(ctx *middleware.Context, handler CreateAuthenticationV1TokenReviewHandler) *CreateAuthenticationV1TokenReview {
	return &CreateAuthenticationV1TokenReview{Context: ctx, Handler: handler}
}

/*CreateAuthenticationV1TokenReview swagger:route POST /apis/authentication.k8s.io/v1/tokenreviews authentication_v1 createAuthenticationV1TokenReview

create a TokenReview

*/
type CreateAuthenticationV1TokenReview struct {
	Context *middleware.Context
	Handler CreateAuthenticationV1TokenReviewHandler
}

func (o *CreateAuthenticationV1TokenReview) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAuthenticationV1TokenReviewParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
