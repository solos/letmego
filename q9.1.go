package main

import "fmt"

type stack struct {
    i int
    data [10] int
}

func (s *stack) pop() int {
    s.i--
    return s.data[s.i]
}

func (s *stack) push(item int) {
    if s.i+1 > 9 {
        return
    }
    s.data[s.i] = item
    s.i++
}

func main() {
    var s stack
    s.push(25)
    fmt.Printf("stack %v\n", s) ;
    s.push(14)
    fmt.Printf("stack %v\n", s) ;
}
