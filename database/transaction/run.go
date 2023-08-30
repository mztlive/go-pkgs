package transaction

import (
	"context"
	"fmt"

	"github.com/mztlive/go-pkgs/database"
)

// Run 开启一个事务并且自动提交或者回滚
// 如果txFunc返回error或者panic，事务会回滚
func Run(ctx context.Context, txFunc func(database.DB) error) error {
	db := database.GetDB()
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction failed. reason: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			// panic(p) // re-throw panic
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()

	err = txFunc(tx)
	return err
}
