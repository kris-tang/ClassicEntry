package main

import "fmt"

func main() {
	var players = make(map[string]int)
	players["tang"] = 100
	players["wang"] = 110
	players["zhao"] = 999
	delete(players, "zhao")
	fmt.Println(players)
}
