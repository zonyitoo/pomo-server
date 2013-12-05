package controllers

import (
	"errors"
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
	"pomo-server/app/models"
	"time"
)

type ActivitiesController struct {
	MgoController
}

func (app ActivitiesController) Queryid(source string, access_token string, id string) revel.Result {
	app.Validation.Required(id)
	app.Validation.Required(source)
	app.Validation.Required(access_token)

	if app.Validation.HasErrors() {
		resp := models.ResponseObject{
			Success: false,
			ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
			Message: "`source`, `access_token` and `id` are REQUIRED",
		}
		return app.RenderJson(resp)
	}

	resp := models.ResponseObject{
		Success: true,
		ErrCode: RESPONSE_STATUS_SUCCESS,
	}

	if !bson.IsObjectIdHex(id) {
		resp.Success = false
		resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
		resp.Message = "Invalid id"
		return app.RenderJson(resp)
	}

	activity := models.Activity{}
	err := app.Db.C(models.ACTIVITY_COLLECTION_NAME).FindId(bson.ObjectIdHex(id)).One(&activity)
	if err != nil {
		resp.Success = false
		resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
		resp.Message = "Invalid id"
		return app.RenderJson(resp)
	}

	resp.Data = activity

	return app.RenderJson(resp)
}

func (app ActivitiesController) processUpdateParams(activity *models.Activity) error {
	if taskid, ok := app.Params.Values["taskid"]; ok {
		if !bson.IsObjectIdHex(taskid[0]) {
			return errors.New("Invalid taskid")
		}

		activity.TaskId = bson.ObjectIdHex(taskid[0])
	}

	if begin, ok := app.Params.Values["begin"]; ok {
		var t time.Time
		err := t.UnmarshalText([]byte(begin[0]))
		if err != nil {
			return errors.New("Invalid begin time")
		}
		activity.Begin = &t
	}

	if end, ok := app.Params.Values["end"]; ok {
		var t time.Time
		err := t.UnmarshalText([]byte(end[0]))
		if err != nil {
			return errors.New("Invalid end time")
		}
		activity.End = &t
	}

	if status, ok := app.Params.Values["status"]; ok {
		stat, err := models.ActivityStatusCode(status[0])
		if err != nil {
			return errors.New("Invalid status")
		}
		activity.Status = stat
	}

	return nil
}

func (app ActivitiesController) Update(source, access_token string) revel.Result {
	app.Validation.Required(source)
	app.Validation.Required(access_token)

	if app.Validation.HasErrors() {
		resp := models.ResponseObject{
			Success: false,
			ErrCode: RESPONSE_STATUS_UNRECOGNIZED_PARAM,
			Message: "`source` and `access_token` are REQUIRED",
		}
		return app.RenderJson(resp)
	}

	resp := models.ResponseObject{
		Success: true,
		ErrCode: RESPONSE_STATUS_SUCCESS,
	}

	activity := models.Activity{}

	if id, ok := app.Params.Values["id"]; ok {
		// Update
		if !bson.IsObjectIdHex(id[0]) {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = "Invalid id"
			return app.RenderJson(resp)
		}

		err := app.Db.C(models.ACTIVITY_COLLECTION_NAME).FindId(bson.ObjectIdHex(id[0])).One(&activity)
		if err != nil {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = "Invalid id"
			return app.RenderJson(resp)
		}

		err = app.processUpdateParams(&activity)
		if err != nil {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = err.Error()
			return app.RenderJson(resp)
		}

		err = app.Db.C(models.ACTIVITY_COLLECTION_NAME).UpdateId(bson.ObjectIdHex(id[0]), activity)
		if err != nil {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_PROCESSING_ERROR
			resp.Message = err.Error()
			return app.RenderJson(resp)
		}
	} else {
		// New

		err := app.processUpdateParams(&activity)
		if err != nil {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = err.Error()
			return app.RenderJson(resp)
		}

		if !activity.TaskId.Valid() {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = "Invalid taskid"
			return app.RenderJson(resp)
		}

		if activity.Begin == nil {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_UNRECOGNIZED_PARAM
			resp.Message = "Invalid begin time"
			return app.RenderJson(resp)
		}

		activity.Id = bson.NewObjectId()
		err = app.Db.C(models.ACTIVITY_COLLECTION_NAME).Insert(activity)

		if err != nil {
			resp.Success = false
			resp.ErrCode = RESPONSE_STATUS_PROCESSING_ERROR
			resp.Message = err.Error()
			return app.RenderJson(resp)
		}
	}

	resp.Data = activity
	return app.RenderJson(resp)
}
