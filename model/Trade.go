package model

const (
	WAIT_TO_PAY = 0
	PAY_SUCCESS = 1
	PAY_FAILURE     =2
)
type Trade struct {
	Id int64 `bson:id`
	UserId int64 `bson:userid`
	Subject string	`bson:subject`
	TradeNo string	`bson:tradeno`
	Amount float64	`bson:amount`
	CreateTime int64 `bson:createtime`
	FinishTime int64 `bson:finishtime`
	Status uint8	`bson:status`
}

var TradeCol = "trade"