func checkInclusion(s1 string, s2 string) bool {
	permMap := make(map[byte]int)
	for i:=0;i<len(s1);i++ {
		permMap[s1[i]]++
	}

	left := 0
	charRemain := len(s1)
	currPermMap := make(map[byte]int)
	for i:=0;i<len(s2);i++ {
		currPermMap[s2[i]]++

		// grow
		_, exist := permMap[s2[i]]
		if exist && currPermMap[s2[i]] <= permMap[s2[i]] {
			charRemain--
		}

		// shrink
		if i >= len(s1) {
			if permMap[s2[left]] > 0 && currPermMap[s2[left]] <= permMap[s2[left]] {
				charRemain++
			} 

			currPermMap[s2[left]]--
			left++
		}

		if charRemain == 0 {
			return true
		}
	}

	if charRemain == 0 {
		return true
	}

	return false
}
