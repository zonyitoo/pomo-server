package controllers

import "github.com/robfig/revel"
import "pomo-server/app/models"

type TasksController struct {
	MgoController
}

func (c TasksController) Queryid(source, access_token, id string) revel.Result {

	//var t models.Task

	//return c.RenderJson(t.ToTaskObject())
	return c.RenderText("Queryid %s %s %s", source, access_token, id)
}

func (c TasksController) Querylist(source, access_token, date, status string) revel.Result {

	var tlist models.TaskObjectList
	tlist.Type = models.TASK_TYPE_NORMAL_STR

	return c.RenderJson(tlist)
}

func (c TasksController) Update(source string) revel.Result {

	return c.RenderText(source)
}

func (c TasksController) Delete() revel.Result {
	return c.RenderText("")
}
