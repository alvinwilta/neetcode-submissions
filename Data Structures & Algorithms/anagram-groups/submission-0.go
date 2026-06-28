func groupAnagrams(strs []string) [][]string {
    // key: sorted string, value: actual string
    anagramMap := map[string][]string{}
    
    for _, str := range strs {
        anagramMapKey := []rune(str)
        sort.Slice(anagramMapKey, func(i, j int) bool {
            return anagramMapKey[i] < anagramMapKey[j]
        })
        anagramMap[string(anagramMapKey)] = append(anagramMap[string(anagramMapKey)], str)
    }

    result := [][]string{}
    for _, v := range anagramMap {
        result = append(result, v)
    }

    return result
}