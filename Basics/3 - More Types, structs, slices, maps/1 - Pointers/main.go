package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println("Value of pointer p is :", p)
	fmt.Println("Value of i through p is :", *p)
	*p = 21
	fmt.Println("Value of i through i is:", i)
	fmt.Println(i)

	p = &j
	fmt.Println("Value of pointer p is :", p)
	fmt.Println("Value of j through p is :", *p)
	*p = *p / 37
	fmt.Println("Value of j through j is:", i)
}
