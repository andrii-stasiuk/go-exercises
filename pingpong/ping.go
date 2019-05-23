package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pingpong(ch <-chan string) {
	amt := time.Duration(100 + rand.Intn(900))
	time.Sleep(time.Millisecond * amt)
	fmt.Println(<-ch)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)
	for {
		go pingpong(ch)
		ch <- "ping"
		go pingpong(ch)
		ch <- "pong"
	}
}
