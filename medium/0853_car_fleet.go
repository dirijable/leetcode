package medium

import "slices"

type Car struct {
	position int
	time     float64
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, 0, len(position))
	for i := range len(position) {
		time := float64((target - position[i])) / float64(speed[i])
		cars = append(cars, Car{time: time, position: position[i]})
	}
	slices.SortFunc(cars, func(c1, c2 Car) int {
		return c2.position - c1.position
	})
	timeStack := make([]float64, 0, len(position))
	for _, car := range cars {
		if len(timeStack) == 0 || car.time > timeStack[len(timeStack)-1] {
			timeStack = append(timeStack, car.time)
		}
	}
	return len(timeStack)
}
