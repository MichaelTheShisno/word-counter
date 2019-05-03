package main

import (
    "fmt"
    "os"
    "log"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println("No command line args...\nExiting program")
        os.Exit(-1)
    } 

    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    var files []string
    for _, arg := range args {
        fpath := dir + "\\" + arg
        files = append(files, fpath)
    }

    fmt.Println(files)
}
