package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Adapter) getTimeStamp() primitive.DateTime {
	t := time.UnixMicro(time.Now().Unix()).UTC()

	return primitive.NewDateTimeFromTime(t)
}
