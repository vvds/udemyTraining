package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
hat variable would need to be package scope

anonymous functions
a function without a name

func expression
assing a func to a variable
*/
