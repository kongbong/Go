package comma

import "testing"

func TestComma(t *testing.T) {
    s := comma("12345")
    if s != "12,345" {
        t.Errorf("comma not correct %v", s)
    }
    
    s = comma("123")
    if s != "123" {
        t.Errorf("comma not correct %v", s)
    }
    
    s = comma("+123")
    if s != "+123" {
        t.Errorf("comma not correct %v", s)
    }
    
    s = comma("-12345")
    if s != "-12,345" {
        t.Errorf("comma not correct %v", s)
    }
    
    s = comma("12.3")
    if s != "12.3" {
        t.Errorf("comma not correct %v", s)
    }
}

func TestIsAnagram(t * testing.T) {
    if !isAnagram("abc", "cba") {
        t.Error("FAIL abc and cba should be anagram")
    }
    
    if isAnagram("abcd", "cba") {
        t.Error("FAIL abcd and cba should be not anagram")
    }
    
    if !isAnagram("가나abc", "나bca가") {
        t.Error("FAIL 가나abc and 나bca가 should be anagram")
    }
}