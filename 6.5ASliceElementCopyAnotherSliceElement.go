package main

import "fmt"

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "tang"
	cheeses[1] = "wei"
	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, cheeses)
	fmt.Println(cheeses)
	fmt.Println(smellyCheeses)
}
