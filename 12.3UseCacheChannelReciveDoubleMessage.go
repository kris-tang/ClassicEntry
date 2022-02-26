package main

import (
	"fmt"
	"time"
)

func receivers(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	messages := make(chan string, 2)
	messages <- "Hello"
	messages <- "world"
	close(messages)
	fmt.Println("Pushed two messages on Channel with no receivers")
	time.Sleep(time.Second * 1)
	receivers(messages)
}
