package main

import "nginx_debugger/api"

func main() {
	httpServer := api.NewHttpServer("0.0.0.0", 9000)
	httpServer.ListenAndServe()
}
