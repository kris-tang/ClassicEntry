package main

import "fmt"

type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) area() float64 {
	return (float64(1) / float64(2)) * t.base * t.height
}

func main() {
	t := Triangle{base: 4, height: 5}
	fmt.Println(t.area())
}
