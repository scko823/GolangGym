package main

import "fmt"

// TimeConvert (num) take the num parameter being passed and return the number of hours and minutes the parameter converts to
//(ie. if num = 63 then the output should be 1:3).
//Separate the number of hours and minutes with a colon.
func TimeConvert(num int) string {
	return fmt.Sprintf("%v:%v", num/60, num%60)
}

func main() {
	fmt.Println(TimeConvert(63))
}
