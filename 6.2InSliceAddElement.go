package main

import "fmt"

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "tang"
	cheeses[1] = "wei"
	cheeses = append(cheeses, "jie")
	fmt.Println(cheeses[2])
	fmt.Println(cheeses[1])
}
