package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "regexp"
    "strings"
)


func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

/*
    Store and return words read from the specified text file.
*/
func getWords(filePath string) []string {
    // Read in data from the specified file path.
    data, err := ioutil.ReadFile(filePath)
    check(err)
    // Use regular expression as word delimiters.
    spaces := regexp.MustCompile(`\s| `)
    trim := regexp.MustCompile("[^a-zA-Z0-9_]+")
    // 
    rawWords := spaces.Split(string(data), -1)
    var cleanWords []string
    for _, word := range rawWords {
        cleanWord := trim.ReplaceAllString(word, "")
        cleanWord = strings.ToLower(cleanWord)
        if cleanWord != "" {
            cleanWords = append(cleanWords, cleanWord)
        }
    }
    return cleanWords
}

/*
    Return a mapping of unique words as strings to the number of times each appears in the text file.
*/
func getWordCountMap(words []string) map[string]int {
    wordMap := make(map[string]int)
    for _, word := range words {
        if _, ok := wordMap[word]; !ok {
            wordMap[word] = 1
        } else {
            wordMap[word]++
        }
    }
    return wordMap
}

/*
    Print data pertinent to the text file.
*/
func printWordStats(words []string) {
    wordMap := getWordCountMap(words)
    fmt.Printf("File contains %d words.\n", len(words))
    fmt.Printf("File contains %d unique words.\n", len(wordMap))
    // Print words from most to least frequently used.
    //rank := 0
    for key, val := range wordMap {
        fmt.Printf("%-12s\t%d\n", key, val)
    }
}

/*
    Main function. Read in command line arguements.
    For each file, print out info about that text file.
*/
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
