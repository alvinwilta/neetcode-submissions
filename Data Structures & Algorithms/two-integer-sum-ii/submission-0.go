func twoSum(numbers []int, target int) []int {
	// two pointers
	left := 0
	right := len(numbers)-1

	for left != right {
		sum := numbers[left]+numbers[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left+1,right+1}
		}
	}

	// not possible to reach here
	return []int{}
}
