func longestConsecutive(nums []int) int {
	// have a map that stores the previous values
	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true
	}

	maxLength := 0
	for _, num := range nums {
		if _, exist := numsMap[num-1]; exist {
			continue
		}
		length := 1
		for {
			if !numsMap[num+length] {
				break
			}
			length++
		}
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
