package easy

import "slices"

func containsDuplicate(nums []int) bool {
	slices.Sort(nums)
	for i := range len(nums) - 1 {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// func containsDuplicate(nums []int) bool {
// 	freq := make(map[int]int, len(nums))
// 	for _, num := range nums {
//         freq[num]++
// 		if freq[num] > 1 {
// 			return true
// 		}
// 	}
// 	return false
// }

// func containsDuplicate(nums []int) bool {
//     set := make(map[int]struct{}, len(nums))
//     for _, num := range nums {
//         if _, ok := set[num]; ok {
//             return true
//         }
//         set[num] = struct{}{}
//     }
//     return false
// }