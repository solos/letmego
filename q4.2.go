package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    s := "asSASA ddd dsjkdsjs dk"
    b := [] byte(s)
    fmt.Printf("String: %s\nLength: %d\nRune: %d\n", s, len(b), utf8.RuneCount(b))
}
