package api

import (
    "encoding/json"
    "net/http"

    "github.com/suedoh/go-commerce/types"
    "github.com/suedoh/go-commerce/store"

    "github.com/anthdm/weavebox"
)

type CreateProductRequest struct {
   SKU string `json:"sku"`
   Name string `json:"name"`
}

type ProductHandler struct {
   store store.ProductStorer 
}

func NewProductHandler(pStore store.ProductStorer) *ProductHandler {
    return &ProductHandler{
        store: pStore,
    } 
}


func (h *ProductHandler) HandlePostProduct(c *weavebox.Context) error {
	productReq := &CreateProductRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(productReq); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, productReq)
}

func (h *ProductHandler) HandleGetProduct(c *weavebox.Context)  error {
    return c.JSON(http.StatusOK, &types.Product{SKU: "SHOE-1111"}) 
}
