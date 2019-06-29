package main

import (
	"fmt"
	"strconv"
)

//board sample case
var board = [][]string{{"1", "1", "1", "1", "0"}, {"1", "1", "0", "1", "0"}, {"1", "1", "0", "0", "0"}, {"0", "0", "0", "0", "0"}}

func main() {
	grid := changeMap(board)
	p(grid)
	fmt.Println(numIslands(grid))
	p(grid)
}

//movement up, right, down, and left
var move = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

//implementation using DFS
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	numIsland := 0
	for r, _ := range grid {
		for c, _ := range grid[0] {
			if int(grid[r][c]) != 1 {
				continue
			}
			numIsland++
			dfs(r, c, grid)
		}
	}

	return numIsland
}

//dfs implementation
func dfs(row, col int, grid [][]byte) {
	if !inRange(row, col, grid) || int(grid[row][col]) != 1 {
		return
	}
	grid[row][col] = byte(5)

	for _, v := range move {
		dRow := row + v[0]
		dCol := col + v[1]
		dfs(dRow, dCol, grid)
	}
}

//check whether location inside the grid
func inRange(r, c int, grid [][]byte) bool {
	return r >= 0 && c >= 0 && r < len(grid) && c < len(grid[0])
}

//chgen string grid to byte
func changeMap(grid [][]string) [][]byte {
	row := len(grid)
	col := len(grid[0])
	m := make([][]byte, row)

	for i := 0; i < row; i++ {
		m[i] = make([]byte, col)
	}

	for r, _ := range grid {
		for c, _ := range grid[0] {
			v, _ := strconv.Atoi(grid[r][c])
			m[r][c] = byte(v)
		}
	}
	return m
}

//print grid function
func p(grid [][]byte) {
	for r, _ := range grid {
		for c, _ := range grid[0] {
			fmt.Print(int(grid[r][c]), " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
