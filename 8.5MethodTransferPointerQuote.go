package main

import "fmt"

type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) changeBase(f float64) {
	t.base = f
	return
}

func main() {
	t := Triangle{base: 10, height: 20}
	t.changeBase(5)
	fmt.Println(t.base)
}
