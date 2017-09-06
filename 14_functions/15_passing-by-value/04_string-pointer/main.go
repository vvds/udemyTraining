package main

import "fmt"

func main() {

	name := "Oscar"
	fmt.Println(&name) // 0x82023c080

	changeMe(&name)

	fmt.Println(&name) // 0x82023c080
	fmt.Println(name)  // Rocky
}

func chengeMe(z *string) {
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Todd
	*z = "Rocky"
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Rocky
}
