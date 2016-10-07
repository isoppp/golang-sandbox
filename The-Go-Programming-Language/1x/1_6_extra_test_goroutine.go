// refer: http://qiita.com/suin/items/82ecb6f63ff4104d4f5d
// refer: http://qiita.com/TakaakiFuruse/items/241578174fd2f00aaa8a

package main

import (
    "fmt"
    "time"
)

func sleepOneSecond() {
    time.Sleep(1000 * time.Millisecond)
    fmt.Println("1 seconds sleeped");
}

func sleepTwoSecond() {
    time.Sleep(2000 * time.Millisecond)
    fmt.Println("2 seconds sleeped");
}

func sleepThreeSecond() {
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("3 seconds sleeped");
}

func test1() {
    go sleepThreeSecond()
    sleepOneSecond()
    sleepTwoSecond()

    // logged 1-3-2
}

func main() {
    fmt.Println(time.Now())
    test1()
    fmt.Println(time.Now())
}
