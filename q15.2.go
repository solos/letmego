package main

import "fmt"

func plusX(x int) func(int) int {
    return func(y int) int { return x + y }
}

func main() {
    p := plusX(2)
    fmt.Printf("%v\n", p(3))
}
