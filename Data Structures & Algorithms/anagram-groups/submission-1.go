func groupAnagrams(strs []string) [][]string {
	// k: sorted key, v: idx of strs
	solutionIndex := map[string][]int{}

	for i, str := range strs {
		strByte := []byte(str)
		sort.Slice(strByte, func(a,b int) bool {
			return strByte[a] > strByte[b]
		})
		solutionIndex[string(strByte)] = append(solutionIndex[string(strByte)],i)
	}

	// build solution
	sol := [][]string{}
	for _, strsIdxArr := range solutionIndex {
		tmp := []string{}
		for _, strIdx := range strsIdxArr {
			tmp = append(tmp, strs[strIdx])
		}
		sol = append(sol, tmp)
	}

	return sol
}
