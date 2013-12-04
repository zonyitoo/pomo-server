package models

import (
	"errors"
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

func TaskStatusString(stat int) (string, error) {
	switch stat {
	case TASK_STATUS_COMPLETED:
		return TASK_STATUS_COMPLETED_STR, nil
	case TASK_STATUS_RUNNING:
		return TASK_STATUS_RUNNING_STR, nil
	case TASK_STATUS_STOPPED:
		return TASK_STATUS_STOPPED_STR, nil
	default:
		return "", errors.New("Invalid status")
	}
}

func TaskStatusCode(stat string) (int, error) {
	switch stat {
	case TASK_STATUS_COMPLETED_STR:
		return TASK_STATUS_COMPLETED, nil
	case TASK_STATUS_RUNNING_STR:
		return TASK_STATUS_RUNNING, nil
	case TASK_STATUS_STOPPED_STR:
		return TASK_STATUS_STOPPED, nil
	default:
		return -1, errors.New("Invalid status")
	}
}

func TaskTypeString(t int) (string, error) {
	switch t {
	case TASK_TYPE_NORMAL:
		return TASK_TYPE_NORMAL_STR, nil
	case TASK_TYPE_URGENT:
		return TASK_TYPE_URGENT_STR, nil

	default:
		return "", errors.New("Invalid type")
	}
}

func TaskTypeCode(t string) (int, error) {
	switch t {
	case TASK_TYPE_NORMAL_STR:
		return TASK_TYPE_NORMAL, nil
	case TASK_TYPE_URGENT_STR:
		return TASK_TYPE_URGENT, nil
	default:
		return -1, errors.New("Invalid type")
	}
}
