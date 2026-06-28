func characterReplacement(s string, k int) int {
	if len(s) < 2 {
		return len(s)
	}

	charMap := make(map[byte]int)
	max := 0
	maxCount := 0
	left := 0 // pointer for sliding window
	for i:=0;i<len(s);i++ {
		// expanding
		charMap[s[i]]++
		if maxCount < charMap[s[i]] {
			maxCount = charMap[s[i]]
		}

		// shrink if not valid window
		// if below condition is not fulfilled, the window is only growing
		if i+1-left-maxCount > k {
			charMap[s[left]]--
			left++
		}

		// update max (important logic here)
		// if it's not valid, we already shrank the left, so max won't update at all
		if max < i+1-left {
			max = i+1-left
		}
	}

	return max
}

