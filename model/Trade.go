package model

const (
	WAIT_TO_PAY = 0
	PAY_SUCCESS = 1
	PAY_FAILURE     =2
)
const (
	LOCKED=1
	UNLOCKED=2
)
type Trade struct {
	Id int64 `bson:id`
	UserId int64 `bson:userid`
	Subject string	`bson:subject`
	TradeNo string	`bson:tradeno`
	Amount string	`bson:amount`
	CreateTime int64 `bson:createtime`
	FinishTime int64 `bson:finishtime`
	Lock uint8	`bson:lock`
	Status uint8	`bson:status`
}

var TradeCol = "trade"