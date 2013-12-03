package models

import "time"
import "labix.org/v2/mgo/bson"

type Activity struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Begin  *time.Time
	End    *time.Time
	Status int
}

const (
	ACTIVITY_STATUS_STOPPED = iota
	ACTIVITY_STATUS_RUNNING
	ACTIVITY_STATUS_COMPLETED
)
