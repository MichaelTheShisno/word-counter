package main

import (
    "crypto/sha1"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strings"
    "regexp"
)

	
func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func getWords(filePath string) []string {
    data, err := ioutil.ReadFile(filePath)
    check(err)
    rawWords := strings.Split(string(data), " ")
    var cleanWords []string
    for _, word := range rawWords {
        cleanWord := strings.TrimSpace(word)
        fmt.Println(cleanWord)
        if hasNonWordCharacter, _ := regexp.MatchString(cleanWord, "[^a-zA-Z0-9_]+"); !hasNonWordCharacter {
            cleanWords = append(cleanWords, cleanWord)
        }
    }
    return cleanWords
}


func getWordCountMap(words []string) map[string]int {
    hashTable := make(map[string]int)
    for _, word := range words {
        h := sha1.New()
        h.Write([]byte(word))
        hash := string(h.Sum(nil))
        // If the hash is not already in the hash table...
        if _, ok := hashTable[hash]; !ok {
            fmt.Println()
            hashTable[hash] = 0
        } else {
            hashTable[hash]++
        }
    }
    return hashTable
}

func printWordStats(words []string) {
    wordMap := getWordCountMap(words)
    numWords := len(words)
    numUniqueWords := len(wordMap)
    fmt.Printf("File contains %d words.\n", numWords)
    fmt.Printf("File contains %d unique words.\n", numUniqueWords)
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
        printWordStats(getWords(filePath))
    }
}
