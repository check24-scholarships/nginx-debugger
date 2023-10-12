package endpoints

import "net/http"

type AnalyzeNginxConfigEndpointHandler struct {
}

func NewAnalyzeNginxConfigEndpointHandler() http.Handler {
	return &AnalyzeNginxConfigEndpointHandler{}
}

func (h *AnalyzeNginxConfigEndpointHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
