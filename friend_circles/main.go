package main

import "fmt"

func main() {
	M := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(findCircleNum(M))
}

var moves = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}

	friend := make([]bool, len(M))
	circle := 0

	for r := 0; r < len(M); r++ {
		if !friend[r] {
			circle++
			doDfs(r, M, friend)
		}
	}
	return circle
}

func doDfs(r int, M [][]int, friend []bool) {
	friend[r] = true
	for c := 0; c < len(M[r]); c++ {
		if M[r][c] == 1 && !friend[c] {
			doDfs(c, M, friend)
		}
	}
}
