package main

import "fmt"

func main() {
	// data := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// data := [][]int{
	// 	{1, 2},
	// 	{3, 4},
	// }

	data := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	print(data)
	print(rotate(data))
}

func rotate(data [][]int) [][]int {
	m := len(data) //total row
	n := len(data) //total col
	row := 0       //start row
	col := 0       //start col

	for row < m-1 && col < n-1 {
		prev := data[row+1][col]

		//substitute top row
		for c := col; c < n; c++ {
			tmp := data[row][c]
			data[row][c] = prev
			prev = tmp
		}
		row++

		//substitute right col
		for r := row; r < m; r++ {
			tmp := data[r][n-1]
			data[r][n-1] = prev
			prev = tmp
		}
		n--

		if row < m {
			//substitute bottom row
			for c := n - 1; c >= col; c-- {
				tmp := data[m-1][c]
				data[m-1][c] = prev
				prev = tmp
			}
			m--
		}

		if col < n {
			//substitute left col
			for r := m - 1; r >= row-1; r-- {
				tmp := data[r][col]
				data[r][col] = prev
				prev = tmp
			}
		}
		col++
		// print(data)
		// fmt.Print()
	}
	return data
}

func print(d [][]int) {
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d[i]); j++ {
			fmt.Print(d[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
