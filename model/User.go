package model

type User struct {
	Id int64 `bson:id`
	Name string `bson:name`
	Money float64 `bson:money`
}


var UserCol = "user"