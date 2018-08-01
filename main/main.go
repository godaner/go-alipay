package main

import (
	"net/http"
	"log"
	"github.com/godaner/go-route/route"
	"go-alipay/handler"
)

const(
	ADDR=":80"
)


func StartServer(){

	//routes
	router := handler.Routes()

	//run server
	runServer(router)
}
func runServer(router route.Router) {

	err := http.ListenAndServe(ADDR,route.GetDispatcherRouter(router))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}




func main() {
	StartServer()
}

