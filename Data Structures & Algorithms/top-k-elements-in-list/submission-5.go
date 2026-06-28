func topKFrequent(nums []int, k int) []int {
	// build freq map
	freqMap := map[int]int{}
	for _, num := range nums {
		freqMap[num] += 1
	}

	// build freq order
	freqOrder := map[int][]int{}
	freqOrderKey := []int{}
	for num, freq := range freqMap {
		if len(freqOrder[freq]) == 0 {
			freqOrderKey = append(freqOrderKey, freq)
		}
		freqOrder[freq] = append(freqOrder[freq], num)
	}

	// sort freq key
	sort.Slice(freqOrderKey, func(a,b int) bool {
		return freqOrderKey[a] > freqOrderKey[b]
	})

	// build solution
	solution := []int{}
	for _, freq := range freqOrderKey {
		if k == 0 {
			break
		}
		for _, num := range freqOrder[freq] {
			if k == 0 {
				break
			}
			solution = append(solution, num)
			k -= 1
		}
	}

	return solution
}
