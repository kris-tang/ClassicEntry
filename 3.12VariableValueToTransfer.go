package main

import "fmt"

func showMemoryAddress(x *string) {
	fmt.Println(x)
	return
}

func main() {
	i := "hello"
	fmt.Println(&i)
	showMemoryAddress(&i)
}
