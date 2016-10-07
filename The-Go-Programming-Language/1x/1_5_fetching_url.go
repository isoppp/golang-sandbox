package main

import (
    "os"
    "net/http"
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        if (!strings.HasPrefix(url, `http://`)) {
            url = `http://` + url
        }

        resp, err := http.Get(url)
        status := resp.Status

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

        fmt.Printf("status code: %v\n body:%s", status, b)
    }
}
