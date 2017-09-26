package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is leng - number of elements reffered to by the slice
	// 5 is capacity - number of elements in the undelying array

	greeting[0] = "Good maorning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Buenos dias!"
	greeting[3] = "Suprabadham"

	fmt.Println(greeting[2])
}
