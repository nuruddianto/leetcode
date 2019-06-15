package main

import (
	"fmt"
	"unicode"
)

func main() {
	in := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(in))
}
func isPalindrome(s string) bool {
	var start int
	finish := len(s) - 1

	for start < finish {
		r1 := rune(s[start])
		r2 := rune(s[finish])
		//fmt.Println(r1, r2)
		if !isAlphanumeric(r1) {
			start++
			continue
		}

		if !isAlphanumeric(r2) {
			finish--
			continue
		}

		if !isSame(r1, r2) {
			return false
		}
		start++
		finish--
	}
	return true
}

func isAlphanumeric(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsNumber(s)
}

func isSame(r1, r2 rune) bool {
	r1 = unicode.ToLower(r1)
	r2 = unicode.ToLower(r2)
	return r1 == r2
}
