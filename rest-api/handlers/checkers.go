/*Package handlers Todo*/
package handlers

import "strconv"

// CheckInt function for basic verification of numbers, can be extended in the future
func CheckInt(id string) bool {
	converted, err := strconv.ParseUint(id, 10, 64)
	if err == nil && converted > 0 {
		return true
	}
	return false
}

// CheckStr function for basic string checking, can be extended in the future
func CheckStr(str string) bool {
	if len(str) > 0 && str != "`" {
		return true
	}
	return false
}
