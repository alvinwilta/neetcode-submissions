func twoSum(num []int, target int) []int {
	left := 0
	right := len(num)-1

	for left < right {
		sum := num[left]+num[right]
		if sum == target {
			return []int{left+1,right+1}
		}

		if sum > target {
			right--
			continue
		}

		if sum < target {
			left++
			continue
		}
	}

	return []int{0,0}
}
