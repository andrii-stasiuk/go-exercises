/*Package errors for error handling*/
package errors

import "fmt"

// DatabaseErrors structure type
type DatabaseErrors struct {
	ErrorID   uint64
	ErrorText string
}

// Application error handling method
func (e DatabaseErrors) Error() string {
	return fmt.Sprintf("An error occurred (ID #%d). %s", e.ErrorID, e.ErrorText)
}
