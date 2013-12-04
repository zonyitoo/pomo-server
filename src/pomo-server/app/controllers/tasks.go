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
			ErrCode: 400,
		}
		return c.RenderJson(resp)
	}
	var t models.Task
	err := c.Db.C(models.TASK_COLLECTION_NAME).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&t)

	if err != nil {
		resp := models.ResponseObject{
			Success: false,
			ErrCode: 404,
		}
		return c.RenderJson(resp)
	}

	resp := models.ResponseObject{
		Success: true,
		ErrCode: 200,
		Data:    t.ToTaskObject(c.Db),
	}

	return c.RenderJson(resp)
}

func (c TasksController) Querylist(source, access_token, tasktype, date, status string) revel.Result {

	var tlist models.TaskObjectList

	cond := bson.M{}
	if date != "" {
		var t time.Time
		err := t.UnmarshalText([]byte(date))
		if err != nil {
			resp := models.ResponseObject{
				Success: false,
				ErrCode: 400,
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
		case "":

		default:
			resp := models.ResponseObject{
				Success: false,
				ErrCode: 400,
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
		case "":

		default:
			resp := models.ResponseObject{
				Success: false,
				ErrCode: 400,
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
		ErrCode: 200,
		Data:    tlist,
	}
	return c.RenderJson(resp)
}

func (c TasksController) Update(source, access_token string) revel.Result {
	id := c.Params.Get("id")
	return c.RenderText(id)
}

func (c TasksController) Delete() revel.Result {
	return c.RenderText("")
}
