package main

import "fmt"

func anotherFunction(f func() string) string {
	return f()
}

func main() {
	fn := func() string {
		return "function called1"
	}
	fmt.Println(anotherFunction(fn))
}
