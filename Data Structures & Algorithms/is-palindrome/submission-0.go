func isPalindrome(s string) bool {
    if len(s) < 2 {
        return true
    }
    
    s = sanitizeString(s)
    return checkReverse(s)
}

func sanitizeString(s string) string {
    newStr := ""
    s = strings.ToLower(s)
    for _, char := range s {
        if ('a' <= char && char <= 'z') || ('0' <= char && char <= '9') {
            newStr += string(char)
        }
        continue
    }

    return newStr
}

func checkReverse(s string) bool {
    for i:=0;i<len(s)/2;i++ {
        if s[i] != s[len(s)-i-1] {
            return false
        }
    }
    return true
}
