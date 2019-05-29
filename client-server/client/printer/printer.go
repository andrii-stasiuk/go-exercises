package printer

import (
	"fmt"
	"log"
)

// MapPrinter function for printing maps
func MapPrinter(stringMap map[string]interface{}, err error) {
	if err == nil {
		fmt.Println()
		for key, value := range stringMap {
			fmt.Println("Key:", key, "Value:", value)
		}
		fmt.Println()
	} else {
		log.Println(err)
	}
}
