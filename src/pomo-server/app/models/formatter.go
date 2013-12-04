package models

import (
	"labix.org/v2/mgo"
)

type ActivityObject struct {
	Id     string `json:"id"`
	Begin  string `json:"begin"`
	End    string `json:"end"`
	Status string `json:"status"`
}

const (
	ACTIVITY_STATUS_COMPLETED_STR = "completed"
	ACTIVITY_STATUS_RUNNING_STR   = "running"
	ACTIVITY_STATUS_STOPPED_STR   = "stopped"
)

type TaskObject struct {
	Id          string           `json:"id"`
	Type        string           `json:"type"`
	Record      []ActivityObject `json:"record"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Create      string           `json:"create"`
	Deadline    string           `json:"deadline"`
	Estimate    int              `json:"estimate"`
	Complete    int              `json:"complete"`
	Interrupt   int              `json:"interrupt"`
	Status      string           `json:"status"`
}

const (
	TASK_STATUS_COMPLETED_STR = "completed"
	TASK_STATUS_RUNNING_STR   = "running"
	TASK_STATUS_STOPPED_STR   = "stopped"
)

const (
	TASK_TYPE_NORMAL_STR = "normal"
	TASK_TYPE_URGENT_STR = "urgent"
)

type TaskObjectList struct {
	Type   string       `json:"type"`
	Date   string       `json:"date"`
	Status string       `json:"status"`
	Tasks  []TaskObject `json:"tasks"`
}

func (a *Activity) ToActivityObject() ActivityObject {
	var result ActivityObject

	result.Id = a.Id.Hex()
	if a.Begin != nil {
		begin, errb := a.Begin.MarshalText()
		if errb != nil {
			result.Begin = ""
		} else {
			result.Begin = string(begin)
		}
	} else {
		result.Begin = ""
	}

	if a.End != nil {
		end, erre := a.End.MarshalText()
		if erre != nil {
			result.End = ""
		} else {
			result.End = string(end)
		}
	} else {
		result.End = ""
	}
	switch a.Status {
	case ACTIVITY_STATUS_STOPPED:
		result.Status = ACTIVITY_STATUS_STOPPED_STR
	case ACTIVITY_STATUS_RUNNING:
		result.Status = ACTIVITY_STATUS_RUNNING_STR
	case ACTIVITY_STATUS_COMPLETED:
		result.Status = ACTIVITY_STATUS_COMPLETED_STR
	default:
		result.Status = ACTIVITY_STATUS_STOPPED_STR
	}

	return result
}

func (t *Task) ToTaskObject(db *mgo.Database) TaskObject {
	var result TaskObject

	result.Id = t.Id.Hex()
	switch t.Type {
	case TASK_TYPE_NORMAL:
		result.Type = TASK_TYPE_NORMAL_STR
	case TASK_TYPE_URGENT:
		result.Type = TASK_TYPE_URGENT_STR
	default:
		result.Type = TASK_TYPE_NORMAL_STR
	}
	records, errr := QueryActivitiesByTask(db, t)
	result.Record = []ActivityObject{}
	if errr == nil {
		for i := range records {
			switch records[i].Status {
			case ACTIVITY_STATUS_COMPLETED:
				result.Complete++
			case ACTIVITY_STATUS_STOPPED:
				result.Interrupt++
			}
			result.Record = append(result.Record, records[i].ToActivityObject())
		}
	}

	result.Title = t.Title
	result.Description = t.Description

	if t.Create != nil {
		c1, err1 := t.Create.MarshalText()
		if err1 != nil {
			result.Create = ""
		} else {
			result.Create = string(c1)
		}
	} else {
		result.Create = ""
	}

	if t.Deadline != nil {
		c2, err2 := t.Deadline.MarshalText()
		if err2 != nil {
			result.Deadline = ""
		} else {
			result.Deadline = string(c2)
		}
	} else {
		result.Deadline = ""
	}

	result.Estimate = t.Estimate

	switch t.Status {
	case TASK_STATUS_COMPLETED:
		result.Status = TASK_STATUS_COMPLETED_STR
	case TASK_STATUS_RUNNING:
		result.Status = TASK_STATUS_RUNNING_STR
	case TASK_STATUS_STOPPED:
		result.Status = TASK_STATUS_STOPPED_STR
	default:
		result.Status = TASK_STATUS_STOPPED_STR
	}

	return result
}

type ResponseObject struct {
	Success bool        `json:"success"`
	ErrCode int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
