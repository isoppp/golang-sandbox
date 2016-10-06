// little edit.

package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    data := make(map[string]int)
    files := os.Args[1:]

    if len(files) == 0 {
        fmt.Println("need args file path")
        os.Exit(0)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, data)
            f.Close()
        }
    }

    for line, n := range data {
        fmt.Printf("%d\t%s\n", n, line)
    }
}

func countLines(f *os.File, data map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        data[input.Text()]++
    }
}
