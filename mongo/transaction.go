package mongo

import (
	"context"
	"fmt"

	g_mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// TransactionHandler is a function that will be executed in a transaction
//
// Must use the database object passed in as a parameter to execute the transaction
type TransactionHandler func(ctx g_mongo.SessionContext, database *g_mongo.Database) (any, error)

// StartTransaction returns the result of the handler function.
// If the handler returns an error, the transaction will be aborted.
// If the handler returns nil, the transaction will be committed.
//
// # example:
//
//	_, err := StartTransaction(ctx, db, func(ctx mongo.SessionContext, session mongo.Session) (any, error) {
//		// do something
//		return nil, nil
//	})
func StartTransaction(ctx context.Context, db *g_mongo.Database, handler TransactionHandler) (any, error) {
	var (
		result  any
		session g_mongo.Session
		err     error
	)

	txo := options.Transaction().SetReadConcern(readconcern.Majority()).SetWriteConcern(writeconcern.Majority())

	if session, err = db.Client().StartSession(); err != nil {
		return nil, fmt.Errorf("start session error: %w", err)
	}
	defer session.EndSession(ctx)

	result, err = session.WithTransaction(ctx, func(sCtx g_mongo.SessionContext) (interface{}, error) {
		return handler(sCtx, db)
	}, txo)

	return result, err
}
