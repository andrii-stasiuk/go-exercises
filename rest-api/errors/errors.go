/*Package errors for error handling*/
package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func jsonResponse(res http.ResponseWriter, data interface{}) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")

	payload, err := json.Marshal(data)
	if ErrorCheck(res, err) {
		return
	}

	fmt.Fprintf(res, string(payload))
}

func ErrorCheck(res http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

type Error struct{}

func (e Error) notFoundError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprintf(w, "404! Not found")
}

// type JsonResponse struct {
// 	// Reserved field to add some meta information to the API response
// 	Meta interface{} `json:"meta"`
// 	Data interface{} `json:"data"`
// }

type JsonErrorResponse struct {
	Error *ApiError `json:"error"`
}

type ApiError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

// Writes the response as a standard JSON response with StatusOK
func WriteOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&m); err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// Writes the error response as a Standard API JSON response with a response code
func WriteErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).
		Encode(&JsonErrorResponse{Error: &ApiError{Status: errorCode, Title: errorMsg}})
}
