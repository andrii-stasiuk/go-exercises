package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ping(ch chan<- string) {
	for {
		ch <- "ping"
	}
}

func pong(ch chan<- string) {
	for {
		ch <- "pong"
	}
}

func printer(ch <-chan string) {
	for {
		time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(900)))
		fmt.Println(<-ch)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)
	go ping(ch)
	go pong(ch)
	go printer(ch)
	for {
	}
}
