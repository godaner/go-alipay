package model

const (
	TRADE_STATUS_WAIT_BUYER_PAY = 1 //（交易创建，等待买家付款）
	TRADE_STATUS_TRADE_CLOSED   = 2   //（未付款交易超时关闭，或支付完成后全额退款）
	TRADE_STATUS_TRADE_SUCCESS  = 3  //（交易支付成功）
	TRADE_STATUS_TRADE_FINISHED = 4 //（交易结束，不可退款）
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