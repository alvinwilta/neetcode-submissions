func combinationSum(nums []int, target int) [][]int {
    return combi(target, nums, []int{})
}

func combi(remain int, nums, subset []int) [][]int {
	if remain < 0 {
		return [][]int{}
	}
	if remain == 0 {
		return [][]int{subset}
	}

	ret := [][]int{}
	for i, num := range nums {
		sub := append([]int{}, subset...)
		res := combi(remain-num, nums[i:], append(sub,num))
		ret = append(ret, res...)
	}
	return ret
}