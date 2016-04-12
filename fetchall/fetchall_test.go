package fetchall

import "testing"

func TestFetchall(t *testing.T) {    
    fetchall([]string{"https://golang.org", "http://gopl.io", "https://godoc.org"})
}