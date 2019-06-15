package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	for c := 0; c < 9; c++ {
		if !checkRow(c, board) {
			return false
		}
	}

	for r := 0; r < 9; r++ {
		if !checkCol(r, board) {
			return false
		}
	}

	for r := 0; r < 9; r += 3 {
		for c := 0; c < 9; c += 3 {
			if !checkBox(r, c, board) {
				return false
			}
		}
	}

	return true
}

func checkRow(col int, board [][]byte) bool {
	m := make(map[int]bool)
	for r := 0; r < 9; r++ {
		val := int(board[r][col])
		if m[val] {
			return false
		}
		m[val] = true
	}
	return true
}

func checkCol(row int, board [][]byte) bool {
	m := make(map[int]bool)
	for c := 0; c < 9; c++ {
		val := int(board[row][c])
		if m[val] {
			return false
		}
		m[val] = true
	}
	return true
}

func checkBox(row, col int, board [][]byte) bool {
	m := make(map[int]bool)
	for r := row; r < row+3; r++ {
		for c := col; c < col+3; c++ {
			val := int(board[r][c])
			if m[val] {
				return false
			}
			m[val] = true
		}
	}
	return true
}
