package api

import (
	"fmt"
	"net/http"
	"nginx_debugger/api/endpoints/analyze"
)

type HttpServer struct {
}

func NewHttpServer(host string, port int) *http.Server {
	serveMux := http.NewServeMux()

	serveMux.Handle("/analyze", analyze.NewAnalyzeNginxConfigEndpointHandler())

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: serveMux,
	}
}
