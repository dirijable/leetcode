package hard

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0, len(nums)-k+1)
	deque := make([]int, 0)
	for i := range nums {
		deque = append(deque, i)
		for len(deque) > 0 && nums[i] >= nums[deque[0]] {
			deque = deque[:len(deque) - 1]
		}
		for len(deque) >0 && deque[0] < i - k + 1 {
			deque = deque[1:]
		}
		result = append(result, nums[deque[0]])
	}
	return result
}

// import "container/heap"

// type Pair struct {
// 	idx   int
// 	value int
// }

// type PriotityQueue []Pair

// func (q *PriotityQueue) Len() int {
// 	return len(*q)
// }

// func (q *PriotityQueue) Less(i, j int) bool {
// 	return (*q)[i].value > (*q)[j].value
// }

// func (q *PriotityQueue) Swap(i, j int) {
// 	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
// }

// func (q *PriotityQueue) Push(x interface{}) {
// 	*q = append(*q, x.(Pair))
// }

// func (q *PriotityQueue) Pop() any {
// 	first := (*q)[len(*q)-1]
// 	*q = (*q)[:len(*q)-1]
// 	return first
// }

// func maxSlidingWindow(nums []int, k int) []int {
// 	var pq PriotityQueue
// 	heap.Init(&pq)
// 	result := make([]int, 0, len(nums)-k+1)
// 	for i := range k {
// 		heap.Push(&pq, Pair{value: nums[i], idx: i})
// 	}
// 	result = append(result, pq[0].value)
// 	for left, right := 1, k; right < len(nums); right, left = right+1, left+1 {
// 		for pq[0].idx < left {
// 			heap.Pop(&pq)
// 		}
// 		heap.Push(&pq, Pair{value: nums[right], idx: right})
// 		result = append(result, pq[0].value)
// 	}
// 	return result
// }
