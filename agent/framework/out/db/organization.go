package db

import (
	"context"
	"log"
	"time"

	model "github.com/codeclout/AccountEd/gateway/core/organization"
	repo "github.com/codeclout/AccountEd/gateway/framework/out/db"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Adapter struct {
	client  *mongo.Client
	ctx     context.Context
	db      string
	repo    repo.OrganizationRepository
	timeout time.Duration
}

func (a *Adapter) ActivateOrganization(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (a *Adapter) DeactivateOrganization(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (a *Adapter) GetOrganization(ctx context.Context, id uuid.UUID) (model.Details, error) {
	ou := make([]model.OrganizationUnit, 0, 100)

	return model.Details{
		ID:    id,
		Name:  "",
		Units: ou,
	}, nil
}

func (a *Adapter) GetOrganizationBatch(ctx context.Context, ids []uuid.UUID) ([]model.Details, error) {
	s := make([]model.Details, 0, 100)

	return s, nil
}

func (a *Adapter) LogOrganizationHistoryEvent(ctx context.Context, event model.OrganizationEvent) error {
	return nil
}

func (a *Adapter) UpsertOrganizationUnit(ctx context.Context, unit model.Details) error {
	return nil
}

func NewAdapter(db, uri string, repo repo.OrganizationRepository, timeout int) (*Adapter, error) {
	t := time.Duration(timeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	client, e := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if e != nil {
		log.Fatalf("db connection failed: %v", e)
	}

	e = client.Ping(ctx, readpref.Primary())
	if e != nil {
		log.Fatalf("db ping failed: %v", e)
	}

	a := Adapter{
		client:  client,
		ctx:     ctx,
		db:      db,
		repo:    repo,
		timeout: time.Duration(t) * time.Second,
	}

	defer a.CloseConnection()

	return &a, nil
}

func (a *Adapter) CloseConnection() {
	if e := a.client.Disconnect(a.ctx); e != nil {
		panic(e)
	}
}
