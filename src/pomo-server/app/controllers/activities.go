package controllers

import (
	"github.com/robfig/revel"
	"pomo-server/app/models"
)

type ActivitiesController struct {
	MgoController
}

func (app ActivitiesController) Queryid(source string, access_token string, id string) revel.Result {
	app.Validation.Required(id)
	app.Validation.Required(source)
	app.Validation.Required(access_token)

	if app.Validation.HasErrors() {
		return app.RenderJson(app.Validation.ErrorMap())
	}

	resp := models.Activity{}

	return app.RenderJson(resp.ToActivityObject())
}

func (app ActivitiesController) Update() revel.Result {
	resp := models.Activity{}
	return app.RenderJson(resp)
}
