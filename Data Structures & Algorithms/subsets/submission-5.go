func subsets(nums []int) [][]int {
	// convert to map for unique
	uniqMap := map[int]bool{}
	for _, num := range nums {
		uniqMap[num] = true
	}
	nNums := []int{}
	for k,_ := range uniqMap {
		nNums = append(nNums, k)
	}
	
	return subset([]int{},nNums)
}

func subset(base, remain []int) [][]int {
	if len(remain) == 0 {
		return [][]int{base}
	}

	ret := [][]int{}
	ret = append(ret, base)
	for i:=0;i<len(remain);i++ {
		newBase := append([]int{}, base...)
		newBase = append(newBase,remain[i])
		ret = append(ret,subset(newBase,remain[i+1:])...)
	}

	return ret
}