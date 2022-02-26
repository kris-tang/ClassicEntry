package main

import "fmt"

func main() {
	s := "c"

	switch s {
	case "a":
		fmt.Println("The letters a!")
	case "b":
		fmt.Println("The letters b!")
	default:
		fmt.Println("I don't recognize that letters!")
	}
}
