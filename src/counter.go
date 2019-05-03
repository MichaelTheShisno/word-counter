package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Hello, Word Counter!")
    
    args := os.Args[1:]
    for _, arg := range args {
        fmt.Println(arg)
    }
}
