package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pingpong(ch <-chan string) {
	time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(900)))
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
