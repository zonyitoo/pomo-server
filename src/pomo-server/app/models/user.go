package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `bson:"name"`
}

type UserAuthInfo struct {
	Expire time.Time `json:"expire"`
	UserId string    `json:"userid"`
}
