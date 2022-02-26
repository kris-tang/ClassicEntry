package main

import "fmt"

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "tang"
	cheeses[1] = "wei"
	cheeses[2] = "jie"
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
}
