package main

import "fmt"

func main() {
    s := "asSASA ddd dsjkdsjs dk"
    n := s[0:3] + "abc" + s[4:]
    fmt.Printf("%s\n", n)
}
