package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)
	for {

		go func() {
			amt := time.Duration(100 + rand.Intn(900))
			time.Sleep(time.Millisecond * amt)
			ch <- "ping"
		}()
		fmt.Println(<-ch)

		go func() {
			amt := time.Duration(100 + rand.Intn(900))
			time.Sleep(time.Millisecond * amt)
			ch <- "pong"
		}()
		fmt.Println(<-ch)

	}

	// var input string
	// fmt.Scanln(&input)
}
