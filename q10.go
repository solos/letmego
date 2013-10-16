package main

import "fmt"

func printr(arg ... int) {
    for _, i := range arg {
        fmt.Printf("%d\n", i)
    }
}

func main() {
    printr(1, 2, 3, 4, 5)
}
