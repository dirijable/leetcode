package medium

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0, len(nums))
	slices.Sort(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, len(nums)-1; left < right && right > i; {
			if nums[left]+nums[right] < -nums[i] {
				left++
			}
			if right == i || nums[left]+nums[right] > -nums[i] {
				right--
			}
			if left < right && nums[left]+nums[right]+nums[i] == 0 {
				result = append(result, []int{nums[left], nums[right], nums[i]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
				left++
			}
		}
	}
	return result
}
