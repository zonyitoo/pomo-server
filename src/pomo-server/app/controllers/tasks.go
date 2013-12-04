package controllers

import (
	"errors"
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
			Message: "Invalid Id",
		}
		return c.RenderJson(resp)
	}
	var t models.Task
	err := c.Db.C(models.TASK_COLLECTION_NAME).FindId(bson.ObjectIdHex(id)).One(&t)

	if err != nil {
		resp := models.ResponseObject{
			Success: false,
			ErrCode: RESPONSE_STATUS_NOT_FOUND,
			Message: id + " not found",
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

	tlist := models.TaskObjectList{Tasks: []models.TaskObject{}}

	tasktype := c.Params.Get("type")

	cond := bson.M{}
	if date != "" {
		var t time.Time
		err := t.UnmarshalText([]byte(date))
		if err != nil {
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
				Message: "Invalid date",
			}
			return c.RenderJson(resp)
		}
		cond["create"] = t
		tlist.Date = date
	}
	if status != "" {
		stat, err := models.TaskStatusCode(status)
		if err != nil {
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
				Message: "Invalid status",
			}
			return c.RenderJson(resp)
		}
		cond["status"] = stat
		tlist.Status = status
	}
	if tasktype != "" {
		t, err := models.TaskTypeCode(tasktype)
		if err != nil {
			resp := models.ResponseObject{
				Success: false,
				ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
				Message: "Invalid task type",
			}
			return c.RenderJson(resp)
		}
		cond["type"] = t
		tlist.Type = tasktype
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

func (c TasksController) processParamsToTask(task *models.Task) error {
	if d, ok := c.Params.Values["description"]; ok {
		task.Description = d[0]
	}
	if t, ok := c.Params.Values["title"]; ok {
		task.Title = t[0]
	}

	if ctimestr, ok := c.Params.Values["create"]; ok {
		if ctimestr[0] == "" {
			if task.Create != nil {
				t := time.Now()
				task.Create = &t
			}
		} else {
			var t time.Time
			err := t.UnmarshalText([]byte(ctimestr[0]))
			if err != nil {
				return errors.New("Invalid create time")
			}
			task.Create = &t
		}
	}

	if dtimestr, ok := c.Params.Values["deadline"]; ok {
		if dtimestr[0] != "" {
			var t time.Time
			err := t.UnmarshalText([]byte(dtimestr[0]))
			if err != nil {
				return errors.New("Invalid deadline time")
			}
			task.Deadline = &t
		}
	}
	c.Params.Bind(&task.Estimate, "estimate")
	if task.Estimate < 0 {
		return errors.New("Invalid estimate")
	}
	if stat, ok := c.Params.Values["status"]; ok {
		s, err := models.TaskStatusCode(stat[0])
		if err != nil {
			return errors.New("Invalid status")
		}
		task.Status = s
	}

	if ttype, ok := c.Params.Values["type"]; ok {

		t, err := models.TaskTypeCode(ttype[0])
		if err != nil {
			return errors.New("Invalid type")
		}
		task.Type = t
	}

	return nil
}

func (c TasksController) Update(source, access_token string) revel.Result {

	resp := models.ResponseObject{
		Success: true,
		ErrCode: RESPONSE_STATUS_SUCCESS,
	}

	if id, ok := c.Params.Values["id"]; !ok {
		// New
		task := models.Task{}

		err := c.processParamsToTask(&task)

		if err != nil {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = err.Error()
			return c.RenderJson(resp)
		}

		err = c.Db.C(models.TASK_COLLECTION_NAME).Insert(&task)

		if err != nil {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = "Invalid type"
			return c.RenderJson(resp)
		}
	} else {
		// Update
		if !bson.IsObjectIdHex(id[0]) {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = "Invalid id"
			return c.RenderJson(resp)
		}

		var task models.Task
		err := c.Db.C(models.TASK_COLLECTION_NAME).FindId(bson.ObjectIdHex(id[0])).One(&task)
		if err != nil {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = "Invalid id"
			return c.RenderJson(resp)
		}

		err = c.processParamsToTask(&task)
		if err != nil {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = err.Error()
			return c.RenderJson(resp)
		}

		err = c.Db.C(models.TASK_COLLECTION_NAME).UpdateId(bson.ObjectIdHex(id[0]), &task)
		if err != nil {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = err.Error()
			return c.RenderJson(resp)
		}
	}

	return c.RenderJson(resp)
}

func (c TasksController) Delete() revel.Result {

	resp := models.ResponseObject{
		Success: true,
		ErrCode: RESPONSE_STATUS_SUCCESS,
	}

	if id, ok := c.Params.Values["id"]; ok && bson.IsObjectIdHex(id[0]) {
		err := c.Db.C(models.TASK_COLLECTION_NAME).RemoveId(bson.ObjectIdHex(id[0]))

		if err != nil {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = err.Error()
		}
	} else {
		resp.Success = false
		resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
		resp.Message = "Invalid id"
	}
	return c.RenderJson(resp)
}
