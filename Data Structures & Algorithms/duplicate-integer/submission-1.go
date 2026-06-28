func hasDuplicate(nums []int) bool {
    duplMap := map[int]bool{}
	for _, num := range nums {
		if duplMap[num] {
			return true
		}

		duplMap[num] = true
	}

	return false
}
