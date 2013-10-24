package main

import "fmt"

func remap(f func(v int) int, list []int) []int {
    values := make([]int, len(list))
    for i := 0; i < len(list); i++ {
        v := f(list[i])
        values[i] = v
    }
    return values
}

func f(arg int) int {
    return arg * arg
}

func Map(f func(int) int, l []int) []int {
    j := make([]int, len(l))
    for k, v := range l {
        j[k] = f(v)
    }
    return j
}

func main() {
    list := []int{1,3,4}
    fmt.Printf("%v\n", remap(f, list))
    fmt.Printf("%v\n", Map(f, list))
}
