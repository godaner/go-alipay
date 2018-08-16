package handler

import (
	"github.com/godaner/go-route/route"
)

func Routes() {
	route.RegistRoutes(
		route.MakeAnyRoute("/alipay/pay/mobile",MobilePayHandler),
		route.MakeAnyRoute("/alipay/pay/pc",PcPayHandler),
		route.MakeAnyRoute("/alipay/pay/qr",QrPayHandler),
		route.MakeAnyRoute("/alipay/payNotify",PayNotifyHandler),
		route.MakeAnyRoute("/alipay/payReturn",PayReturnHandler),
		route.MakeAnyRoute("/alipay/refund",RefundHandler),)
}

