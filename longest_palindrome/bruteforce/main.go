package main

import "fmt"

func main() {
	input := "cbbc"
	fmt.Println(longestPalindrome(input))
}

var m map[string]int

func longestPalindrome(s string) string {
	var curr string
	var ans string
	for start := range s {
		for i := start + 1; i <= len(s); i++ {
			curr = s[start:i]

			res := palindrome(curr)
			if len(res) > len(ans) {
				ans = res
			}
		}
	}
	return ans
}

func palindrome(subs string) string {
	n := len(subs)
	if n == 0 {
		return ""
	}
	j := n - 1
	i := 0

	for i <= j {
		if subs[i] != subs[j] {
			return ""
		}
		i++
		j--
	}
	return subs
}
