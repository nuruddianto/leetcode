package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	i := 1
	start := "1"
	if n == 1 {
		return start
	}
	res := ""
	for i < n {
		curr := start[0]
		count := 1
		for j := 1; j < len(start); j++ {
			if curr != start[j] {
				res = res + strconv.Itoa(count) + string(curr)
				curr = start[j]
				count = 1
				continue
			}
			count++
		}
		start = res + strconv.Itoa(count) + string(curr)
		res = ""
		i++
	}
	return start
}
