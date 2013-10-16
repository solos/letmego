package main

import (
    "strconv"
    "fmt"
)

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

func (s stack) String() string {
    var str string
        for i := 0 ; i <= s.i ; i++ {
            str = str + "[" +
                strconv.Itoa(i) + ":" + strconv.
                Itoa(s.data[i]) + "]"
        }
    return str
}

func main() {
    var s stack
    s.push(25)
    fmt.Printf("%s\n", s.String())
    s.push(14)
    fmt.Printf("%s\n", s.String())
}
