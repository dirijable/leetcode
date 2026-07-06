package medium

/*
короче придумал такое: заводим стэк для более теплой температуры, в него добавляем температуру.
можем начинать всегда с предпоследнего элемента, так как для последнего всегда 0 будет в конце.
когда мы берем очередную температуру, то смотрим в стэк более теплых температур: если эта температура меньше, просто удаляем ее;
если она больше, то записываем разницу индексов, но не удаляем, так как могут быть те температуры, которым тоже нужна эта температура.
вот только проблема в том, что для запоминания индексов еще одну структуру сделать для хранения значения и индекса либо можно вместо значений температуры хранить индексы.
*/
func dailyTemperatures(temperatures []int) []int {
	warmerTempIdxStack := make([]int, 0, len(temperatures))
	result := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(warmerTempIdxStack) > 0 && temperatures[i] >= temperatures[warmerTempIdxStack[len(warmerTempIdxStack)-1]] {
			warmerTempIdxStack = warmerTempIdxStack[:len(warmerTempIdxStack)-1]
		}
		if len(warmerTempIdxStack) > 0 {
			result[i] = warmerTempIdxStack[len(warmerTempIdxStack)-1] - i
		} else {
			result[i] = 0

		}
		warmerTempIdxStack = append(warmerTempIdxStack, i)

	}
	return result
}
