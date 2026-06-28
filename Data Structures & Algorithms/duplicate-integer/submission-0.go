func hasDuplicate(nums []int) bool {
    duplMap := map[int]bool{}

    for _, val := range nums {
        if duplMap[val] {
            return true
        }
        duplMap[val] = true
    }
    return false
}
