package easy

func gcdOfOddEvenSums(n int) int {
	oddSum, evenSum := n*n, n*(n+1)
	remainder := evenSum % oddSum
	lastRemainder := remainder
	for remainder > 0 {
		lastRemainder = remainder
		remainder = oddSum % remainder

	}
	return lastRemainder
}
