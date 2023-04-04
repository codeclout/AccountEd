package storage

import (
	"context"

	"github.com/codeclout/AccountEd/pkg/storage"
)

type HomeschoolActionPort interface {
	Login()
	Register(ctx context.Context, data *[]interface{}) ([]storage.MongoInsertOutput, error)
}
