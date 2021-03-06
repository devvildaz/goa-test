// Code generated by goa v3.7.2, DO NOT EDIT.
//
// products endpoints
//
// Command:
// $ goa gen github.com/devvildaz/goa-code/pkg/design

package products

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "products" service endpoints.
type Endpoints struct {
	GetAllProducts goa.Endpoint
}

// NewEndpoints wraps the methods of the "products" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetAllProducts: NewGetAllProductsEndpoint(s),
	}
}

// Use applies the given middleware to all the "products" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetAllProducts = m(e.GetAllProducts)
}

// NewGetAllProductsEndpoint returns an endpoint function that calls the method
// "get_all_products" of service "products".
func NewGetAllProductsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetAllProducts(ctx)
	}
}
