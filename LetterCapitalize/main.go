package main

import (
	"fmt"
	"strings"
)

//LetterCapitalize (str) take the str parameter being passed and capitalize the first letter of each word.
// Words will be separated by only one space.
func LetterCapitalize(s string) string {
	result := ""
	words := strings.Split(s, " ")
	for _, word := range words {
		chars := strings.Split(word, "")
		chars[0] = strings.ToUpper(chars[0])
		result += strings.Join(chars, "")
		result += " "
	}
	return result
}

func main() {
	fmt.Println(LetterCapitalize("hello my World! I am gopher"))
}
