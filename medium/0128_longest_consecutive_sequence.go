package medium

func longestConsecutive(nums []int) int {
	allNumbers := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		allNumbers[n] = struct{}{}
	}
	resultCount := 0
	for num := range allNumbers {
		if _, ok := allNumbers[num-1]; ok  {
			continue
		}
		currentCount := 1 
		for i:= num + 1 ; ;i++  {
			if _, ok := allNumbers[i]; !ok {
				break
			}
			currentCount++
		}
		if currentCount > resultCount {
			resultCount = currentCount
		}

	}
	return resultCount
}
