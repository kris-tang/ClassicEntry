package main

import (
	"fmt"
	"time"
)

func slowSleep() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished")
}

func main() {
	go slowSleep()
	fmt.Println("I am not shown slowSleep() completes")
	time.Sleep(time.Second * 3)
}
