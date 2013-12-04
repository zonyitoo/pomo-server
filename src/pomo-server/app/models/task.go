package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Task struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Type        int           `bson:"type"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	Create      *time.Time    `bson:"create"`
	Deadline    *time.Time    `bson:"deadline"`
	Estimate    int           `bson:"estimate"`
	Status      int           `bson:"status"`
}

const (
	TASK_STATUS_RUNNING = iota
	TASK_STATUS_COMPLETED
	TASK_STATUS_STOPPED
)

const (
	TASK_TYPE_NORMAL = iota
	TASK_TYPE_URGENT
)

const (
	TASK_COLLECTION_NAME = "Task"
)
