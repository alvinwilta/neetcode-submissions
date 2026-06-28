func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	duplMap := make(map[byte]int)
	var maxLength, anchor int

	for i:=0;i<len(s);i++ {
		// slide anchor
		idx, exist := duplMap[s[i]]
		if exist {
			anchor = max(anchor,idx+1) // +1 to not include the dupl on i
		}

		// calculate length
		duplMap[s[i]] = i
		length := i-anchor+1
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func max(a,b int) int {
	if a>b {return a}
	return b
}
