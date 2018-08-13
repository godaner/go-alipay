package main

import (
	"go-alipay/handler"
	"github.com/godaner/go-route/route"
)

const(
	ADDR=":80"
)



func main() {

	//routes
	handler.Routes()

	//run
	route.Start(ADDR)
}

