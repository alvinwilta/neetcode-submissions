// using bucket
func topKFrequent(nums []int, k int) []int {
	freqMap := map[int]int{}
	for _, num := range nums {
		freqMap[num] += 1
	}

	bucket := make([][]int, len(nums)+1)
	max := 0
	// +1 to account for "0"
	// so that the index can be the actual length
	for num, freq := range freqMap {
		bucket[freq] = append(bucket[freq], num)
		if freq > max {
			max = freq
		}
	}

	// build solution
	solution := []int{}
	for i:=max;i>0;i-- {
		solution = append(solution, bucket[i]...)
		if len(solution) >= k {
			break
		}
	}

	return solution[:k]
}
