package inc

import "fmt"

// MapPrinter function for printing maps
func MapPrinter(stringMap map[string]interface{}) {
	for key, value := range stringMap {
		fmt.Println("Key:", key, "Value:", value)
	}
}
