package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Activity struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	TaskId bson.ObjectId `bson:"taskid"`
	Begin  *time.Time    `bson:"begin"`
	End    *time.Time    `bson:"end"`
	Status int           `bson:"status"`
}

const (
	ACTIVITY_STATUS_STOPPED = iota
	ACTIVITY_STATUS_RUNNING
	ACTIVITY_STATUS_COMPLETED
)

const (
	ACTIVITY_COLLECTION_NAME = "Activity"
)

func QueryActivitiesByTask(db *mgo.Database, task *Task) ([]Activity, error) {
	var activities []Activity
	err := db.C(ACTIVITY_COLLECTION_NAME).Find(bson.M{"taskid": task.Id}).All(&activities)
	return activities, err
}
