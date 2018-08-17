package handler

import (
	"github.com/godaner/go-route/route"
	"github.com/smartwalle/alipay"
	"go-alipay/mgosess"
	"go-alipay/model"
	"log"
	"github.com/skip2/go-qrcode"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
	"gopkg.in/mgo.v2"
	"github.com/godaner/go-util"
)

const(
	APP_ID="2016081500252906"
	PARTNER_ID="2088102171304735"
	PUBLIC_KEY="MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtztDI8yGwWyEicfCS15dKJMoL3MuuQxuZdnFwPLgnC94xbQ3xMDSkcRucNhDqF1rgxn2cGnj8maJ1s5K4UIlN7YnUe0RAQ+ZORcjAOPncc5uNdEBPsuv6uvF1Vvet7re9DtFNLL5Sv09gbim1mdykOtCmDhikgXxUMo4arRc+Mj+Ax5V4qeZHcqZKYBxUZ2yB0FsMRYY2zHS6pfOA0Zlw/Jp53+FISWwMiMkr/R0iL/N89ouTTDd0Uswl7ynTCRsy8k/IdxSSMCLgTEhhRQJk+d1MAn8MpLQ3syFtmmr7MwcusMbWGY+DI2YvvA8COcuywaBLbROsA0U/F/uhgsYDwIDAQAB"
	PRIVATE_KEY   ="MIIEpAIBAAKCAQEA3NvjiANKofzgB/G32qBzpxMCv+LeA385GgkgsSb32eqGkQ0z4+/xKOJD1HNbcw9A1GlFOBkd3gtPpRqXNg1UzrbbfD20GGKgWxvvu59wWX7+yGv1VZxd4Ned/bsqt/Ehiae9Uzzu+QdT7LkpToDyFpbpmlUDFtU8eVRjooFwrhjuiHQDCQvxgfJQFr6Wgev2zJuOAhbZPKr1q6hU0Vwp+M/KEcJ0H6y3ME2nq2y3UHI/uVWxlzIFKUhUOeHzwU6njlGZVpxSWAGO4PR94D1FjauXHVT3VYmkho1lBMX/atGH50dztpZCNmuASTICXm7QzQNjoytc3Inz7f2IiY/6ZQIDAQABAoIBAQDTlNQlOQIGrXryIV9A0vX42P2JVo3aS8coahYnoG1RFpxT2ZIci4E86YaRwPGbCMHSvAbT9zvlvHBQe2jhH2RfpE4gd/xQYCu/HMRNujCnjEIJP2OI4IQPIoD40pXaIec/OLLnSzik2aBM6BXqXx2NBoZEk6yh8yY8FdxdkiOgo/mON1uyZiHqKlQ80KHHFlO+j4uoyvWvcZDBa/2rZNixPyfKjqxQY+Z1AQP0F7KrVHGCIYbNGazGmeUKjTegKTKSCJW5OvQlz7NKcyX77CLdWJDw+tTE98gpUiLFzWoJwVAZUeuw3zdzRzQeP1KvN6y1l4XXvLgryCmCpHWNSzABAoGBAPubyVhS3FKrhMKF8iYSlQdFYQ1DC01SFINV7HChG1OxT3e6d1FCj2mf4A5cT67axMqu35+MWzv188Tnb+BE9MTi17X4anvkZVzyNQIZwXX4li+Aswuj528HeCerq096sV06eGxYRMOELy3RzsjyAbu5ExDPLGLxQ1yebr2W2q4xAoGBAOC2tbeW3RMolHE4919gPHgvZxXI3omXLpWeCFsXD3/P4pnNzrJDmtx3vLKqmNSEQyEuZVUt75oYr3TREP2kJNWkMGWUJSBuHURaoMD/j0HX6FTIr7pfpD476NPSxnaxlHLaE6PulXDKP9CeD4Smk9nqsKsZqVEsxl+uwCVMNr51AoGAKY4pSvUkfmLlolsNC0y2YGUNCmuwSizsRmTCkDeLk98NCPE7E6iylvdYwgZtAqwPJIqNVISR7O1KPZMb3yaEZqAwkFVfAOdP9nHqX1ZxpVJRO7c92wZPCv9o0OSBfrb85DDsArjxYnlKcX7dcELIFoESrQwKgKzoOzvPfnCT3YECgYAyTdSKH7SnyErKudTuve1rfjkWcFwY8wYTqkf2lEf89b2lbqQZ1faB6jAM0xHZaS5Z7Df1BI6BsYr9nJuwHCG+zb2jz7Er/FqC5cLc0ZjxRXMWH8Lu9uAeqmyplqKlCYXQ3C7PYOaFVFFK00doC55Hhzk7ZEg6Csrxun37G3+ZzQKBgQCpmYiICnFz6D8FvRB+R1kaMxhmq3NXFTBRH+8uazfYf2q38o/X5CyEpElqT5oyW8LMthAxnganCgr9qnHx5+T0yDdE4JkiE4tsaJKtB0jUjvVJllr46VK20faPqhPAKm40sP63yqC6qIvvNByq0aFNc5r+iuSOweY1If9Z+ssOmw=="
	//RETURN_URL    ="http://zk.godaner.link/alipay/payReturn"
	RETURN_URL    =""//如果return_url为空，那么支付宝支付完后页面不会跳转
	NOTIFY_URL    ="http://zk.godaner.link/alipay/payNotify"
	PRODUCT_CODE  ="FAST_INSTANT_TRADE_PAY"
	IS_PRODUCTION =false
	TIMEOUT_EXPRESS = "30m"//最晚付款时间
	USER_ID =1
	UNIQUE_ID =1
)

//alipay client
var client = alipay.New(APP_ID, PARTNER_ID, PUBLIC_KEY, PRIVATE_KEY, IS_PRODUCTION)
func init(){
	checkUnSuccessTrade()
}
func checkUnSuccessTrade() {
	ticker := time.NewTicker(time.Second * 30)
	go func() {
		session:=mgosess.OpenSession()
		defer session.Close()
		c:=session.DB(mgosess.DB).C(model.TradeCol)
		for _ = range ticker.C {
			selector:=bson.M{"status":model.TRADE_STATUS_WAIT_BUYER_PAY}
			trades:=make([]model.Trade,0,0)
			c.Find(selector).All(&trades)
			for _,trade:=range trades{
				tradeNo:=trade.TradeNo
				var p = alipay.AliPayTradeQuery{} //mobile wap page , it will try to open alipay app
				p.OutTradeNo = tradeNo
				results,err:=client.TradeQuery(p)

				if err!=nil {
					log.Println("checkUnSuccessTrade call TradeQuery api error ! tradeno is: ",tradeNo," , err is: ",err)
					continue
				}
				//log.Println("checkUnSuccessTrade get trade status ! tradeno is: ",tradeNo," , trade status is: ",results.AliPayTradeQuery.TradeStatus)
				if results.AliPayTradeQuery.TradeStatus == alipay.K_TRADE_STATUS_TRADE_SUCCESS {
					checkParam:=map[string]interface{}{"tradeNo":tradeNo,"totalAmount":results.AliPayTradeQuery.TotalAmount,"partnerId":PARTNER_ID}
					newSetParam:=makeCheckSetTrade(results)
					err:=updateTradeSuccess(c,checkParam,newSetParam)
					if err!=nil {
						log.Println("checkUnSuccessTrade update trade status to success fail ! tradeno is: ",tradeNo," , err is: ",err)
						continue
					}
					log.Println("checkUnSuccessTrade update trade status to success is ok , status syncing ! tradeNo is : ",tradeNo)
				}
			}

		}
	}()
}
func makeCheckSetTrade(results *alipay.AliPayTradeQueryResponse) model.Trade {
	buyerPayAmount,_:=strconv.ParseFloat(results.AliPayTradeQuery.BuyerPayAmount,64)
	pointAmount,_:=strconv.ParseFloat(results.AliPayTradeQuery.PointAmount,64)
	invoiceAmount,_:=strconv.ParseFloat(results.AliPayTradeQuery.InvoiceAmount,64)
	receiptAmount,_:=strconv.ParseFloat(results.AliPayTradeQuery.ReceiptAmount,64)
	return model.Trade{
		BuyerId:results.AliPayTradeQuery.BuyerUserId,
		BuyerPayAmount:buyerPayAmount,
		PointAmount:pointAmount,
		InvoiceAmount:invoiceAmount,
		ReceiptAmount:receiptAmount,
	}
}


//mobile wap page , it will try to open alipay app
//url like: http://localhost/alipay/pay/mobile?subject=支付午餐&amount=10000
func MobilePayHandler(response route.RouteResponse, request route.RouteRequest) {
	//var
	subject:=request.Params["subject"].(string)
	body:=request.Params["body"].(string)
	tradeNo:= go_util.GetSnowFlakeIdStr(UNIQUE_ID)
	amountStr:= fmt.Sprintf("%s",request.Params["amount"])
	amount,_:=strconv.ParseFloat(amountStr,64)
	amount=go_util.Round(amount,2)
	amountStr=fmt.Sprintf("%.2f",amount)

	//param
	var p = alipay.AliPayTradeWapPay{} //mobile wap page , it will try to open alipay app
	//var p = alipay.AliPayTradePagePay{} //pc web page
	p.NotifyURL = NOTIFY_URL
	p.ReturnURL = RETURN_URL
	p.Subject = subject
	p.OutTradeNo = tradeNo
	p.TotalAmount = amountStr
	p.ProductCode = PRODUCT_CODE
	p.TimeoutExpress=TIMEOUT_EXPRESS


	log.Println("treade is :" ,p)
	//new trade
	var url, err = client.TradeWapPay(p)
	//var url, err = client.TradePagePay(p) //
	if err != nil {
		log.Println(err)
		return
	}
	//save
	session:=mgosess.OpenSession()
	defer session.Close()
	c:=session.DB(mgosess.DB).C(model.TradeCol)
	c.Insert(model.Trade{
		Id:go_util.GetSnowFlakeId(UNIQUE_ID),
		UserId:USER_ID,
		Subject:subject,
		Body:body,
		TimeoutExpress:TIMEOUT_EXPRESS,
		TradeNo:tradeNo,
		Amount:amount,
		SellerId:PARTNER_ID,
		Status:model.TRADE_STATUS_WAIT_BUYER_PAY,
		CreateTime:go_util.Unix(),
	})
	//res
	var payURL = url.String()
	log.Println("payURL is : " + payURL)

	go_util.WriteTemplate(response.ResponseWriter,getForwardAlipayHtml(payURL))

}
//qr pay
func QrPayHandler(response route.RouteResponse, request route.RouteRequest) {

	//var
	subject:=request.Params["subject"].(string)
	body:=request.Params["body"].(string)
	tradeNo:= go_util.UnixStr()
	amountStr:= fmt.Sprintf("%s",request.Params["amount"])
	amount,_:=strconv.ParseFloat(amountStr,64)
	//param
	var p = alipay.AliPayTradePreCreate{} //mobile wap page , it will try to open alipay app
	//var p = alipay.AliPayTradePagePay{} //pc web page
	p.NotifyURL = NOTIFY_URL
	p.ReturnURL = RETURN_URL
	p.Subject = subject
	p.OutTradeNo = tradeNo
	p.TotalAmount = amountStr
	p.TimeoutExpress=TIMEOUT_EXPRESS


	log.Println("treade is :" ,p)
	//new trade
	var results, err = client.TradePreCreate(p)
	//var url, err = client.TradePagePay(p) //
	if err != nil {
		log.Println(err)
		return
	}
	//save
	session:=mgosess.OpenSession()
	defer session.Close()
	c:=session.DB(mgosess.DB).C(model.TradeCol)
	c.Insert(model.Trade{
		Id:go_util.GetSnowFlakeId(UNIQUE_ID),
		UserId:USER_ID,
		Subject:subject,
		Body:body,
		TimeoutExpress:TIMEOUT_EXPRESS,
		TradeNo:tradeNo,
		Amount:amount,
		SellerId:PARTNER_ID,
		Status:model.TRADE_STATUS_WAIT_BUYER_PAY,
		CreateTime:go_util.Unix(),
	})
	//res
	var qrCode = results.AliPayPreCreateResponse.QRCode
	log.Println("qrCode is : " + qrCode)
	q, err := qrcode.New(qrCode, qrcode.Medium)
	if err != nil {
		return
	}
	png, err := q.PNG(256)
	if err != nil {
		return
	}
	response.ResponseWriter.Header().Set("Content-Type", "image/png")
	response.ResponseWriter.Header().Set("Content-Length", fmt.Sprintf("%d", len(png)))
	response.ResponseWriter.Write(png)
}
//pc web page
//url like: http://localhost/alipay/pay/pc?subject=支付午餐&amount=10000
func PcPayHandler(response route.RouteResponse, request route.RouteRequest) {
	//var
	subject:=request.Params["subject"].(string)
	body:=request.Params["body"].(string)
	tradeNo:= go_util.UnixStr()
	amountStr:= fmt.Sprintf("%s",request.Params["amount"])
	amount,_:=strconv.ParseFloat(amountStr,64)
	//param
	//var p = alipay.AliPayTradeWapPay{} //mobile wap page , it will try to open alipay app
	var p = alipay.AliPayTradePagePay{} //pc web page
	p.NotifyURL = NOTIFY_URL
	p.ReturnURL = RETURN_URL
	p.Subject = subject
	p.OutTradeNo = tradeNo
	p.TotalAmount = amountStr
	p.ProductCode = PRODUCT_CODE
	p.TimeoutExpress=TIMEOUT_EXPRESS

	log.Println("treade is :" ,p)
	//new trade
	//var url, err = client.TradeWapPay(p) //mobile wap page , it will try to open alipay app
	var url, err = client.TradePagePay(p) //pc web page
	if err != nil {
		log.Println(err)
		return
	}
	//save
	session:=mgosess.OpenSession()
	defer session.Close()
	c:=session.DB(mgosess.DB).C(model.TradeCol)
	c.Insert(model.Trade{
		Id:go_util.GetSnowFlakeId(UNIQUE_ID),
		UserId:USER_ID,
		Subject:subject,
		Body:body,
		TimeoutExpress:TIMEOUT_EXPRESS,
		TradeNo:tradeNo,
		Amount:amount,
		SellerId:PARTNER_ID,
		Status:model.TRADE_STATUS_WAIT_BUYER_PAY,
		CreateTime:go_util.Unix(),
	})
	//res
	var payURL = url.String()
	log.Println("payURL is : " + payURL)

	go_util.WriteTemplate(response.ResponseWriter,getForwardAlipayHtml(payURL))

}

func getForwardAlipayHtml(payURL string)(string){
	return `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
</body>
<script>
	location.href="`+payURL+`"
</script>
</html>
	`
}
func PayNotifyHandler(response route.RouteResponse, request route.RouteRequest) {
	PayOverHandler(response,request)
}
func PayReturnHandler(response route.RouteResponse, request route.RouteRequest) {
	PayOverHandler(response,request)
}

//when pay success , alipay will call this func
func PayOverHandler(response route.RouteResponse, request route.RouteRequest){
	// verify sign
	var noti, _ = client.GetTradeNotification(request.Request)
	// handler service
	if noti!=nil{
		log.Println("PayOverHandler start ! tradeno is: ",noti.OutTradeNo,", trade status is : ",noti.TradeStatus)
		switch noti.TradeStatus {
		case alipay.K_TRADE_STATUS_TRADE_SUCCESS:
			tradeSuccess(noti,request,response)
		case alipay.K_TRADE_STATUS_TRADE_FINISHED:
			tradeFinished(noti,request,response)
		case alipay.K_TRADE_STATUS_TRADE_CLOSED:
			tradeClosed(noti,request,response)
		case alipay.K_TRADE_STATUS_WAIT_BUYER_PAY:
			tradeWaitBuyerPay(noti,request,response)
		}
	}


}

func tradeSuccess(notification *alipay.TradeNotification, request route.RouteRequest, response route.RouteResponse) {
	tradeNo:=notification.OutTradeNo
	session:=mgosess.OpenSession()
	defer session.Close()
	c:=session.DB(mgosess.DB).C(model.TradeCol)

	selector:=bson.M{"tradeno":tradeNo}
	q:=c.Find(selector)
	//is exits ?
	if n,_:=q.Count();n==0{
		log.Println("PayOverHandler Trade is not exits ! tradeno is: ",tradeNo)
		return
	}
	trade:=model.Trade{}
	q.One(&trade)
	//is success ?
	if trade.Status== model.TRADE_STATUS_TRADE_SUCCESS {
		log.Println("PayOverHandler trade have pay success ! tradeno is: ",tradeNo)
		notifyAliPaySuccess(response)
		return
	}

	// update status to "pay success" , only one can access this program
	checkParam:=map[string]interface{}{"tradeNo":tradeNo,"totalAmount":notification.TotalAmount,"partnerId":notification.SellerId}
	newSetParam:= makeNotifySetTrade(notification)
	err:=updateTradeSuccess(c,checkParam,newSetParam)
	if err!=nil {
		log.Println("PayOverHandler Trade Update fail ! tradeno is: ",tradeNo," , err is: ",err)
		return
	}
	notifyAliPaySuccess(response)
	log.Println("tradeSuccess handle trade success ! trade status is : ",notification.TradeStatus," , tradeno is: ",tradeNo)
}
func makeNotifySetTrade(notification *alipay.TradeNotification) model.Trade {
	gmtPayment:=notification.GmtPayment
	gmtCreate:=notification.GmtCreate
	buyerPayAmount,_:=strconv.ParseFloat(notification.BuyerPayAmount,64)
	pointAmount,_:=strconv.ParseFloat(notification.PointAmount,64)
	invoiceAmount,_:=strconv.ParseFloat(notification.InvoiceAmount,64)
	receiptAmount,_:=strconv.ParseFloat(notification.ReceiptAmount,64)
	return model.Trade{
		BuyerId:notification.BuyerId,
		GmtPayment:gmtPayment,
		GmtCreate:gmtCreate,
		BuyerPayAmount:buyerPayAmount,
		PointAmount:pointAmount,
		InvoiceAmount:invoiceAmount,
		ReceiptAmount:receiptAmount,
	}
}
func tradeFinished(notification *alipay.TradeNotification, request route.RouteRequest, response route.RouteResponse) {
	tradeNo:=notification.OutTradeNo

	notifyAliPaySuccess(response)
	log.Println("tradeFinished handle trade success ! trade status is : ",notification.TradeStatus," , tradeno is: ",tradeNo)
}
func tradeClosed(notification *alipay.TradeNotification, request route.RouteRequest, response route.RouteResponse) {
	tradeNo:=notification.OutTradeNo

	notifyAliPaySuccess(response)
	log.Println("tradeClosed handle trade success ! trade status is : ",notification.TradeStatus," , tradeno is: ",tradeNo)
}
func tradeWaitBuyerPay(notification *alipay.TradeNotification, request route.RouteRequest, response route.RouteResponse) {
	tradeNo:=notification.OutTradeNo

	notifyAliPaySuccess(response)
	log.Println("tradeWaitBuyerPay handle trade success ! trade status is : ",notification.TradeStatus," , tradeno is: ",tradeNo)
}

func updateTradeSuccess(c *mgo.Collection,checkParam map[string]interface{},newSet model.Trade)(error){
	partnerId:=checkParam["partnerId"]
	totalAmountStr:=checkParam["totalAmount"].(string)
	amountFloat64,_:=strconv.ParseFloat(totalAmountStr,64)
	tradeNo:=checkParam["tradeNo"]

	selector:=bson.M{"$and":[]bson.M{{"tradeno":tradeNo},{"sellerid":partnerId},{"amount":amountFloat64},{"status":model.TRADE_STATUS_WAIT_BUYER_PAY}}}
	update:=bson.M{"$set":bson.M{
		"status":model.TRADE_STATUS_TRADE_SUCCESS,
		"paymenttime":go_util.Unix(),
		"buyerid":newSet.BuyerId,
		"gmtcreate":newSet.GmtCreate,
		"gmtpayment":newSet.GmtPayment,
	}}
	return c.Update(selector,update)
}
func notifyAliPaySuccess(response route.RouteResponse){
	i,err:=response.ResponseWriter.Write([]byte("success"))
	log.Println("notifyAliPaySuccess notigy alipay finish ! len(notify) is : ",i," , err is : ",err)
}



func RefundHandler(response route.RouteResponse, request route.RouteRequest){
	tradeNo:=request.Params["tradeNo"].(string)
	log.Println("RefundHandler trade start ! trade no is :",tradeNo)
	//exits trade ?
	session:=mgosess.OpenSession()
	defer session.Close()
	c:=session.DB(mgosess.DB).C(model.TradeCol)
	selector:=bson.M{"tradeno":tradeNo}
	q:=c.Find(selector)
	if n,_:=q.Count();n==0{
		log.Println("RefundHandler trade is not exits ! trade no is :",tradeNo)
		return
	}
	// fetch alipay trade status
	var pp = alipay.AliPayTradeQuery{} //mobile wap page , it will try to open alipay app
	pp.OutTradeNo = tradeNo
	res,er:=client.TradeQuery(pp)

	if er!=nil||!res.IsSuccess() {
		log.Println("RefundHandler call TradeQuery api fail ! tradeno is: ",tradeNo," , err is: ",er)
		return
	}
	if res.AliPayTradeQuery.TradeStatus != alipay.K_TRADE_STATUS_TRADE_SUCCESS {
		log.Println("RefundHandler alipay trade status is not TRADE_SUCCESS ! tradeno is: ",tradeNo," , aplipay status is: ",res.AliPayTradeQuery.TradeStatus)
		return
	}
	//check local status
	trade:=model.Trade{}
	q.One(&trade)
	if trade.Status!=model.TRADE_STATUS_TRADE_SUCCESS{
		log.Println("RefundHandler local trade status is not TRADE_SUCCESS ! tradeno is: ",tradeNo," , local status is: ",res.AliPayTradeQuery.TradeStatus)
		return
	}

	//build refund request
	//param
	var p = alipay.AliPayTradeRefund{} //mobile wap page , it will try to open alipay app
	//var p = alipay.AliPayTradePagePay{} //pc web page
	p.OutTradeNo = tradeNo
	p.RefundAmount = fmt.Sprintf("%.2f",trade.Amount)
	var results, err = client.TradeRefund(p)
	//var url, err = client.TradePagePay(p) //
	if err != nil||!results.IsSuccess() {
		log.Println("RefundHandler call refund api fail ! trade no is :",tradeNo," , err is : ",err)
		return
	}
	selector=bson.M{"tradeno":tradeNo}
	update:=bson.M{"$set":bson.M{"status":model.TRADE_STATUS_TRADE_CLOSED}}
	err=c.Update(selector,update)
	if err!=nil {
		log.Println("RefundHandler trade update to close fail ! trade no is :",tradeNo)
		return
	}
	log.Println("RefundHandler trade is refund success ! trade no is :",tradeNo)
}
