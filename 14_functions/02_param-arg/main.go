package main

import "fmt"

func main() {
	greet("Jane")
	greet("John")
}

func greet(name string) {
	fmt.Println(name)
}

// greet is declaed with a parameter
// when calling greet, pass an argument
