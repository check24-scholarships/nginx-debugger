package main

import "nginx_debugger/api"

func main() {
	httpServer := api.NewHttpServer("127.0.0.1", 9000)
	httpServer.ListenAndServe()
}
