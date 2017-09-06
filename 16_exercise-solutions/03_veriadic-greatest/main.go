package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	greatest := max(4, 7, 9, 123, 543, 53, 125)
	fmt.Println(greatest)
}

/*
FYI
for your code to also work with only negative numbers such as

greatest := max(-200 -700)

include this as your range statement
for i, v := range numbers {
        if v > largest || i == 0 {
        largest = v
      }
}

what does that code do?

the first time through the range loop
the index, it will be zero
so largest will be set to the first numbers

Originally largest is set to the zero value first an int, which is zero

Zero would be greater than any negative numbers

if you only have negative numbers
you need largest to be something less than zero

thanks to Ricardo G for this code improvement
*/
