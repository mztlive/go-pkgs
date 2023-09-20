package mongo

import (
	"context"

	"github.com/mztlive/go-pkgs/structure"
	"github.com/spf13/cast"
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

func QueryCollection[T any](ctx context.Context, params QueryParams, dest *structure.Collection[T], db *mongo.Database) error {

	offset := params.Paginator.Offset()
	limit := params.Paginator.Limit()
	matchStage := bson.M{"$match": params.Filter}
	sortStage := bson.M{"$sort": bson.M{"created_at": -1}}
	skipStage := bson.M{"$skip": offset}
	limitStage := bson.M{"$limit": limit}
	countStage := bson.M{"$count": "count"}

	pipeline := mongo.Pipeline{
		bson.D{
			{Key: "$facet", Value: bson.M{
				"list":  bson.A{matchStage, sortStage, skipStage, limitStage},
				"count": bson.A{matchStage, countStage},
			}},
		},
	}

	items, total, err := QueryItemsAndTotalFromPipeline[T](ctx, pipeline, params.CollectionName, db)
	if err != nil {
		return err
	}

	for _, item := range items {
		dest.Items = append(dest.Items, *item)
	}

	dest.Total = total
	return err
}

// QueryItemsAndTotalFromPipeline takes a mongo.Pipeline and a collection name
// and returns a slice of items and a total count of items
//
// Warning:
// The pipeline must have a $facet stage with a "list" and "count" stage
//
// Panics:
// If the pipeline does not have a $facet stage with a "list" and "count" stage
func QueryItemsAndTotalFromPipeline[T any](ctx context.Context, pipeline mongo.Pipeline, cName string, db *mongo.Database) (items []*T, total int64, err error) {
	var (
		response []bson.M
		cursor   *mongo.Cursor
		list     bson.A
		count    bson.A
	)

	if cursor, err = db.Collection(cName).Aggregate(ctx, pipeline); err != nil {
		return
	}

	if err = cursor.All(ctx, &response); err != nil {
		return
	}

	list = response[0]["list"].(bson.A)
	count = response[0]["count"].(bson.A)

	for _, item := range list {
		ptr := new(T)
		bsonData, _ := bson.Marshal(item)
		if err := bson.Unmarshal(bsonData, ptr); err != nil {
			return nil, 0, err
		}

		items = append(items, ptr)
	}

	total = cast.ToInt64(count[0].(bson.M)["count"])
	return
}
