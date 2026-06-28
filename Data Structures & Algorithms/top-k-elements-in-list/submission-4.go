func topKFrequent(nums []int, k int) []int {
	// store count for each element, k: num, v: count
	numCountMap := make(map[int]int)

	// calculate count
	for _, num := range nums {
		numCountMap[num] += 1
	}

	// maps count to element, k: count, v: arr of el for that count
	bucket := make([][]int, len(nums)+1) // to account for index 0

	// map count
	for num, count := range numCountMap {
		bucket[count] = append(bucket[count], num)
	}

	// build sorted k frequent element
	result := []int{}
	for i:=len(nums);i>0;i-- {
		if len(bucket[i]) == 0 {
			continue
		}
		result = append(result, bucket[i]...)
		if len(result) >= k {
			break
		}
	}

	return result[:k]
}
