package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // address

	var b = &a
	fmt.Println(b)  // same address
	fmt.Println(*b) // 43

	// b is an int (interger) pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is nown as dereferencing
	// the * is an operator in this case
}
