package controllers

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
	"pomo-server/app/models"
	"time"
)

type TasksController struct {
	MgoController
}

func (c TasksController) Queryid(source, access_token, id string) revel.Result {

	if !bson.IsObjectIdHex(id) {
		resp := models.ResponseObject{
			Success: false,
			ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
		}
		return c.RenderJson(resp)
	}
	var t models.Task
	err := c.Db.C(models.TASK_COLLECTION_NAME).FindId(bson.ObjectIdHex(id)).One(&t)

	if err != nil {
		resp := models.ResponseObject{
			Success: false,
			ErrCode: RESPONSE_STATUS_NOT_FOUND,
		}
		return c.RenderJson(resp)
	}

	resp := models.ResponseObject{
		Success: true,
		ErrCode: RESPONSE_STATUS_SUCCESS,
		Data:    t.ToTaskObject(c.Db),
	}

	return c.RenderJson(resp)
}

func (c TasksController) Querylist(source, access_token, date, status string) revel.Result {

	var tlist models.TaskObjectList

	tasktype := c.Params.Get("type")

	cond := bson.M{}
	if date != "" {
		var t time.Time
		err := t.UnmarshalText([]byte(date))
		if err != nil {
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
			}
			return c.RenderJson(resp)
		}
		cond["create"] = t
		tlist.Date = date
	}
	if status != "" {
		switch status {
		case models.TASK_STATUS_STOPPED_STR:
			cond["status"] = models.TASK_STATUS_STOPPED
		case models.TASK_STATUS_RUNNING_STR:
			cond["status"] = models.TASK_STATUS_RUNNING
		case models.TASK_STATUS_COMPLETED_STR:
			cond["status"] = models.TASK_STATUS_COMPLETED
		default:
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
			}
			return c.RenderJson(resp)
		}
		tlist.Status = status
	}
	if tasktype != "" {
		switch tasktype {
		case models.TASK_TYPE_NORMAL_STR:
			cond["type"] = models.TASK_TYPE_NORMAL
		case models.TASK_TYPE_URGENT_STR:
			cond["type"] = models.TASK_TYPE_URGENT
		default:
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
			}
			return c.RenderJson(resp)
		}
	}

	iter := c.Db.C(models.TASK_COLLECTION_NAME).Find(cond).Iter()

	var task models.Task
	for iter.Next(&task) {
		tlist.Tasks = append(tlist.Tasks, task.ToTaskObject(c.Db))
	}

	resp := models.ResponseObject{
		Success: true,
		ErrCode: RESPONSE_STATUS_SUCCESS,
		Data:    tlist,
	}
	return c.RenderJson(resp)
}

func (c TasksController) Update(source, access_token string) revel.Result {
	id := c.Params.Get("id")

	task := models.Task{}

	ttype := c.Params.Get("type")
	task.Title = c.Params.Get("title")
	task.Description = c.Params.Get("description")
	create := c.Params.Get("create")
	deadline := c.Params.Get("deadline")
	c.Params.Bind(&task.Estimate, "estimate")
	status := c.Params.Get("status")

	if ttype != "" {
		switch ttype {
		case models.TASK_TYPE_NORMAL_STR:
			task.Type = models.TASK_TYPE_NORMAL
		case models.TASK_TYPE_URGENT_STR:
			task.Type = models.TASK_TYPE_URGENT
		default:
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
			}
			return c.RenderJson(resp)
		}
	}

	if create != "" {
		var tcreate time.Time
		err1 := tcreate.UnmarshalText([]byte(create))
		if err1 != nil {
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
			}
			return c.RenderJson(resp)
		}
		task.Create = &tcreate
	} else {
		task.Create = nil
	}

	if deadline != "" {
		var tdeadline time.Time
		err2 := tdeadline.UnmarshalText([]byte(deadline))
		if err2 != nil {
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
			}
			return c.RenderJson(resp)
		}
		task.Deadline = &tdeadline
	} else {
		task.Deadline = nil
	}

	if status != "" {
		switch status {
		case models.TASK_STATUS_STOPPED_STR:
			task.Status = models.TASK_STATUS_STOPPED
		case models.TASK_STATUS_COMPLETED_STR:
			task.Status = models.TASK_STATUS_COMPLETED
		case models.TASK_STATUS_RUNNING_STR:
			task.Status = models.TASK_STATUS_RUNNING
		default:
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
			}
			return c.RenderJson(resp)
		}
	}

	if id != "" {
		if !bson.IsObjectIdHex(id) {
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
			}
			return c.RenderJson(resp)
		}

		err := c.Db.C(models.TASK_COLLECTION_NAME).UpdateId(bson.ObjectIdHex(id), &task)

		if err != nil {
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_PROCESSING_ERROR,
			}
			return c.RenderJson(resp)
		}
	} else {
		err := c.Db.C(models.TASK_COLLECTION_NAME).Insert(&task)
		if err != nil {
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_PROCESSING_ERROR,
			}
			return c.RenderJson(resp)
		}
	}

	resp := models.ResponseObject{
		Success: true,
		ErrCode: RESPONSE_STATUS_SUCCESS,
	}

	return c.RenderJson(resp)
}

func (c TasksController) Delete() revel.Result {
	return c.RenderText("")
}
