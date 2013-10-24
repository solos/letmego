package main

import "fmt"

func max(list [] int) int {
    r := list[0] 
    for _, v := range list {
        if (v > r) {
            r = v
        }
    }
    return r
}


func min(list [] int) int {
    r := list[0] 
    for _, v := range list {
        if (v < r) {
            r = v
        }
    }
    return r
}

func main() {
    l := []int{1, 3, 5, 7, 9, 7, 5, 3, 1}
    fmt.Printf("max of %v is %d", l, max(l))
    fmt.Printf("min of %v is %d", l, min(l))
}
