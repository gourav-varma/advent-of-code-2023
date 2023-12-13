package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "slices"
	// "sort"
)
 
func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

	inputText := ""
    for scanner.Scan() {
		curText := scanner.Text();
		inputText += curText
    }
}

// curSum := 0;
// arr := []int{}
// arr = append(arr, curSum)
// i, err := strconv.Atoi(curText)
// slices.Max(arr)
// slices.Sort(arr)
// sort.Sort(sort.Reverse(sort.IntSlice(arr)))