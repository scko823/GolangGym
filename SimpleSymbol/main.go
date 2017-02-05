package main

import "fmt"

// SimpleSymbols (str) take the str parameter being passed and determine if it is an acceptable sequence by either returning the string true or false.
// The str parameter will be composed of + and = symbols with several letters between them (ie. ++d+===+c++==a)
// and for the string to be true each letter must be surrounded by a + symbol.
// So the string to the left would be false. The string will not be empty and will have at least one letter.
func SimpleSymbols(s string) bool {
	if len(s) < 3 {
		return false
	}
	if string(s[0]) != "+" || string(s[len(s)-1]) != "+" {
		return false
	}
	for i, r := range s {
		if i == 0 || i == len(s)-1 {
			continue
		}
		if r >= 97 || r <= 122 {
			prev := string(s[i-1])
			next := string(s[i+1])
			if prev != "+" || next != "+" {
				return false
			}
		}
		return true
	}
	return true
}

func main() {
	fmt.Println(SimpleSymbols("2+a+a+"))
}
