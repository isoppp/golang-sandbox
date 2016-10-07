package main

import (
    "os"
    "net/http"
    "fmt"
    "strings"
    "io"
    "log"
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

        if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
        fmt.Println()
        fmt.Printf("status code: %v", status)
        resp.Body.Close()
    }
}
