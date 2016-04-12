package fetch

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

func fetch(urls []string) {
    for _, url := range urls {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        b, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)            
        }
        fmt.Printf("%s", b)
    }
}

func fetch2(urls []string) {
    for _, url := range urls {
        if !strings.HasPrefix(url, "http://") {
            url = "http://" + url
        }
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        fmt.Println(resp.Status)        
        _, err = io.Copy(os.Stdout, resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)            
        }        
    }
}