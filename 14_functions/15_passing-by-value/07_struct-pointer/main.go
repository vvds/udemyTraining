package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c1 := customer{"Oscar", 31}
	fmt.Println(&c1.name) // 0x82023c080

	changeMe(&c1)

	fmt.Println(c1)       // {Rocky 44}
	fmt.Println(&c1.name) // 0x82023c080
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Oscar 31}
	fmt.Println(&z.name) // 0x82023c080
	z.name = "Rocky"
	fmt.Println(z)       // &{Rocky 44}
	fmt.Println(&z.name) // 0x82023c080
}
