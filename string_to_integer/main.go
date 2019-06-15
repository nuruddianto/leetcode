package main

import (
	"fmt"
	"math"
)

func main() {
	in := "9223372036854775808"
	fmt.Println(myAtoi(in))
}

func myAtoi(str string) int {
	var res int
	var isNeg bool
	var isPlus bool
	var isStart bool

	for _, v := range str {
		if int(v) == 45 && !isNeg {
			if isPlus {
				return reformat(res, isNeg)
			}
			isNeg = true
			continue
		}

		if int(v) == 43 && !isPlus {
			if isNeg {
				return reformat(res, isNeg)
			}
			isPlus = true
			continue
		}
		//continue if space
		if int(v) == 32 {
			if isStart {
				return reformat(res, isNeg)
			}
			continue
		}

		if !isNumber(v) {
			return reformat(res, isNeg)
		}
		if res == 0 {
			isStart = true
			res = toNumber(v)
		} else {
			res = res*10 + toNumber(v)
		}
	}
	return reformat(res, isNeg)
}

func isNumber(v rune) bool {
	return int(v) >= 48 && int(v) <= 57
}

func toNumber(v rune) int {
	return int(v) - 48
}

func reformat(res int, isNeg bool) int {
	fmt.Println(res)
	if isNeg {
		res *= -1
	}
	max := int(math.Pow(2, 31)) - 1
	min := (max + 1) * (-1)

	if res >= max {
		return max
	} else if res <= min {
		return min
	}

	return res
}
