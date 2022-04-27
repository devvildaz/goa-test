package design

import (
  .  "goa.design/goa/v3/dsl"
)

var _ = API("products", func(){
  Title("Product Service")
  Description("HTTP Product Service")
  Server("products",func() { 
    Host("localhost", func() { URI("http://localhost:8088") })
  })
})

var _ = Service("products", func(){
  Description("the product service")
  Method("get_all_products", func() {
    Result(Int)
    HTTP(func(){
      GET("/products")
      Response(StatusOK)
    })
  })
})

var StoredProduct = ResultType("application", func() {
  
}) 
