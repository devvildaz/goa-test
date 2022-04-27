package productsapi

import (
	"context"
	"log"

	products "github.com/devvildaz/goa-code/gen/products"
)

// products service example implementation.
// The example methods log the requests and return zero values.
type productssrvc struct {
	logger *log.Logger
}

// NewProducts returns the products service implementation.
func NewProducts(logger *log.Logger) products.Service {
	return &productssrvc{logger}
}

// GetAllProducts implements get_all_products.
func (s *productssrvc) GetAllProducts(ctx context.Context) (res int, err error) {
	s.logger.Print("products.get_all_products")
	return 10, nil
}
