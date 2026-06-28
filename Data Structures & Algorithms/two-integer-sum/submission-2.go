func twoSum(nums []int, target int) []int {
	remainMap := map[int]int{}
    for i, num := range nums {
		j, exist := remainMap[num]
		if exist {
			return []int{j, i}
		}

		remainMap[target-num] = i
	}

	return []int{0,0}
}
