package main

import "fmt"

func main() {
	in1 := "mississippi"
	in2 := "issipi"
	fmt.Println(strStr(in1, in2))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}
	j := 0
	for i, v := range haystack {
		if byte(v) == needle[j] {
			for k := j; k < len(needle); k++ {
				if i+k >= len(haystack) || haystack[i+k] != needle[k] {
					break
				}
				if k == len(needle)-1 {
					return i
				}
			}
		}
	}
	return -1
}
