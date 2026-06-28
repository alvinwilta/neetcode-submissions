
func twoSum(number []int, target int) []int {
	left := 0
	right := len(number)-1
	for left < right {
		sum := number[left]+number[right]
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

	return []int{}
}