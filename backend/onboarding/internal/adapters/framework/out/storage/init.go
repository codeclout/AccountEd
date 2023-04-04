package storage

import "github.com/codeclout/AccountEd/pkg/storage/adapters/framework/out"

type Adapter struct {
	mongoActions      *out.MongoActions
	log               func(level, msg string)
	ResponseItemLimit int32
}

func NewAdapter(m *out.MongoActions, logger func(level, msg string)) *Adapter {
	return &Adapter{
		mongoActions: m,
		log:          logger,
	}
}
