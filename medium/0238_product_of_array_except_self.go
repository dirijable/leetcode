package medium

func productExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, len(nums))
	prefix := 1
	for i := range length {
		result[i] = prefix
		prefix *= nums[i]
	}
	// 1 2 3 4
	// 1 1 2 6

	suffix := 1
	for i := length - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}
	return result
}


//
// prefix := make([]int, length)
// prefix[0] = nums[0]
// postfix := make([]int, length)
// postfix[length-1] = nums[length-1]
// for start, end := 1, length-2; start < length && end >= 0; start, end = start+1, end-1 {
// 	prefix[start] = nums[start-1] * nums[start]
// 	postfix[end] = nums[end+1] * nums[end]
// }
// result := make([]int, length)
// result[0] = postfix[1]
// result[length-1] = prefix[length-2]
// for i := 1; i < length-1; i++ {
// 	result[i] = prefix[i-1] * postfix[i+1]
// }
// return result
