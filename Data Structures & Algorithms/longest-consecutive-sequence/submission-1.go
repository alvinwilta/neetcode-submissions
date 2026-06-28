func longestConsecutive(nums []int) int {
	// key: the "next" element, val: the current length excl the "next"
	set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}

	// find the solution
	max := 0
	for num := range set {
		// skip non-start num
		if set[num-1] {
			continue
		}

		length := 1
		for set[num+length] {
			length++
		}
		if length > max {
			max = length
		}
	}

	return max
}
