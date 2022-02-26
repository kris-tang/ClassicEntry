package main

import "fmt"

type Superhero struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Number int
	Street string
	City   string
}

func main() {
	e := Superhero{
		Name: "tangweijie",
		Age:  32,
		Address: Address{
			Number: 1007,
			Street: "Mountain Drive",
			City:   "Gotham",
		},
	}
	fmt.Println("%+v\n", e)
}
