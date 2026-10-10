import (
	"slices"
)


func combinationSum2(candidates []int, target int) [][]int {
	// sort candidates
	slices.Sort(candidates)
	return combi(target, candidates, []int{})
}

func combi(remain int, subset, combination []int) [][]int {
	if remain < 0 {
		return [][]int{}
	}
	if remain == 0 {
		return [][]int{combination}
	}

	res := [][]int{}
	for i:=0;i<len(subset);i++ {
		if i > 0 && subset[i] == subset[i-1] {
			continue
		}
		if subset[i] > remain {
			break
		}
		nCombi := append([]int{}, combination...)
		res = append(res, combi(remain-subset[i], subset[i+1:], append(nCombi,subset[i]))...)
		
	}

	return res
}
