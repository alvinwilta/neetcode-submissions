type TimeMap struct {
	store map[string][][]string
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][][]string),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], []string{strconv.Itoa(timestamp),value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	n, exist := this.store[key]
	if !exist {
		return "" 
	}

	var val string
	left, right := 0, len(n)-1

	for left <= right {
		mid := (left+right)/2

		t, _ := strconv.Atoi(n[mid][0]) 
		if t <= timestamp {
			val = n[mid][1]
			left = mid+1
		} else {
			right = mid-1
		}
	}

	return val
}
