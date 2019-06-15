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
	rotate(data)
}

func rotate(data [][]int) [][]int {

	//transpose matrix
	for r := 0; r < len(data); r++ {
		for c := r; c < len(data); c++ {
			tmp := data[r][c]
			data[r][c] = data[c][r]
			data[c][r] = tmp
		}
	}

	//reverse column
	for r := 0; r < len(data); r++ {
		for c := 0; c < len(data)/2; c++ {
			tmp := data[r][c]
			data[r][c] = data[r][len(data)-1-c]
			data[r][len(data)-1-c] = tmp
		}
	}
	return data
}
func reverse() {

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
