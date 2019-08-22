// Code generated by go-swagger; DO NOT EDIT.

package blacklist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddToBlackListHandlerFunc turns a function with the right signature into a add to black list handler
type AddToBlackListHandlerFunc func(AddToBlackListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddToBlackListHandlerFunc) Handle(params AddToBlackListParams) middleware.Responder {
	return fn(params)
}

// AddToBlackListHandler interface for that can handle valid add to black list params
type AddToBlackListHandler interface {
	Handle(AddToBlackListParams) middleware.Responder
}

// NewAddToBlackList creates a new http.Handler for the add to black list operation
func NewAddToBlackList(ctx *middleware.Context, handler AddToBlackListHandler) *AddToBlackList {
	return &AddToBlackList{Context: ctx, Handler: handler}
}

/*AddToBlackList swagger:route POST /blacklist blacklist addToBlackList

Add a new blacklist IP to the store

*/
type AddToBlackList struct {
	Context *middleware.Context
	Handler AddToBlackListHandler
}

func (o *AddToBlackList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddToBlackListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}