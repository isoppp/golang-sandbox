// little edit.
package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") {
            if line == "" {
                continue;
            }

            counts[line]++
        }
    }
    for line, n := range counts {
        fmt.Printf("%d\t%s\n", n, line)
    }
}
