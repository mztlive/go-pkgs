package mongo

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/mztlive/go-pkgs/reflect_utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCasLock = errors.New("cas lock")
)

func retryWait(i int) {
	waitSec := (i + 1) * rand.Intn(500)
	time.Sleep(time.Millisecond * time.Duration(waitSec))
}

// UpdateDocumentWithNotTimestamp 更新文档，不更新时间戳
func UpdateDocumentWithNotTimestamp(ctx context.Context, entity EntityInterface, db *mongo.Database) error {

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
		return ErrCasLock
	}

	return err
}

func UpdateDocument(ctx context.Context, entity EntityInterface, db *mongo.Database) error {

	collectionName := reflect_utils.GetSnakeNameFromStruct(entity)
	filter := bson.M{
		"identity": entity.GetIdentity(),
		"version":  entity.GetVersion(),
	}

	entity.AddVersion()
	entity.UpdateNow()
	upRes, err := db.Collection(collectionName).UpdateOne(ctx, filter, bson.M{
		"$set": entity,
	})

	if upRes.ModifiedCount == 0 {
		return ErrCasLock
	}

	return err
}

// UpdateDocumentWithCasRetry 乐观锁更新
// 重试次数为3次
func UpdateDocumentWithCasRetry(ctx context.Context, entity EntityInterface, db *mongo.Database) error {
	for i := 0; i < 3; i++ {
		err := UpdateDocument(ctx, entity, db)

		if err == nil {
			return nil
		}

		if err != ErrCasLock {
			return err
		}

		retryWait(i)
	}

	return ErrCasLock
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
