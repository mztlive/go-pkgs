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

	cursor, err := db.Collection(params.CollectionName).Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}

	response := []bson.M{}
	if err = cursor.All(context.TODO(), &response); err != nil {
		return err
	}

	list := response[0]["list"].(bson.A)
	count := response[0]["count"].(bson.A)

	if len(list) == 0 {
		return nil
	}

	for _, item := range list {
		ptr := new(T)
		bsonData, _ := bson.Marshal(item)
		bson.Unmarshal(bsonData, ptr)
		dest.Items = append(dest.Items, *ptr)
	}

	dest.Total = cast.ToInt64(count[0].(bson.M)["count"])
	return nil
}
