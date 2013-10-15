package main

import "fmt"

func main() {
    i := 0
    for i < 100 {
        if (i % 3 == 0) && (i % 5 == 0) {
            fmt.Printf("FizzBuzz\n")
        } else {
            if i % 3 == 0 {
                fmt.Printf("Buzz\n")
            } else {
                fmt.Printf("Fizz\n")
            }
        }
        i ++
    }
}
