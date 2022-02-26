package main

import "fmt"

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	if a == b {
		fmt.Printf("a and b are the same\n")
	}
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
}
