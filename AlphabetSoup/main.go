package main

import (
	"fmt"
	"sort"
	"strings"
)

// AlphabetSoup (str) take the str string parameter being passed and return the string with the letters in alphabetical order (ie. hello becomes ehllo).
// Assume numbers and punctuation symbols will not be included in the string
func AlphabetSoup(s string) string {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}

func main() {
	fmt.Println(AlphabetSoup("quack"))
}
