package comma

import (
    "bytes"
    "strings"
)

func comma(s string) string {
    var buf bytes.Buffer
    
    // is first letter a sign
    if s[0] == '+' || s[0] == '-' {
        buf.WriteByte(s[0])
        s = s[1:]
    }
    
    // if there is a floating-point, then print until floating-point
    pt := strings.Index(s, ".")
    n := len(s)
    if pt < 0 {
        pt = n
    }     
    
    for i := 0; i < pt; i++ {
        buf.WriteByte(s[i])
        d := (pt-i-1)
        if d > 0 && d%3 == 0 {
            buf.WriteByte(',')
        }
    }
    
    if pt < n {
        buf.WriteString(s[pt:])
    }
    
    return buf.String()
}

func isAnagram(s1, s2 string) bool {    
    if len(s1) != len(s2) {
        return false;
    }
    
    m := make(map[rune]bool)
    for _, r := range s1 {
        m[r] = true
    }
    
    for _, r := range s2 {
        if _, ok := m[r]; !ok {
            return false
        }
    }
    return true
}