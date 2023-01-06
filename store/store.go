package store

type ProductStorer interface {
   Insert(*types.Product) error 
   GetByID(string) (*types.Product, error)
}
