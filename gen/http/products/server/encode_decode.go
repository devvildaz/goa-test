// Code generated by goa v3.7.2, DO NOT EDIT.
//
// products HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/devvildaz/goa-code/pkg/design

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// EncodeGetAllProductsResponse returns an encoder for responses returned by
// the products get_all_products endpoint.
func EncodeGetAllProductsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(int)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}