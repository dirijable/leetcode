package hard

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]
	water := 0
	for left < right {
		if height[right] < height[left] {
			right--
			maxRight = max(maxRight, height[right])
			water += maxRight - height[right]
		} else {
			left++
			maxLeft = max(maxLeft, height[left])
			water += maxLeft - height[left]
		}
	}
	return water
}