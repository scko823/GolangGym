package main

/*
Have the function LetterChanges(str) take the str parameter being passed and modify it using the following algorithm.
 Replace every letter in the string with the letter following it in the alphabet (ie. c becomes d, z becomes a).
Then capitalize every vowel in this new string (a, e, i, o, u)
 and finally return this modified string.

*/

import (
	"fmt"
	"strings"
)

// func split(s string) []string {
// 	return strings.Split(s, "")
// }

func shiftOne(s string) string {
	r := int([]rune(s)[0])
	if r <= 64 || (r >= 91 && r <= 96) || r >= 123 {
		return s
	}
	switch r {
	case 90:
		r = 65
	case 122:
		r = 97
	default:
		r = r + 1
	}
	return string(r)
}

func vowelToUpper(c string) string {
	if strings.Contains("aeiou", c) {
		return strings.ToUpper(c)
	}
	return c
}

// LetterChange wiring all together
func LetterChange(s string) string {
	results := ""
	for _, c := range s {
		results += vowelToUpper(shiftOne(string(c)))
	}
	return results
}

func main() {

	str := "this long cake@&"

	fmt.Println(LetterChange(str))
}
