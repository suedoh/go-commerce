package store

import (
    "github.com/suedoh/go-commerce/types"
    "go.mongo.org/mongo-driver/mongo"
)

type MongoProductStore struct {
    client *mongo.Client
}

func NewMongoProductStore(c *mongo.Client) *MongoProductStore {
    return &MongoProductStore {
        client: c,
    }
}

func (s *MongoProductStore) Insert(p *types.Product) error {
   return nil
}

func (s *MongoProductStore) GetByID(id string) (*types.Product, error)  {
   return nil, nil 
}
