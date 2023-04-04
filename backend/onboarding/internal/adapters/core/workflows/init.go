package workflows

import (
	"github.com/codeclout/AccountEd/onboarding/internal/ports/framework/out/storage"
	"github.com/codeclout/AccountEd/pkg/storage/adapters/framework/out"
)

type logger func(level, msg string)
type persist storage.HomeschoolActionPort

type Adapter struct {
	log          logger
	mongoActions out.MongoActions
	storageOut   persist
}

func NewAdapter(l logger, s persist, a *out.MongoActions) *Adapter {
	return &Adapter{
		log:          l,
		mongoActions: *a,
		storageOut:   s,
	}
}
