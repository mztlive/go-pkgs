package mongo

type BaseEntity struct {
	Identity  string `bson:"identity"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
	DeletedAt int64  `bson:"deleted_at"`
}
