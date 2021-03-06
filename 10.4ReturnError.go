package main

import "fmt"

func Half(numberToHalf int) (int, error) {
	if numberToHalf%2 != 0 {
		return -1, fmt.Errorf("Cannot half %v", numberToHalf)
	}
	return numberToHalf / 2, nil
}

func main() {
	n, err := Half(36)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
