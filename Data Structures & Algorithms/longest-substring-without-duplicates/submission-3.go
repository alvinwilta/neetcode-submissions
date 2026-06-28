func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	maxLength := 0
	prevIdx := 0
	charMap := make(map[byte]int)
	for i:=0;i<len(s);i++ {
		duplIdx, ok := charMap[s[i]]
		if ok && prevIdx <= duplIdx {
			// duplicates
			prevIdx = duplIdx+1
		}
		charMap[s[i]] = i


		// check max length
		if maxLength < i-prevIdx+1 {
			maxLength = i-prevIdx+1
		}
	}

	return maxLength
}
