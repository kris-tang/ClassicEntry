package main

import (
	"errors"
	"fmt"
)

type Robot interface {
	Poweron() error
}

type T850 struct {
	Name string
}

func (t *T850) Poweron() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) Poweron() error {
	if r.Broken {
		return errors.New("R2D2 Robot is broken")
	} else {
		return nil
	}
}

func Boot(r Robot) error {
	return r.Poweron()
}

func main() {
	t := T850{
		Name: "The Terminator",
	}

	r := R2D2{
		Broken: true,
	}
	fmt.Println(t.Name)
	fmt.Println(r.Broken)

	err := Boot(&r)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The R2D2 Robot is Power on!")
	}

	err = Boot(&t)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The T850 Robot is Power on!")
	}
}
