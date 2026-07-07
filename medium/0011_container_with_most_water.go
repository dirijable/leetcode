package medium

func maxArea(height []int) int {
	maxSquare := 0
	for left, right := 0, len(height)-1; left < right; {
		maxSquare = max(maxSquare, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxSquare
}
