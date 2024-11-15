package http

import (
	"encoding/json"
	"net/http"
)

type Handler func(r *http.Request) (statusCode int, data map[string]interface{})

func (h Handler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	statusCode, data := h(r)
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}
