package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // []
	changeMe(m)
	fmt.Println(m) // [Oscar]
}

func changeMe(z []string) {
	z[0] = "Oscar"
	fmt.Println(z) // [Oscar]
}
