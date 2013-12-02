package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
}

type TestJsonResponse struct {
	Hello   string
	Boolean bool
	List    []string
}

func (c App) Index() revel.Result {
	//return c.Render()
	resp := TestJsonResponse{"World", true, []string{"str1", "str2"}}
	return c.RenderJson(resp)
}
