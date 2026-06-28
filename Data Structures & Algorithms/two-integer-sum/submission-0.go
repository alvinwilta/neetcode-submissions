func twoSum(nums []int, target int) []int {
    if len(nums) <= 2 {
        return []int{0,1}
    }
    // key: pairing, val: index 
    remain := map[int]int{}

    for i, num := range nums {
        rIdx, exist := remain[num] 
        if exist {
            return []int{rIdx,i}
        }

        remain[target-num] = i
    }

    return []int{0,1}
}
