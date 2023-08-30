package exec

import (
	"context"
	"time"

	"github.com/mztlive/go-pkgs/reflect_utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateExecutor struct{}

func (a *UpdateExecutor) UpdateDocument(ctx context.Context, entity EntityInterface, db *mongo.Database) error {

	collectionName := reflect_utils.GetSnakeNameFromStruct(entity)
	filter := bson.M{
		"identity": entity.GetIdentity(),
	}

	_, err := db.Collection(collectionName).UpdateOne(ctx, filter, bson.M{
		"$set": entity,
	})

	return err
}

// SoftDelete 软删除 (将deleted_at字段设置为当前时间)
//
// filter是查询条件
func (a *UpdateExecutor) SoftDeleteMany(ctx context.Context, filter bson.M, collection string, db *mongo.Database) error {
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now().Unix(),
		},
	}
	_, err := db.Collection(collection).UpdateMany(ctx, filter, update)
	return err
}
