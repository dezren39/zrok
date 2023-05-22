// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/openziti/zrok/rest_model_zrok"
)

// GetShareMetricsHandlerFunc turns a function with the right signature into a get share metrics handler
type GetShareMetricsHandlerFunc func(GetShareMetricsParams, *rest_model_zrok.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetShareMetricsHandlerFunc) Handle(params GetShareMetricsParams, principal *rest_model_zrok.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetShareMetricsHandler interface for that can handle valid get share metrics params
type GetShareMetricsHandler interface {
	Handle(GetShareMetricsParams, *rest_model_zrok.Principal) middleware.Responder
}

// NewGetShareMetrics creates a new http.Handler for the get share metrics operation
func NewGetShareMetrics(ctx *middleware.Context, handler GetShareMetricsHandler) *GetShareMetrics {
	return &GetShareMetrics{Context: ctx, Handler: handler}
}

/*
	GetShareMetrics swagger:route GET /metrics/share/{shrToken} metrics getShareMetrics

GetShareMetrics get share metrics API
*/
type GetShareMetrics struct {
	Context *middleware.Context
	Handler GetShareMetricsHandler
}

func (o *GetShareMetrics) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetShareMetricsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *rest_model_zrok.Principal
	if uprinc != nil {
		principal = uprinc.(*rest_model_zrok.Principal) // this is really a rest_model_zrok.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
