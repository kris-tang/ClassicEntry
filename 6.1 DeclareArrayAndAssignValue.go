package main

import "fmt"

func main() {
	var cheeses [2]string
	cheeses[0] = "Marilla"
	cheeses[1] = "Epoisses de Bourgogne"
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	fmt.Println(cheeses)
}
