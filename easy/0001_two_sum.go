package easy

func twoSum(nums []int, target int) []int {
    numsAndIdx := make(map[int]int, len(nums))
	for idx, num := range nums {
		if i, ok := numsAndIdx[target - num]; ok {
			return []int{idx, i}
		}
		numsAndIdx[num] = idx
	}
	return []int{}
}