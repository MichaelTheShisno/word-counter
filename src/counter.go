package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println("No command line args...\nExiting program")
        os.Exit(-1)
    } else {
        for _, arg := range args {
            fmt.Println(arg)
        }
    }
}
