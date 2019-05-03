package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "log"
)

	
func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func main() {
    // Read in args, check if there are no args
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println("No command line args...\nExiting program")
        os.Exit(-1)
    } 
    // Get working directory
    dir, err := os.Getwd()
    check(err)
    // Store full path of files in array
    var files []string
    for _, arg := range args {
        fpath := dir + "\\" + arg
        files = append(files, fpath)
    }
    // Loop through files and check if they exist
    for _, filePath := range files {
        _, err := ioutil.ReadFile(filePath)
        check(err)
    }
}
