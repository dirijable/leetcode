package medium

type TimestampValue struct {
	timestamp int
	value     string
}

type TimeMap struct {
	data map[string][]TimestampValue
}

func ConstructorTM() TimeMap {
	return TimeMap{
		data: map[string][]TimestampValue{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], TimestampValue{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if data, ok := this.data[key]; ok {
		return binarySearchNearest(data, timestamp)
	}
	return ""
}

func binarySearchNearest(data []TimestampValue, target int) string {
	left, right := 0, len(data)-1
	for left < right {
		mid := (left + right + 1) / 2
		if data[mid].timestamp <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return data[right].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
