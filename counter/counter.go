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
    data, err := ioutil.ReadFile(filePath)
    check(err)
    // Use regular expression as word delimiters and whitespace trimmers.
    spaces := regexp.MustCompile(`\s| `)
    trim := regexp.MustCompile("[^a-zA-Z0-9_]+")
    // Store data from file in array, delimit on whitespace.
    rawWords := spaces.Split(string(data), -1)
    // Store each word in array after it has been stripped of whitespace and made lowercase.
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
func getFrequencyMap(wordMap map[string]int) map[int][]string {
    freqMap := make(map[int][]string)
    for word, frequency := range wordMap {
        if _, inMap := freqMap[frequency]; !inMap {
            freqMap[frequency] = []string{}
        }
        freqMap[frequency] = append(freqMap[frequency], word)
    }
    return freqMap
}

/*
    Print data pertinent to the text file.
*/
func printWordStats(words []string) {
    wordMap := getWordCountMap(words)
    freqMap := getFrequencyMap(wordMap)
    fmt.Printf("File contains %d words.\n", len(words))
    fmt.Printf("File contains %d unique words.\n", len(wordMap))
    var rankSlice []int
    for frequency := range freqMap {
        rankSlice = append(rankSlice, frequency)
    }
    // Sort in descending order of frequency.
    sort.Slice(rankSlice, func(i, j int) bool {
        return rankSlice[i] > rankSlice[j]
    })
    // Print words from most to least frequently used.      
    size := len(rankSlice)
    rankSlice = rankSlice[:size-1]
    for rank, frequency := range rankSlice {
        fmt.Printf("Rank: %3d\tFrequency: %3d---> ", rank+1, frequency)
        fmt.Println(freqMap[frequency])
    }
}

/*
    Print word info for each file in the working directory.
*/
func printDirectory() {
    files, err := ioutil.ReadDir(".")
    check(err)
    dir, err := os.Getwd()
    check(err)
    for _, file := range files {
        mode := file.Mode()
        if mode.IsRegular() {
            printWordStats(getWords(fmt.Sprintf("%s\\%s", dir, file.Name())))
        }
    }
}

/*
    Print word info for each file specified in the working directory.
    'args' is the array of desired file.
*/
func printFiles(args []string) {
    dir, err := os.Getwd()
    check(err)
    for _, arg := range args {
        fileInfo, _ := os.Stat(arg)
        if fileInfo.Mode().IsRegular() {
            printWordStats(getWords(fmt.Sprintf("%s\\%s", dir, arg)))
        }
    }
}

/*
    Main function. Read in command line arguements.
    For each file, print out info about that text file.
    If no files specified, print info for files in working directory.
*/
func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        printDirectory()
    } else {
        printFiles(args)
    }
}
