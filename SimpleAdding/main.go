package main

import "fmt"

//SimpleAdding (num) add up all the numbers from 1 to num.
// For example: if the input is 4 then your program should return 10 because 1 + 2 + 3 + 4 = 10.
// For the test cases, the parameter num will be any number from 1 to 1000.
//
func SimpleAdding(n int) int {
	if n <= 0 {
		return -1
	}
	result := int(0)
	for i := int(1); i <= n; i++ {
		result += i
	}
	return result
}

func main() {
	fmt.Println(SimpleAdding(1000))
}
