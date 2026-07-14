package medium

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	left, right := 0, n*m-1
	for left <= right {
		mid := (left + right + 1) / 2
		r, c := mid/m, mid%m
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
