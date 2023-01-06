package main

import (
    "context"

    "github.com/suedoh/go-commerce/api"
    "github.com/suedoh/go-commerce/store"

    "github.com/anthdm/weavebox"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main()  {
    app := weavebox.New()
    adminRoute := app.Box("/admin")



    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        panic(err)
    }


    productStore := store.NewMongoProductStore(client)
    productHandler := api.NewProductHandler(productStore)


    adminRoute.Get("/product", productHandler.HandleGetProduct)
    adminRoute.Post("/product", productHandler.HandlePostProduct)

    app.Serve(3001)
}
