// Code generated by go-swagger; DO NOT EDIT.

package authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAuthorizationV1beta1SelfSubjectAccessReviewHandlerFunc turns a function with the right signature into a create authorization v1beta1 self subject access review handler
type CreateAuthorizationV1beta1SelfSubjectAccessReviewHandlerFunc func(CreateAuthorizationV1beta1SelfSubjectAccessReviewParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAuthorizationV1beta1SelfSubjectAccessReviewHandlerFunc) Handle(params CreateAuthorizationV1beta1SelfSubjectAccessReviewParams) middleware.Responder {
	return fn(params)
}

// CreateAuthorizationV1beta1SelfSubjectAccessReviewHandler interface for that can handle valid create authorization v1beta1 self subject access review params
type CreateAuthorizationV1beta1SelfSubjectAccessReviewHandler interface {
	Handle(CreateAuthorizationV1beta1SelfSubjectAccessReviewParams) middleware.Responder
}

// NewCreateAuthorizationV1beta1SelfSubjectAccessReview creates a new http.Handler for the create authorization v1beta1 self subject access review operation
func NewCreateAuthorizationV1beta1SelfSubjectAccessReview(ctx *middleware.Context, handler CreateAuthorizationV1beta1SelfSubjectAccessReviewHandler) *CreateAuthorizationV1beta1SelfSubjectAccessReview {
	return &CreateAuthorizationV1beta1SelfSubjectAccessReview{Context: ctx, Handler: handler}
}

/*CreateAuthorizationV1beta1SelfSubjectAccessReview swagger:route POST /apis/authorization.k8s.io/v1beta1/selfsubjectaccessreviews authorization_v1beta1 createAuthorizationV1beta1SelfSubjectAccessReview

create a SelfSubjectAccessReview

*/
type CreateAuthorizationV1beta1SelfSubjectAccessReview struct {
	Context *middleware.Context
	Handler CreateAuthorizationV1beta1SelfSubjectAccessReviewHandler
}

func (o *CreateAuthorizationV1beta1SelfSubjectAccessReview) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAuthorizationV1beta1SelfSubjectAccessReviewParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
