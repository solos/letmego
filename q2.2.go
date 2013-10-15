package main

import "fmt"

func main() {
    i := 0
    LOOP:
        if (i < 10) {
            fmt.Printf("%d\n", i)
            i++
            goto LOOP
        }
}
