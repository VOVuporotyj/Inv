package models

import (
	"gopkg.in/mgo.v2/bson"
)

//Place BLA BLA
type Place struct {
	Number   bson.ObjectId `json:"number" bson:"_number"`
	UserName string        `json:"username" `
	Sistem   string        `json:"sistem" `
	Monitor  string        `json:"monitor" `
	Keybord  string        `json:"keybord" `
	Mouse    string        `json:"mouse" `
	Headset  string        `json:"headset" `
}
