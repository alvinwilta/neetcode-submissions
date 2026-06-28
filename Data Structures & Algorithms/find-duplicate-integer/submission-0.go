func findDuplicate(nums []int) int {
    loc := nums[0]
	count := len(nums)
	for count > 0 {
		if nums[loc] == loc {
			return loc
		}

		tmp := nums[loc]
		nums[loc] = loc
		loc = tmp
		count--
	}

	return 0
}
