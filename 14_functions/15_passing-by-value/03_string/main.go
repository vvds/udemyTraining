package main

import "fmt"

func main() {

	name := "Oscar"
	fmt.Println(name) // Oscar

	changeMe(name)

	fmt.Println(name) // Oscar
}

func changeMe(z string) {
	fmt.Println(z) // Oscar
	z = "Rocky"
	fmt.Println(z) // Rocky
}
