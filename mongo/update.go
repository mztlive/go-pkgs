package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/mztlive/go-pkgs/reflect_utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateDocument(ctx context.Context, entity EntityInterface, db *mongo.Database) error {

	collectionName := reflect_utils.GetSnakeNameFromStruct(entity)
	filter := bson.M{
		"identity": entity.GetIdentity(),
		"version":  entity.GetVersion(),
	}
	entity.AddVersion()
	upRes, err := db.Collection(collectionName).UpdateOne(ctx, filter, bson.M{
		"$set": entity,
	})
	if upRes.ModifiedCount == 0 {
		return errors.New("update failed. maybe the document is not exist or the version is not match")
	}
	return err
}

// SoftDelete 软删除 (将deleted_at字段设置为当前时间)
//
// filter是查询条件
func SoftDeleteMany(ctx context.Context, filter bson.M, collection string, db *mongo.Database) error {
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now().Unix(),
		},
	}
	_, err := db.Collection(collection).UpdateMany(ctx, filter, update)
	return err
}
