package main

import (
    "github.com/suedoh/go-commerce/api"

    "github.com/anthdm/weavebox"
)

func main()  {
   app := weavebox.New()
   adminRoute := app.Box("/admin")

   productHandler := &api.ProductHandler{}


   adminRoute.Get("/product", productHandler.HandleGetProduct)
   adminRoute.Post("/product", productHandler.HandlePostProduct)

   app.Serve(3001)
}
