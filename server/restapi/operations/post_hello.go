// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostHelloHandlerFunc turns a function with the right signature into a post hello handler
type PostHelloHandlerFunc func(PostHelloParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostHelloHandlerFunc) Handle(params PostHelloParams) middleware.Responder {
	return fn(params)
}

// PostHelloHandler interface for that can handle valid post hello params
type PostHelloHandler interface {
	Handle(PostHelloParams) middleware.Responder
}

// NewPostHello creates a new http.Handler for the post hello operation
func NewPostHello(ctx *middleware.Context, handler PostHelloHandler) *PostHello {
	return &PostHello{Context: ctx, Handler: handler}
}

/* PostHello swagger:route POST /hello postHello

say hello to someone

*/
type PostHello struct {
	Context *middleware.Context
	Handler PostHelloHandler
}

func (o *PostHello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostHelloParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostHelloOKBody post hello o k body
//
// swagger:model PostHelloOKBody
type PostHelloOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this post hello o k body
func (o *PostHelloOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post hello o k body based on context it is used
func (o *PostHelloOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostHelloOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostHelloOKBody) UnmarshalBinary(b []byte) error {
	var res PostHelloOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}