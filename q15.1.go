package main

import "fmt"

func plusTwo() func(int) int {
    return func(x int) int { return x + 2 }
}

func main() {
    p := plusTwo()
    fmt.Printf("%v\n", p(2))
}
