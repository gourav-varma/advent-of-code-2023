package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	part1(scanner)

    file1, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner1 := bufio.NewScanner(file1)
	part2(scanner1)
}

func part1(scanner *bufio.Scanner){
	sum := 0
    for scanner.Scan() {
		curText := scanner.Text();
		strLength := len([]rune(curText))
		curNum := "" 
		for i := 0; i < strLength; i++{ 
			num, err := strconv.Atoi(curText[i:(i+1)])
			if err == nil {
				curNum += strconv.Itoa(num);
				break;
			}
		} 
		for i := strLength - 1; i >=0 ; i--{ 
			num, err := strconv.Atoi(curText[i:(i+1)]);
			if err == nil {
				curNum += strconv.Itoa(num);
				break;
			}
		} 
		num, err := strconv.Atoi(curNum)
		if err == nil {
			sum += num;
		}
    }
	fmt.Println(sum)
}

func part2(scanner *bufio.Scanner){
	sum := 0
    for scanner.Scan() {
		curText := scanner.Text();
		numStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		// for idx, str := range numStrings{
		// 	curText = strings.ReplaceAll(curText, str, strconv.Itoa(idx + 1))
		// }
		strLength := len([]rune(curText))
		curNum := "" 
		for i := 0; i < strLength; i++{ 
			num, err := strconv.Atoi(curText[i:(i+1)])
			if err == nil {
				curNum += strconv.Itoa(num);
				break;
			}else {
				flag := false
				for idx, str := range numStrings{
					if strings.HasPrefix(curText[i:], str) {
						curNum += strconv.Itoa(idx + 1);
						flag = true
						break;
					}
				}
				if flag {
					break;
				}
			}
		} 
		for i := strLength - 1; i >=0 ; i--{ 
			num, err := strconv.Atoi(curText[i:(i+1)]);
			if err == nil {
				curNum += strconv.Itoa(num);
				break;
			}else {
				flag := false
				for idx, str := range numStrings{
					if strings.HasPrefix(curText[i:], str) {
						curNum += strconv.Itoa(idx + 1);
						flag = true
						break;
					}
				}
				if flag {
					break;
				}
			}
		} 
		num, err := strconv.Atoi(curNum)
		if err == nil {
			sum += num;
		}
    }
	fmt.Println(sum)
}

// curSum := 0;
// arr := []int{}
// arr = append(arr, curSum)
// i, err := strconv.Atoi(curText)
// slices.Max(arr)
// slices.Sort(arr)
// sort.Sort(sort.Reverse(sort.IntSlice(arr)))