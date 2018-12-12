package common

import (
	"encoding/json"
	"net/http"
)

// JSONResponse sends response as JSON
func JSONResponse(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "an unpexected error occured encoding response"}`))
		return
	}
}

// BaseError is used to respond an error
type BaseError struct {
	Message string `json:"message"`
}
