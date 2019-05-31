/*Package errors for error handling*/
package errors

import "net/http"

func ErrorCheck(res http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}
