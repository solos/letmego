package main 

import "fmt"

func bubblesort(list []int) []int {
    for i := 0; i < len(list) - 1; i++ {
        for j:= i+1; j < len(list); j++ {
            if (list[j] < list[i]) {
                list[i], list[j] = list[j], list[i]
            }
        }
    }
    return list
}

func main() {
    l := []int{1,3,5,7,0,2}
    fmt.Printf("%v\n", bubblesort(l))
}
