package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	t := time.NewTicker(time.Second * 1)
	for {
		c <- "ping"
		<-t.C
	}
}

func main() {
	messages := make(chan string)
	go pinger(messages)
	msg := <-messages
	fmt.Println(msg)
}
