package model

type User struct {
	Id string `bson:id`
	Name string `bson:name`
	Money float64 `bson:money`
}


var UserCol = "user"