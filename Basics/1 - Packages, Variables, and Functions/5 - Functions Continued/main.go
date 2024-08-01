package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Sum of 2 and 3 is :", add(2, 3))
}
