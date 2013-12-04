package controllers

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
)

type MgoController struct {
	*revel.Controller
	Db *mgo.Database
}

var G_DBSession *mgo.Session = nil

func init() {
	revel.OnAppStart(BeginSession)
	revel.InterceptMethod((*MgoController).OpenDB, revel.BEFORE)
	revel.InterceptMethod((*MgoController).CloseDB, revel.AFTER)
}

func BeginSession() {
	sec, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	G_DBSession = sec
}

func (mc *MgoController) OpenDB() revel.Result {
	mc.db = G_DBSession.DB("pomo")
	return nil
}

func (mc *MgoController) CloseDB() revel.Result {
	mc.db = nil
	return nil
}
