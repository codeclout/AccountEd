package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Adapter) getTimeStamp() primitive.DateTime {
	t := time.Unix(time.Now().Unix(), 0).UTC()

	return primitive.NewDateTimeFromTime(t)
}
