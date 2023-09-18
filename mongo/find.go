package mongo

import (
	"context"

	"github.com/mztlive/go-pkgs/reflect_utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetByIdentity 查找一个实体
//
// collection_name 是实体的名称, entity是实体的指针
func GetByIdentity(ctx context.Context, identity string, entity any, db *mongo.Database) error {
	collectionName := reflect_utils.GetSnakeNameFromStruct(entity)
	return db.Collection(collectionName).FindOne(ctx, bson.M{
		"identity": identity,
	}).Decode(entity)

}

// Get 查找一个实体
//
// collection_name 是实体的名称, entity是实体的指针
func Get(ctx context.Context, filters bson.M, entity any, db *mongo.Database) error {
	collectionName := reflect_utils.GetSnakeNameFromStruct(entity)
	return db.Collection(collectionName).FindOne(ctx, filters).Decode(entity)
}
