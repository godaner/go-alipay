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
	Body string	`bson:body`
	TradeNo string	`bson:tradeno`
	Amount float64	`bson:amount`
	TimeoutExpress string	`bson:timeoutexpress`
	BuyerId string	`bson:buyerid`
	SellerId string	`bson:sellerid`
	BuyerPayAmount float64 `buyerpayamount`
	PointAmount float64 `pointamount`
	InvoiceAmount float64 `invoiceamount`
	ReceiptAmount float64 `receiptamount`
	GmtCreate string `bson:gmtcreate`
	GmtPayment string `bson:gmtpayment`
	GmtClose string `bson:gmtclose`
	GmtRefund string `bson:gmtrefund`
	GmtFinish string `bson:gmtfinish`
	CreateTime int64 `bson:createtime`
	PaymentTime int64 `bson:paymenttime`
	CloseTime int64 `bson:closetime`
	RefundTime int64 `bson:refundtime`
	FinishTime int64 `bson:finishtime`
	Status uint8	`bson:status`
}

var TradeCol = "trade"