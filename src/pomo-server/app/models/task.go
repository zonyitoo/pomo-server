package models

import "labix.org/v2/mgo/bson"
import "time"

type Task struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Type        int
	Record      []Activity
	Title       string
	Description string
	Create      *time.Time
	Deadline    *time.Time
	Estimate    int
	Complete    int
	Interrupt   int
	Status      int
}

const (
	TASK_STATUS_STOPPED = iota
	TASK_STATUS_COMPLETED
	TASK_STATUS_RUNNING
)

const (
	TASK_TYPE_NORMAL = iota
	TASK_TYPE_URGENT
)
