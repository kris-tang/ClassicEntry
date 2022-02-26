package main

import "fmt"

func sayHi() (x, y, z string) {
	x = "Hello"
	y = "world"
	z = "8888"
	return
}

func main() {
	fmt.Println(sayHi())
}
