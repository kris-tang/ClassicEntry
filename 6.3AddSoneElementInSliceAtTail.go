package main

import "fmt"

func main() {
	var chesses = make([]string, 2)
	chesses[0] = "tang"
	chesses[1] = "wei"
	chesses = append(chesses, "jie", "wang", "jun", "lan")
	fmt.Println(chesses)
}
