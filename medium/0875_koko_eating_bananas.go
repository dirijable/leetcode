package medium

import (
	"slices"
)

func minEatingSpeed(piles []int, h int) int {
	low, high := 1, slices.Max(piles)
	var minSpeed int
	for low <= high {
		mid := (low + high + 1) / 2
		if canEatAll(mid, h, piles) {
			minSpeed = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return minSpeed
}

func canEatAll(speed, hours int, piles []int) bool {
	for _, v := range piles {
		hours -= (v + speed - 1) / speed
		if hours < 0 {
			return false
		}
	}
	return true
}

/*
как я понял у нас надо искать из другого диапазона значений.
нижний это видимо 1, ну а большой - самая большая.
начинаем снова с середины, короче как в бинарном поиске, но при выборе числа для ответа начинаем смотреть хватитли времени,
если хватает, значит,можно попробовать взять меньше сохраняя самое маленькое, а если меньше скорости не хватает,
то и дальше меньшие смотреть нет смысла потому что для еще меньших точно не хватит.
и как понял сложность будет O(nlogm) , где n - число куч, а m - число значений в диапазоне
*/
