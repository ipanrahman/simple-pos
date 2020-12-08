package category

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "categories"

type IRepository interface {
	FindAll(ctx context.Context) ([]Category, error)
	FindByID(ctx context.Context, id string) (Category, error)
}

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) IRepository {
	return Repository{db: db}
}

func (r Repository) FindAll(ctx context.Context) ([]Category, error) {
	var categories []Category
	cursor, err := r.db.Collection(collectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		if err := cursor.Decode(&categories); err != nil {
			return nil, err
		}
	}
	return categories, err
}

func (r Repository) FindByID(ctx context.Context, id string) (Category, error) {
	panic("Implement me")
}
