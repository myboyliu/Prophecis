// Code generated by go-swagger; DO NOT EDIT.

package auths

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CheckMLFlowResourceAccessHandlerFunc turns a function with the right signature into a check m l flow resource access handler
type CheckMLFlowResourceAccessHandlerFunc func(CheckMLFlowResourceAccessParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CheckMLFlowResourceAccessHandlerFunc) Handle(params CheckMLFlowResourceAccessParams) middleware.Responder {
	return fn(params)
}

// CheckMLFlowResourceAccessHandler interface for that can handle valid check m l flow resource access params
type CheckMLFlowResourceAccessHandler interface {
	Handle(CheckMLFlowResourceAccessParams) middleware.Responder
}

// NewCheckMLFlowResourceAccess creates a new http.Handler for the check m l flow resource access operation
func NewCheckMLFlowResourceAccess(ctx *middleware.Context, handler CheckMLFlowResourceAccessHandler) *CheckMLFlowResourceAccess {
	return &CheckMLFlowResourceAccess{Context: ctx, Handler: handler}
}

/*CheckMLFlowResourceAccess swagger:route GET /cc/v1/auth/access/mlflow/uri/{request_uri} Auths checkMLFlowResourceAccess

Use

Use in MLFlow Resource Check

*/
type CheckMLFlowResourceAccess struct {
	Context *middleware.Context
	Handler CheckMLFlowResourceAccessHandler
}

func (o *CheckMLFlowResourceAccess) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCheckMLFlowResourceAccessParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
