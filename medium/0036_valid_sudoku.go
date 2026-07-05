package medium

func isValidSudoku(board [][]byte) bool {
	var rows, cols, blocks [9][9]bool
	for r := range 9 {
		for c := range 9 {
			if board[r][c] == '.' {
				continue
			}
			rcbIdx := int(board[r][c]-'0') - 1
			blockIdx := r / 3 * 3 + c/3
			if rows[r][rcbIdx] || cols[c][rcbIdx] || blocks[blockIdx][rcbIdx] {
				return false
			}
			rows[r][rcbIdx] = true
			cols[c][rcbIdx] = true
			blocks[blockIdx][rcbIdx] = true
		}
	}
	return true
}