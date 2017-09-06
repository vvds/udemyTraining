package main

import "fmt"

var x int

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by ultiple functons
without closure, for two or more funcs to have acess to the same variable
that variable would need to be package scope
*/
