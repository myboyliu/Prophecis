// Code generated by go-swagger; DO NOT EDIT.

package proxy_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetProxyUserByUserIDHandlerFunc turns a function with the right signature into a get proxy user by user Id handler
type GetProxyUserByUserIDHandlerFunc func(GetProxyUserByUserIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProxyUserByUserIDHandlerFunc) Handle(params GetProxyUserByUserIDParams) middleware.Responder {
	return fn(params)
}

// GetProxyUserByUserIDHandler interface for that can handle valid get proxy user by user Id params
type GetProxyUserByUserIDHandler interface {
	Handle(GetProxyUserByUserIDParams) middleware.Responder
}

// NewGetProxyUserByUserID creates a new http.Handler for the get proxy user by user Id operation
func NewGetProxyUserByUserID(ctx *middleware.Context, handler GetProxyUserByUserIDHandler) *GetProxyUserByUserID {
	return &GetProxyUserByUserID{Context: ctx, Handler: handler}
}

/*GetProxyUserByUserID swagger:route GET /cc/v1/proxyUser/{name}/user/{userName} proxyUser getProxyUserByUserId

get proxy user by name and userid.

Optional extended description in Markdown.

*/
type GetProxyUserByUserID struct {
	Context *middleware.Context
	Handler GetProxyUserByUserIDHandler
}

func (o *GetProxyUserByUserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetProxyUserByUserIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
