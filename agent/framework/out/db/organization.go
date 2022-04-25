package db

import (
	"context"
	"errors"
	"time"

	model "github.com/codeclout/AccountEd/gateway/core/organization"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrorNotFound       = errors.New("not found")
	ErrorNotImplemented = errors.New("not implemented")
)

type Adapter struct {
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
}

func NewAdapter(ctx context.Context, client *mongo.Client, db string) *Adapter {
	collection := client.Database(db).Collection("organization")

	return &Adapter{
		client:     client,
		collection: collection,
		ctx:        ctx,
	}
}

// ActivateOrganization - find the record for which the id field matches the id
// and set isActive field to 1. Updates isPending & isMarkedForDeletion fields to 0.
// return ErrorNotFound if the record does not exist
func (a *Adapter) ActivateOrganization(id uuid.UUID) error {
	filter := bson.D{primitive.E{Key: "internalID", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "isActive", Value: 1},
		{Key: "isMarkedForDeletion", Value: 0},
		{Key: "isPending", Value: 0},
		{Key: "updatedAt", Value: time.Now()},
	}}}

	var record bson.M
	e := a.collection.FindOneAndUpdate(a.ctx, filter, update, options.FindOneAndUpdate()).Decode(&record)

	if e != nil {
		if errors.Is(e, mongo.ErrNoDocuments) {
			return ErrorNotFound
		}
		return e
	}

	return nil
}

func (a *Adapter) DeactivateOrganization(id uuid.UUID) error {
	return ErrorNotImplemented
}

func (a *Adapter) GetOrganization(id uuid.UUID) (model.Details, error) {
	ou := make([]model.OrganizationUnit, 0, 100)

	return model.Details{
		ID:    id,
		Name:  "",
		Units: ou,
	}, ErrorNotImplemented
}

func (a *Adapter) GetOrganizationBatch(ids []uuid.UUID) ([]model.Details, error) {
	s := make([]model.Details, 0, 100)

	return s, ErrorNotImplemented
}

func (a *Adapter) LogOrganizationHistoryEvent(ctx context.Context) error {
	return ErrorNotImplemented
}

func (a *Adapter) UpsertOrganizationUnit(unit model.Details) error {
	return ErrorNotImplemented
}
