package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

// DB 数据库操作接口
// 该接口的实现类有sqlx.DB和sqlx.Tx
// 这个接口是与sqlx强耦合的，如果要替换sqlx，需要修改这个接口
// 定义这个接口的目的是为了在业务代码中使用这个接口，而不是直接使用sqlx.DB和sqlx.Tx，这样便可以传入db或者tx到repository中执行具体的sql
type DB interface {

	// SelectContext 执行查询操作
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// GetContext 执行查询操作，返回单条数据
	// 如果找不到数据会返回sql.ErrNoRows
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// ExecContext 执行非查询操作
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// PrepareNamedContext 准备一个命名查询
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)

	// Rebind 重新绑定查询参数
	Rebind(query string) string
}

type BaseEntity struct {
	Id        int64      `db:"id"`
	DeletedAt *time.Time `db:"deleted_at"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}
