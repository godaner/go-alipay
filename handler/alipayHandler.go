package handler

import (
	"github.com/godaner/go-route/route"
	"github.com/smartwalle/alipay"
	"go-util/httputil"
	"go-alipay/mgosess"
	"go-alipay/model"
	"log"
	"gopkg.in/mgo.v2/bson"
	"go-util/timeutil"
)

const(
	APP_ID="2016081500252906"
	PARTNER_ID="2088102171304735"
	PUBLIC_KEY="MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2yoq51fxpfdGbSIIUQfy5ixLvWGDOYinZceg5LwfPevpJiiotvvxxwE25FeNq3jeNrJeoVmCMQabPA4R4psXjnHOhjpGxrcLYLcNruf8lkjaLIS2g5WgiCc/wUFN8CB/IkXdLUBJ3KCOdE3Hb9xNfhMRdnGTQctquQygiqDZcD6m+yCCK/iGr2/hnNScMH/nyUrpXni+FpNQtj5wmwiSTY3HM34EOgZ7t4Bi10E8Q9FGgrequ20W5/u7thD2EfwxRaHf2t8u0/ygsAd79KDcD7DXdYEt3kmV/IxxQVF2ImSONZlYcV167YmM8yR3jH1DqJW10tTo+fKds1WqGDjjOQIDAQAB"
	PRIVATE_KEY="MIIEpAIBAAKCAQEA2yoq51fxpfdGbSIIUQfy5ixLvWGDOYinZceg5LwfPevpJiiotvvxxwE25FeNq3jeNrJeoVmCMQabPA4R4psXjnHOhjpGxrcLYLcNruf8lkjaLIS2g5WgiCc/wUFN8CB/IkXdLUBJ3KCOdE3Hb9xNfhMRdnGTQctquQygiqDZcD6m+yCCK/iGr2/hnNScMH/nyUrpXni+FpNQtj5wmwiSTY3HM34EOgZ7t4Bi10E8Q9FGgrequ20W5/u7thD2EfwxRaHf2t8u0/ygsAd79KDcD7DXdYEt3kmV/IxxQVF2ImSONZlYcV167YmM8yR3jH1DqJW10tTo+fKds1WqGDjjOQIDAQABAoIBAC+0QPx0FwQyC6FLjBXfLg7Ny3qgVAjc5trvleTT0dUPmYMVzItv0ZOofwM6z3poZ63uK0zhh8YWEwoCYgA6E+mMehCbdLZiN1bI4XCVzFQF7X9NL9D6a6PXLzhod4dx+1pBbMAhwzIDvJ4yI9wETtXHXsCyPgRPO56l0ff6xPl6MoEvZWfUy/vJg6DBI1O3kd4uyOMdSjXF+8Ird8hFHODHfNG078USagYQWF1G5mfNR481J+lvST3QrbeMyGT50c3uKkZmD68Km9AkVQ7jEZoZxJUeWCsDUc56e28prlHKs0YjsoiHND+Lq7+l/gwT9oi2LhFfkxj+Wdcc05IWa8UCgYEA9I8DTnO1m2pwWY/wJQ9WgSj8d2h59iICaQwr3hlZFo35J8aCUAqXOyA116GsU71D7V/qI4x60UPdznY7tyy7b8gqKxzZcgZxJ+y9E5qqWNyfBJi5dhH0Tu0x0puYbEOk2uDzb27qp4Nz4fM8yDlfVn3kLmtsvX2vGgPMfuC2zFsCgYEA5WsFabcb9ib7xfWVsuRzwjolXY7fzOhFAG1W0mYtFb6/nFXkHT27+EcjO/wBjyjlwwI8ZXz0RRU7a6OamV2hiuu15zt1hZTGu2s6F1I51lKKUacx0+CV3lk7ie0S+bE4K8Tv1fJkPV9IV53e3GLImZ2tw6OIUhhVDPblKnJacvsCgYEAsZk99HYRF4lHh7MA2Vj3IBsMpQaJM1ZlW7YMEWFlEf0OSHVfYxMd2kE8+JgfFjznHHZACYrWEixv4qR4H0Dr6XR6Mw7jVmwZNr40XQ0/0gJ9tI/Yd0b4nWyhdump2k2RZaZhAraQ0A+lUxwaMul/M8d+srsun7mrNIA3vJgiB+sCgYAcA+SmUzD77yjBtMqyDMnCBx22/hKAUEDU24VFriqFRETbz3VKyNYibHV2BsEd/U2JccV0Uzz1DrGx5EmlvtRSRZRyB0XDqTZXfrGaVXUwoeW4MOWZzQwgGd92aVfu2+BTH/p1suLgx7jq0iF74ihC8gldIaQs+kGwkpnLhCldbwKBgQDg7SN6p2nqsRo4BAJHyuC8wfolfOrXBGT5FPYHO6ONg0wTLTcD5QryykhU4mJCDSNvSGZeoas/mo6WO32VQdmshgw2FsckVSVwpcCWWSgYZPy6dD4mPF6j+9Bp/aBIZB+piFc5NJ2AgO7d/ae7jAYxglNytmPUcOOzHzbIw7MeXA=="
	RETURN_URL="http://zk.ngrok.xiaomiqiu.cn/alipay/payReturn"
	NOTIFY_URL="http://zk.ngrok.xiaomiqiu.cn/alipay/payNotify"
	PRODUCT_CODE="FAST_INSTANT_TRADE_PAY"
)
//alipay client
var client = alipay.New(APP_ID, PARTNER_ID, PUBLIC_KEY, PRIVATE_KEY, false)

//url like: http://localhost/alipay/pay?subject=支付午餐&amount=10000
func PayHandler(response route.RouteResponse, request route.RouteRequest) {
	//var
	subject:=request.Params["subject"].(string)
	tradeNo:= timeutil.UnixStr()
	amount:=request.Params["amount"].(string)
	//param
	var p = alipay.AliPayTradeWapPay{}
	p.NotifyURL = NOTIFY_URL
	p.ReturnURL = RETURN_URL
	p.Subject = subject
	p.OutTradeNo = tradeNo
	p.TotalAmount = amount
	p.ProductCode = PRODUCT_CODE

	log.Println("treade is :" ,p)
	//new trade
	var url, err = client.TradeWapPay(p)
	if err != nil {
		log.Println(err)
		return
	}
	//save
	session:=mgosess.OpenSession()
	c:=session.DB(mgosess.DB).C(model.TradeCol)
	c.Insert(model.Trade{
		Subject:subject,
		TradeNo:tradeNo,
		Amount:amount,
		Status:model.UN_FINISH,
	})
	//res
	var payURL = url.String()
	log.Println("payURL is : " + payURL)

	httputil.WriteTemplate(response.ResponseWriter,getForwardAlipayHtml(payURL))

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
func PayOverHandler(response route.RouteResponse, request route.RouteRequest){
	if request.Params == nil{
		log.Println("支付失败")
		return
	}
	tradeNo:=request.Params["out_trade_no"]
	session:=mgosess.OpenSession()
	c:=session.DB(mgosess.DB).C(model.TradeCol)
	c.Update(bson.M{"tradeNo":tradeNo},bson.M{"status":model.SUCCESS})
	log.Println("TreadeNo is : ",tradeNo," - 支付成功")

}