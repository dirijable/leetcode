package hard

type Rectangle struct {
	height int
	idx    int
}

func largestRectangleArea(heights []int) int {
	stack := make([]Rectangle, 0, len(heights))
	heights = append(heights, 0)
	maxSquare := -1
	for i, h := range heights {
		topIdx := i
		for len(stack) > 0 && stack[len(stack)-1].height > h {
			rect := stack[len(stack)-1]
			maxSquare = max(maxSquare, rect.height * (i - rect.idx))
			topIdx = rect.idx
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, Rectangle{height: h, idx: topIdx})
	}
	return maxSquare
}
