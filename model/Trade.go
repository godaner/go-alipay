package model

const (
	UN_FINISH = 0
	SUCCESS = 1
	FAILURE=2
)
type Trade struct {
	Subject string	`bson:subject`
	TradeNo string	`bson:tradeno`
	Amount string	`bson:amount`
	Status uint8	`bson:status`
}

var TradeCol = "trade"