package mongo

import (
	"context"

	"github.com/mztlive/go-pkgs/reflect_utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// Insert 插入一个实体
//
// collection_name 是实体的名称
func Insert(ctx context.Context, entity any, db *mongo.Database) (*mongo.InsertOneResult, error) {
	collectionName := reflect_utils.GetSnakeNameFromStruct(entity)
	return db.Collection(collectionName).InsertOne(ctx, entity)
}

// InsertMany 插入多个实体
//
// collection_name 是实体的名称
func InsertMany(ctx context.Context, entities []any, db *mongo.Database) (*mongo.InsertManyResult, error) {
	collectionName := reflect_utils.GetSnakeNameFromStruct(entities[0])
	return db.Collection(collectionName).InsertMany(ctx, entities)
}
