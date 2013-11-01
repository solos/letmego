/*
    stack
 */
package stack

type Stack struct {
    i int
    data [10] int
}

func (s *Stack) Pop() int {
    s.i--
    return s.data[s.i]
}

func (s *Stack) Push(item int) {
    if s.i+1 > 9 {
        return
    }
    s.data[s.i] = item
    s.i++
}
