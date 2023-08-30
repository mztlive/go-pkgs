package mongo

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	entity        *mongo.Client
	lock          sync.Once
	connectionURI string

	databaseMap sync.Map
)

// InitMongo 初始化 MongoDB 连接
//
// uri 参数是 MongoDB 连接字符串, 如果连接失败会触发 panic
func Initialize(uri string) {
	connectionURI = uri
}

// Connect 连接 MongoDB
func Connect() {
	lock.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
		if err != nil {
			panic(err)
		}

		entity = client
	})
}

// GetDatabase 获取指定名称的数据库实例
//
// 如果数据库连接未初始化，会触发 panic
// 如果数据库连接已经断开，会触发 panic
func GetDatabase(name string) *mongo.Database {

	if entity == nil {
		Connect()
	}

	err := entity.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	database, ok := databaseMap.Load(name)
	if !ok {
		database = entity.Database(name)
		databaseMap.Store(name, database)
	}

	return database.(*mongo.Database)
}
