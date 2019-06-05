/*Package responses (json content)*/
package responses

import (
	"encoding/json"
	"net/http"
)

type JsonErrorResponse struct {
	Error *ApiError `json:"error"`
}

type ApiError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

// WriteOKResponse - writes the response as a standard JSON response with StatusOK
func WriteOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&m); err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// WriteErrorResponse - writes the error response as a Standard API JSON response with a response code
func WriteErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	if err := json.NewEncoder(w).Encode(&JsonErrorResponse{Error: &ApiError{Status: errorCode, Title: errorMsg}}); err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}
