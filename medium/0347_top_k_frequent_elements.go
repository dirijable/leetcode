package medium

import (
	"slices"
)

type KeyValue struct {
	key   int
	value int
}

func topKFrequent(nums []int, k int) []int {
	elementFreq := map[int]int{}
	for _, num := range nums {
		elementFreq[num]++
	}
	keyValue := make([]KeyValue, len(elementFreq))
	for k, v := range elementFreq {
		keyValue = append(keyValue, KeyValue{key: k, value: v})
	}

	slices.SortFunc(keyValue, func(first, second KeyValue) int {
		return second.value - first.value
	})
	result := make([]int, 0, k)
	for i := range k {
		result = append(result, keyValue[i].key)
	}
	return result
}
