package handler

import (
	"github.com/godaner/go-route/route"
	"github.com/smartwalle/alipay"
	"github.com/godaner/go-util/httputil"
	"go-alipay/mgosess"
	"go-alipay/model"
	"log"
	"github.com/godaner/go-util/timeutil"
	"github.com/skip2/go-qrcode"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

const(
	APP_ID="2016081500252906"
	PARTNER_ID="2088102171304735"
	PUBLIC_KEY="MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtztDI8yGwWyEicfCS15dKJMoL3MuuQxuZdnFwPLgnC94xbQ3xMDSkcRucNhDqF1rgxn2cGnj8maJ1s5K4UIlN7YnUe0RAQ+ZORcjAOPncc5uNdEBPsuv6uvF1Vvet7re9DtFNLL5Sv09gbim1mdykOtCmDhikgXxUMo4arRc+Mj+Ax5V4qeZHcqZKYBxUZ2yB0FsMRYY2zHS6pfOA0Zlw/Jp53+FISWwMiMkr/R0iL/N89ouTTDd0Uswl7ynTCRsy8k/IdxSSMCLgTEhhRQJk+d1MAn8MpLQ3syFtmmr7MwcusMbWGY+DI2YvvA8COcuywaBLbROsA0U/F/uhgsYDwIDAQAB"
	PRIVATE_KEY="MIIEpAIBAAKCAQEA3NvjiANKofzgB/G32qBzpxMCv+LeA385GgkgsSb32eqGkQ0z4+/xKOJD1HNbcw9A1GlFOBkd3gtPpRqXNg1UzrbbfD20GGKgWxvvu59wWX7+yGv1VZxd4Ned/bsqt/Ehiae9Uzzu+QdT7LkpToDyFpbpmlUDFtU8eVRjooFwrhjuiHQDCQvxgfJQFr6Wgev2zJuOAhbZPKr1q6hU0Vwp+M/KEcJ0H6y3ME2nq2y3UHI/uVWxlzIFKUhUOeHzwU6njlGZVpxSWAGO4PR94D1FjauXHVT3VYmkho1lBMX/atGH50dztpZCNmuASTICXm7QzQNjoytc3Inz7f2IiY/6ZQIDAQABAoIBAQDTlNQlOQIGrXryIV9A0vX42P2JVo3aS8coahYnoG1RFpxT2ZIci4E86YaRwPGbCMHSvAbT9zvlvHBQe2jhH2RfpE4gd/xQYCu/HMRNujCnjEIJP2OI4IQPIoD40pXaIec/OLLnSzik2aBM6BXqXx2NBoZEk6yh8yY8FdxdkiOgo/mON1uyZiHqKlQ80KHHFlO+j4uoyvWvcZDBa/2rZNixPyfKjqxQY+Z1AQP0F7KrVHGCIYbNGazGmeUKjTegKTKSCJW5OvQlz7NKcyX77CLdWJDw+tTE98gpUiLFzWoJwVAZUeuw3zdzRzQeP1KvN6y1l4XXvLgryCmCpHWNSzABAoGBAPubyVhS3FKrhMKF8iYSlQdFYQ1DC01SFINV7HChG1OxT3e6d1FCj2mf4A5cT67axMqu35+MWzv188Tnb+BE9MTi17X4anvkZVzyNQIZwXX4li+Aswuj528HeCerq096sV06eGxYRMOELy3RzsjyAbu5ExDPLGLxQ1yebr2W2q4xAoGBAOC2tbeW3RMolHE4919gPHgvZxXI3omXLpWeCFsXD3/P4pnNzrJDmtx3vLKqmNSEQyEuZVUt75oYr3TREP2kJNWkMGWUJSBuHURaoMD/j0HX6FTIr7pfpD476NPSxnaxlHLaE6PulXDKP9CeD4Smk9nqsKsZqVEsxl+uwCVMNr51AoGAKY4pSvUkfmLlolsNC0y2YGUNCmuwSizsRmTCkDeLk98NCPE7E6iylvdYwgZtAqwPJIqNVISR7O1KPZMb3yaEZqAwkFVfAOdP9nHqX1ZxpVJRO7c92wZPCv9o0OSBfrb85DDsArjxYnlKcX7dcELIFoESrQwKgKzoOzvPfnCT3YECgYAyTdSKH7SnyErKudTuve1rfjkWcFwY8wYTqkf2lEf89b2lbqQZ1faB6jAM0xHZaS5Z7Df1BI6BsYr9nJuwHCG+zb2jz7Er/FqC5cLc0ZjxRXMWH8Lu9uAeqmyplqKlCYXQ3C7PYOaFVFFK00doC55Hhzk7ZEg6Csrxun37G3+ZzQKBgQCpmYiICnFz6D8FvRB+R1kaMxhmq3NXFTBRH+8uazfYf2q38o/X5CyEpElqT5oyW8LMthAxnganCgr9qnHx5+T0yDdE4JkiE4tsaJKtB0jUjvVJllr46VK20faPqhPAKm40sP63yqC6qIvvNByq0aFNc5r+iuSOweY1If9Z+ssOmw=="
	RETURN_URL="http://zk.godaner.link/alipay/payReturn"
	NOTIFY_URL="http://zk.godaner.link/alipay/payNotify"
	PRODUCT_CODE="FAST_INSTANT_TRADE_PAY"
	IS_PRODUCTION=false
	TRADE_SUCCESS="TRADE_SUCCESS"
)
//alipay client
var client = alipay.New(APP_ID, PARTNER_ID, PUBLIC_KEY, PRIVATE_KEY, IS_PRODUCTION)

//mobile wap page , it will try to open alipay app
//url like: http://localhost/alipay/pay/mobile?subject=支付午餐&amount=10000
func MobilePayHandler(response route.RouteResponse, request route.RouteRequest) {
	//var
	subject:=request.Params["subject"].(string)
	tradeNo:= timeutil.UnixStr()
	amount:=request.Params["amount"].(string)
	//param
	var p = alipay.AliPayTradeWapPay{} //mobile wap page , it will try to open alipay app
	//var p = alipay.AliPayTradePagePay{} //pc web page
	p.NotifyURL = NOTIFY_URL
	p.ReturnURL = RETURN_URL
	p.Subject = subject
	p.OutTradeNo = tradeNo
	p.TotalAmount = amount
	p.ProductCode = PRODUCT_CODE

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
	c:=session.DB(mgosess.DB).C(model.TradeCol)
	c.Insert(model.Trade{
		Subject:subject,
		TradeNo:tradeNo,
		Amount:amount,
		Status:model.UN_FINISH,
		CreateTime:timeutil.Unix(),
	})
	//res
	var payURL = url.String()
	log.Println("payURL is : " + payURL)

	httputil.WriteTemplate(response.ResponseWriter,getForwardAlipayHtml(payURL))

}
//qr pay
func QrPayHandler(response route.RouteResponse, request route.RouteRequest) {

	//var
	subject:=request.Params["subject"].(string)
	tradeNo:= timeutil.UnixStr()
	amount:=request.Params["amount"].(string)
	//param
	var p = alipay.AliPayTradePreCreate{} //mobile wap page , it will try to open alipay app
	//var p = alipay.AliPayTradePagePay{} //pc web page
	p.NotifyURL = NOTIFY_URL
	p.ReturnURL = RETURN_URL
	p.Subject = subject
	p.OutTradeNo = tradeNo
	p.TotalAmount = amount

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
	c:=session.DB(mgosess.DB).C(model.TradeCol)
	c.Insert(model.Trade{
		Subject:subject,
		TradeNo:tradeNo,
		Amount:amount,
		Status:model.UN_FINISH,
		CreateTime:timeutil.Unix(),
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
	tradeNo:= timeutil.UnixStr()
	amount:=request.Params["amount"].(string)
	//param
	//var p = alipay.AliPayTradeWapPay{} //mobile wap page , it will try to open alipay app
	var p = alipay.AliPayTradePagePay{} //pc web page
	p.NotifyURL = NOTIFY_URL
	p.ReturnURL = RETURN_URL
	p.Subject = subject
	p.OutTradeNo = tradeNo
	p.TotalAmount = amount
	p.ProductCode = PRODUCT_CODE

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
	c:=session.DB(mgosess.DB).C(model.TradeCol)
	c.Insert(model.Trade{
		Subject:subject,
		TradeNo:tradeNo,
		Amount:amount,
		Status:model.UN_FINISH,
		CreateTime:timeutil.Unix(),
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

	if request.Params == nil {
		return
	}

	if request.Params["trade_status"] == TRADE_SUCCESS {
		session:=mgosess.OpenSession()
		c:=session.DB(mgosess.DB).C(model.TradeCol)

		tradeNo:=request.Params["out_trade_no"]
		// exits
		trade := &model.Trade{}
		c.Find(bson.M{"tradeno":tradeNo}).One(trade)
		if trade==nil {
			log.Println("tradeno is : ",tradeNo," 不存在")
			return
		}
		if trade.Status != model.UN_FINISH {
			log.Println("tradeno is : ",tradeNo," 已经结束")
			return
		}
		// update
		c.Update(bson.M{"tradeno":tradeNo},bson.M{"$set":bson.M{"status":model.SUCCESS}})
		log.Println("tradeno is : ",tradeNo," - 支付成功")
		response.ResponseWriter.Write([]byte("success"))
	}


}