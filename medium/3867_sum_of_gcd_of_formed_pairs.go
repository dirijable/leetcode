package medium

import "slices"

func gcdSum(nums []int) int64 {
	prefixGCD := make([]int, 0, len(nums))
	greatest := nums[0]
	for _, v := range nums {
		greatest = max(greatest, v)
		currentGcd := gcd(greatest, v)
		prefixGCD = append(prefixGCD, currentGcd)
	}
	slices.Sort(prefixGCD)
	var sum int64
	for left, right := 0, len(prefixGCD)-1; left < right; left, right = left+1, right-1 {
		sum += int64(gcd(prefixGCD[right], prefixGCD[left]))
	}
	return sum
}

func gcd(greatest, least int) int {
	for least > 0 {
		least, greatest = greatest%least, least
	}
	return greatest
}
