package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "regexp"
    "strings"
    "sort"
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
            wordMap[word] = 0
        }
        wordMap[word]++
    }
    return wordMap
}

/*
    Return mapping of frequencies to array of strings with the corresponding frequency.
*/
func getFrequencyMap(wMap map[string]int) map[int][]string {
    fMap := make(map[int][]string)
    for word, frequency := range wMap {
        if _, inMap := fMap[frequency]; !inMap {
            fMap[frequency] = []string{}
        }
        fMap[frequency] = append(fMap[frequency], word)
    }
    return fMap
}

/*
    Print data pertinent to the text file.
*/
func printWordStats(words []string) {
    wordMap := getWordCountMap(words)
    freqMap := getFrequencyMap(wordMap)
    fmt.Printf("File contains %d words.\n", len(words))
    fmt.Printf("File contains %d unique words.\n", len(wordMap))
    // Print words from most to least frequently used.      
    var rankSlice []int
    for frequency := range freqMap {
        rankSlice = append(rankSlice, frequency)
    }
    sort.Slice(rankSlice, func(i, j int) bool {
        return rankSlice[i] > rankSlice[j]
    })
    size := len(rankSlice)
    rankSlice = rankSlice[:size-1]
    for rank, frequency := range rankSlice {
        fmt.Printf("Rank: %3d\tFrequency: %3d---> ", rank+1, frequency)
        fmt.Println(freqMap[frequency])
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
