package query

import (
	"context"

	"github.com/mztlive/project-template/pkg/structure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QueryParams struct {
	Filter         bson.M
	Paginator      structure.Paginator
	CollectionName string
}

// QuerySlice 根据filter查询数据
// dests是一个slice，用于接收查询结果
//
// 查询结果会根据paginator进行分页并且使用created_at进行降序
func QuerySlice[T any](ctx context.Context, params QueryParams, dests *[]T, db *mongo.Database) error {
	offset := params.Paginator.Offset()
	limit := params.Paginator.Limit()

	option := &options.FindOptions{
		Sort: bson.M{"created_at": -1},
	}

	if offset != 0 && limit != 0 {
		option = option.SetSkip(offset)
		option = option.SetLimit(limit)
	}

	cursor, err := db.Collection(params.CollectionName).Find(ctx, params.Filter, option)

	if err != nil {
		return err
	}

	return cursor.All(ctx, dests)
}
