package main

import "fmt"

func sort(arg1, arg2 int) (int, int) {
    if (arg1 > arg2) {
        return arg2, arg1
    }
    return arg1, arg2
}

func main() {
    arg1 := 2
    arg2 := 1
    a, b := sort(arg1, arg2)
    fmt.Printf("%d\t%d\n", a, b)
}
