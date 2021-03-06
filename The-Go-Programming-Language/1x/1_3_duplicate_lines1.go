// little edit.
package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)

    for input.Scan() {
        if input.Text() == "" {
            break;
        }
        counts[input.Text()]++
    }
    fmt.Println(counts)
    for line, n := range counts {
        fmt.Printf("%d\t%s\n", n, line)
    }
}

