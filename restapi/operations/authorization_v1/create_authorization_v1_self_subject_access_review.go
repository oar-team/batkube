// Code generated by go-swagger; DO NOT EDIT.

package authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAuthorizationV1SelfSubjectAccessReviewHandlerFunc turns a function with the right signature into a create authorization v1 self subject access review handler
type CreateAuthorizationV1SelfSubjectAccessReviewHandlerFunc func(CreateAuthorizationV1SelfSubjectAccessReviewParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAuthorizationV1SelfSubjectAccessReviewHandlerFunc) Handle(params CreateAuthorizationV1SelfSubjectAccessReviewParams) middleware.Responder {
	return fn(params)
}

// CreateAuthorizationV1SelfSubjectAccessReviewHandler interface for that can handle valid create authorization v1 self subject access review params
type CreateAuthorizationV1SelfSubjectAccessReviewHandler interface {
	Handle(CreateAuthorizationV1SelfSubjectAccessReviewParams) middleware.Responder
}

// NewCreateAuthorizationV1SelfSubjectAccessReview creates a new http.Handler for the create authorization v1 self subject access review operation
func NewCreateAuthorizationV1SelfSubjectAccessReview(ctx *middleware.Context, handler CreateAuthorizationV1SelfSubjectAccessReviewHandler) *CreateAuthorizationV1SelfSubjectAccessReview {
	return &CreateAuthorizationV1SelfSubjectAccessReview{Context: ctx, Handler: handler}
}

/*CreateAuthorizationV1SelfSubjectAccessReview swagger:route POST /apis/authorization.k8s.io/v1/selfsubjectaccessreviews authorization_v1 createAuthorizationV1SelfSubjectAccessReview

create a SelfSubjectAccessReview

*/
type CreateAuthorizationV1SelfSubjectAccessReview struct {
	Context *middleware.Context
	Handler CreateAuthorizationV1SelfSubjectAccessReviewHandler
}

func (o *CreateAuthorizationV1SelfSubjectAccessReview) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAuthorizationV1SelfSubjectAccessReviewParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
