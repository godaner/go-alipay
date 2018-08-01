package handler

import (
	"github.com/godaner/go-route/route"
)

func Routes() route.Router{
	return route.RegistRoutes(
		route.MakeAnyRoute("/alipay/pay",PayHandler),
		route.MakeAnyRoute("/alipay/payNotify",PayNotifyHandler),
		route.MakeAnyRoute("/alipay/payReturn",PayReturnHandler),)
}

