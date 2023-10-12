package analyze

import (
	"encoding/json"
	"net/http"
)

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

	var domainRequest Request
	err := json.NewDecoder(request.Body).Decode(&domainRequest)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	//Domain logic
	response := Response{
		Explanation: map[int]string{0: "Hello, World!"},
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
