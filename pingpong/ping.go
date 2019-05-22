package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ping(ch chan string) {
	amt := time.Duration(100 + rand.Intn(900))
	time.Sleep(time.Millisecond * amt)
	ch <- "ping"
}

func pong(ch chan string) {
	amt := time.Duration(100 + rand.Intn(900))
	time.Sleep(time.Millisecond * amt)
	ch <- "pong"
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)
	for {
		go ping(ch)
		fmt.Println(<-ch)
		go pong(ch)
		fmt.Println(<-ch)
	}
}
