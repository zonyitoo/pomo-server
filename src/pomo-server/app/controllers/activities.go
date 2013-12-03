package controllers

import "github.com/robfig/revel"

type ActivitiesController struct {
	*revel.Controller
}

const (
	ACTIVITY_TYPE_NORMAL = "normal"
	ACTIVITY_TYPE_URGENT = "urgent"

	ACTIVITY_STATUS_COMPLETED   = "completed"
	ACTIVITY_STATUS_RUNNING     = "running"
	ACTIVITY_STATUS_STOPPED     = "stopped"
	ACTIVITY_STATUS_INCOMPLETED = "incompleted"
)

type ActivityObject struct {
	Id        string
	Type      string
	Record    string
	Date      string
	Deadline  string
	Estimate  int
	Complete  int
	Interrupt int
	Status    string
}

func (app ActivitiesController) Queryid(source string, access_token string, id string) revel.Result {
	app.Validation.Required(id)
	app.Validation.Required(source)
	app.Validation.Required(access_token)

	if app.Validation.HasErrors() {
		return app.RenderJson(app.Validation.Errors)
	}

	date := "2013-12-03T02:11Z+08:00"
	resp := ActivityObject{
		Id:        id,
		Type:      ACTIVITY_TYPE_NORMAL,
		Record:    "",
		Date:      date,
		Deadline:  "",
		Estimate:  10,
		Complete:  1,
		Interrupt: 0,
		Status:    ACTIVITY_STATUS_INCOMPLETED,
	}
	return app.RenderJson(resp)
}

type ActivityObjectList struct {
	Type       string
	Date       string
	Status     string
	Activities []ActivityObject
}

func (app ActivitiesController) Querylist(source, access_token, Type, date, status string) revel.Result {

	resp := ActivityObjectList{
		Type:       ACTIVITY_TYPE_NORMAL,
		Date:       "",
		Status:     ACTIVITY_STATUS_COMPLETED,
		Activities: []ActivityObject{},
	}

	return app.RenderJson(resp)
}

func (app ActivitiesController) Update() revel.Result {
	resp := ActivityObject{}
	return app.RenderJson(resp)
}
