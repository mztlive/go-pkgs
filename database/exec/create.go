package exec

import (
	"context"

	"github.com/mztlive/go-pkgs/database"
	"github.com/mztlive/go-pkgs/database/convert"
	"github.com/spf13/cast"
)

// CreateForNamed 将传入的entity插入到数据库中，并返回主键ID
//
// Entity的名字将转为表名，字段的db tag将转为列名
func CreateForNamed[T any](ctx context.Context, entity *T, db database.DB) (string, error) {
	insertStmt, err := convert.ConvertEntityToInsertNamedStmt(entity)
	if err != nil {
		return "", err
	}

	return namedInsert(ctx, db, insertStmt, entity)
}

// namedInsert 使用sqlx.NamedExec执行插入操作
// 返回插入数据的主键ID
func namedInsert(ctx context.Context, database database.DB, insertStmt string, data interface{}) (string, error) {
	stmt, err := database.PrepareNamedContext(ctx, insertStmt)
	if err != nil {
		return "", err
	}

	result, err := stmt.ExecContext(ctx, data)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return cast.ToString(id), nil
}
