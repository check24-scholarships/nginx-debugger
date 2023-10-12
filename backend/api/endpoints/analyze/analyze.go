package analyze

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nginx_debugger/explain"
	"nginx_debugger/internal/parser"
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

	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	rawCfg := string(bytes)
	configParser := parser.NewParser(rawCfg)
	parsedCfg, err := configParser.Parse()
	if err != nil {
		http.Error(writer, fmt.Sprintf("parser error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	domainResponse := explain.ExplainNginxConfig(*parsedCfg)

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(domainResponse)
}
