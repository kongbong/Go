package main

import (
    "log"
    "net/http"
    "strconv"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        err := r.ParseForm()
        cycles := 5
        if err == nil {
            cycles, err = strconv.Atoi(r.Form.Get("cycles"))    
        }
        lissajous(w, cycles)
    })
    log.Fatal(http.ListenAndServe("localhost:8000", nil))    
}
