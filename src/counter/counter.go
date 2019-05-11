package main

import (
    "fmt"
    "os"
    "crypto/sha1"
    "io/ioutil"
    "log"
    "strings"
)

	
func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func getWordCountMap(words []string) map[string]int {
    hashTable := make(map[string]int)
    for _, word := range words {
        h := sha1.New()
        h.Write([]byte(word))
        hash := string(h.Sum(nil))
        // If the hash is not already in the hash table...
        if _, ok := hashTable[hash]; !ok {
            hashTable[hash] = 0
        } else {
            hashTable[hash]++
        }
    }
    return hashTable
}

func printWordStats(words []string) {
    numWords := len(words)
    //wordMap := getWordCountMap(words)

    fmt.Printf("File contains %d words.\n", numWords)
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
    // For every file, print its word information
    for _, filePath := range files {
        data, err := ioutil.ReadFile(filePath)
        check(err)
        words := strings.Split(string(data), " ")
        printWordStats(words)
    }
}
