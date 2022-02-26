package main

import "fmt"

func main() {
	var players = make(map[string]int)
	players["tang"] = 99
	players["wang"] = 88
	players["zhao"] = 66
	fmt.Println(players["tang"])
	fmt.Println(players["wang"])
	fmt.Println(players["zhao"])
}
