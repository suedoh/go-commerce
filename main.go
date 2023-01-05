package main

import (
    "github.com/suedoh/go-commerce/api"

    "github.com/anthdm/weavebox"
)

func main()  {
   app := weavebox.New()

   productHandler := &api.ProductHandler{}

   app.Get("/product", productHandler.HandleGetProduct)
   app.Serve(3001)
}
