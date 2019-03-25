// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTransactionReceivedHandlerFunc turns a function with the right signature into a get transaction received handler
type GetTransactionReceivedHandlerFunc func(GetTransactionReceivedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTransactionReceivedHandlerFunc) Handle(params GetTransactionReceivedParams) middleware.Responder {
	return fn(params)
}

// GetTransactionReceivedHandler interface for that can handle valid get transaction received params
type GetTransactionReceivedHandler interface {
	Handle(GetTransactionReceivedParams) middleware.Responder
}

// NewGetTransactionReceived creates a new http.Handler for the get transaction received operation
func NewGetTransactionReceived(ctx *middleware.Context, handler GetTransactionReceivedHandler) *GetTransactionReceived {
	return &GetTransactionReceived{Context: ctx, Handler: handler}
}

/*GetTransactionReceived swagger:route GET /transactions/account/{receiver}/received/{transaction} transactions getTransactionReceived

Get a specific transaction received by a given account

Get a specific transaction received by a given account

*/
type GetTransactionReceived struct {
	Context *middleware.Context
	Handler GetTransactionReceivedHandler
}

func (o *GetTransactionReceived) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTransactionReceivedParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
