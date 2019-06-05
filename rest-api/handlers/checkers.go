package handlers

import "strconv"

// CheckInt basic check of integer
func CheckInt(id string) bool {
	converted, err := strconv.ParseUint(id, 10, 64)
	if err == nil && converted > 0 {
		return true
	}
	return false
}

// CheckStr basic check of string
func CheckStr(str string) bool {
	if len(str) > 0 && str != "`" {
		return true
	}
	return false
}
