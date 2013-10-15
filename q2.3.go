package main

import "fmt"

func main() {
    arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    for _, value := range arr {
        fmt.Printf("%d\n", value)
    }
}
