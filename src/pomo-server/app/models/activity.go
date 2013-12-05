package models

import (
	"errors"
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
	err := db.C(ACTIVITY_COLLECTION_NAME).FindId(task.Id).All(&activities)
	return activities, err
}

func ActivityStatusString(stat int) (string, error) {
	switch stat {
	case ACTIVITY_STATUS_STOPPED:
		return ACTIVITY_STATUS_STOPPED_STR, nil
	case ACTIVITY_STATUS_COMPLETED:
		return ACTIVITY_STATUS_COMPLETED_STR, nil
	case ACTIVITY_STATUS_RUNNING:
		return ACTIVITY_STATUS_RUNNING_STR, nil
	default:
		return "", errors.New("Invalid status")
	}
}

func ActivityStatusCode(stat string) (int, error) {
	switch stat {
	case ACTIVITY_STATUS_STOPPED_STR:
		return ACTIVITY_STATUS_STOPPED, nil
	case ACTIVITY_STATUS_COMPLETED_STR:
		return ACTIVITY_STATUS_COMPLETED, nil
	case ACTIVITY_STATUS_RUNNING_STR:
		return ACTIVITY_STATUS_RUNNING, nil
	default:
		return -1, errors.New("Invalid status")
	}
}
