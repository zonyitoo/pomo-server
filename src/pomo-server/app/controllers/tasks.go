package controllers

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"pomo-server/app/models"
	"time"
)

type TasksController struct {
	MgoController
}

func (c TasksController) Queryid(source, access_token, id string) revel.Result {

	if !bson.IsObjectIdHex(string) {
		resp := models.RESTResponseObject{
			Success: false,
			ErrCode: 400,
		}
		return c.RenderJson(resp)
	}
	var t models.Task
	err := c.Db.C(models.TASK_COLLECTION_NAME).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&t)

	if err != nil {
		resp := models.RESTResponseObject{
			Success: false,
			ErrCode: 404,
		}
		return c.RenderJson(resp)
	}

	return c.RenderJson(t.ToTaskObject(c.Db))
}

func (c TasksController) Querylist(source, access_token, tasktype, date, status string) revel.Result {

	var tlist models.TaskObjectList

	cond := bson.M{}
	if date {
		var t time.Time
		err := t.UnmarshalText(date)
		if err != nil {
			resp := models.RESTResponseObject{
				Success: false,
				ErrCode: 400,
			}
			return c.RenderJson(resp)
		}
		cond["create"] = t
		tlist.Date = date
	}
	if status {
		switch status {
		case models.TASK_STATUS_STOPPED_STR:
			cond["status"] = models.TASK_STATUS_STOPPED
		case models.TASK_STATUS_RUNNING_STR:
			cond["status"] = models.TASK_STATUS_RUNNING
		case models.TASK_STATUS_COMPLETED_STR:
			cond["status"] = models.TASK_STATUS_COMPLETED
		case "":

		default:
			resp := models.RESTResponseObject{
				Success: false,
				ErrCode: 400,
			}
			return c.RenderJson(resp)
		}
		tlist.Status = status
	}
	if tasktype {
		switch tasktype {
		case models.TASK_TYPE_NORMAL_STR:
			cond["type"] = models.TASK_TYPE_NORMAL
		case models.TASK_TYPE_URGENT_STR:
			cond["type"] = models.TASK_TYPE_URGENT
		case "":

		default:
			resp := models.RESTResponseObject{
				Success: false,
				ErrCode: 400,
			}
			return c.RenderJson(resp)
		}
	}

	iter := c.Db.C(models.TASK_COLLECTION_NAME).Find(cond).Iter()

	var task models.Task
	for iter.Next(&task) {
		tlist.Task = append(tlist.Task, task.ToTaskObject(c.Db))
	}

	return c.RenderJson(tlist)
}

func (c TasksController) Update(source, access_token string) revel.Result {

	return c.RenderText(source)
}

func (c TasksController) Delete() revel.Result {
	return c.RenderText("")
}
