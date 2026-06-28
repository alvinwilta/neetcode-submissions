func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    runeMap := map[rune]int{}
    for _, letter := range s {
        runeMap[letter] += 1
    }

    for _, letter := range t {
        runeMap[letter] -= 1
    }

    for _, v := range runeMap {
        if v > 0 {
            return false
        }
    }
    
    return true
}

