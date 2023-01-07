package store

import (
    "context"

    "github.com/suedoh/go-commerce/types"
)

type ProductStorer interface {
   Insert(context.Context, *types.Product) error 
   GetByID(context.Context, string) (*types.Product, error)
}
