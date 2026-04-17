package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespError represents the error response
type ErrorAPI struct {
	Error string `json:"error"`
}

// JSON returns a JSON response to the client
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// StatusCodeErrorHandled returns a JSON response to the client with the error message
func StatusCodeErrorHandled(w http.ResponseWriter, r *http.Response) {
	var errorResp ErrorAPI
	json.NewDecoder(r.Body).Decode(&errorResp)
	JSON(w, r.StatusCode, errorResp)
}
