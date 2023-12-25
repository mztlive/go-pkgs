package mongo

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/mztlive/go-pkgs/snowflake"
	"github.com/spf13/cast"
)

type EntityInterface interface {
	SetBaseEntity(baseEntity BaseEntity)
	GetIdentity() string
	GetVersion() int64
	AddVersion()
	UpdateNow()
}

type BaseEntity struct {
	Identity  string `bson:"identity" json:"identity"`
	CreatedAt int64  `bson:"created_at" json:"created_at"`
	DeletedAt int64  `bson:"deleted_at" json:"deleted_at"`
	UpdatedAt int64  `bson:"updated_at" json:"updated_at"`
	Version   int64  `bson:"version" json:"version"`
}

// FromAny 从任意类型转换为当前类型
//
// 将任意类型的数据转换为当前类型，使用 github.com/jinzhu/copier 库实现
// 如果转换失败，将返回错误
//
// 如果传入的any的字段并不完全与当前类型的字段匹配，将会忽略多余的字段
func (a *BaseEntity) FromAny(any interface{}) error {
	return copier.Copy(a, any)
}

func (a *BaseEntity) GetIdentity() string {
	return a.Identity
}

func (a *BaseEntity) GetVersion() int64 {
	return a.Version
}

func (a *BaseEntity) AddVersion() {
	a.Version++
}

func (a *BaseEntity) UpdateNow() {
	a.UpdatedAt = time.Now().Unix()
}

func (a *BaseEntity) SetBaseEntity(baseEntity BaseEntity) {
	a.Identity = baseEntity.Identity
	a.CreatedAt = baseEntity.CreatedAt
	a.DeletedAt = baseEntity.DeletedAt
	a.UpdatedAt = baseEntity.UpdatedAt
	a.Version = baseEntity.Version
}

// NewBaseEntity 返回一个 BaseEntity 实例
func NewBaseEntity() BaseEntity {
	return BaseEntity{
		Identity:  cast.ToString(snowflake.GetID()),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		DeletedAt: 0,
		Version:   0,
	}
}
