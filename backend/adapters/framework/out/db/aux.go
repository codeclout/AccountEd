package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Adapter) getTimeStamp() primitive.DateTime {
	now := time.Now()
	t := time.Unix(0, now.UnixNano()).UTC()

	return primitive.NewDateTimeFromTime(t)
}
