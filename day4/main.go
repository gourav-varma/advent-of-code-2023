package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	// "slices"
	"strconv"

	// "strconv"
	"strings"
	// "unicode/utf8"
	// "sort"
)
 
func main() {
    // file, err := os.Open("sample.txt")
    // file, err := os.Open("input.txt")
    // if err != nil {
    //     fmt.Println(err)
    // }
    // defer file.Close()
    // scanner := bufio.NewScanner(file)
	// part1(scanner)

    file1, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file1.Close()
    scanner1 := bufio.NewScanner(file1)
	part2(scanner1)
}

// func part1(scanner *bufio.Scanner){
// 	sum := 0
//     for scanner.Scan() {
// 		winning := [100]int{}
// 		mine := [100]int{}
// 		curText := scanner.Text()
// 		split := strings.Split(curText, ":")
// 		split = strings.Split(split[1], "|")
// 		for _, v := range strings.Split(strings.Trim(split[0], " "), " ") {
// 			val, err := strconv.Atoi(v)
// 			if err != nil {
// 				fmt.Println("wow and error")
// 			}
// 			winning[val] += 1;
// 		}
// 		for _, v := range strings.Split(strings.Trim(split[1], " "), " ") {
// 			if v == " " || v == "" {continue}
// 			val, err := strconv.Atoi(v)
// 			if err != nil {
// 				fmt.Println("wow and error")
// 			}
// 			mine[val]++;
// 		}
// 		cur_val := 0
// 		for i, v := range winning {
// 			if v != 0 && mine[i] != 0 {
// 				temp := v
// 				for {
// 					temp--;
// 					if temp == -1 {break}
// 					if cur_val == 0 {
// 						cur_val = 1
// 						continue
// 					}
// 					cur_val *= 2
// 				}
// 			}
// 		}
// 		sum += cur_val
// 	}
// 	fmt.Println(sum)
// }

func part2(scanner *bufio.Scanner){
	total := [250]int{}
	for i:=0; i<250; i++ {
		total[i] = 1
	}
	idx := 0 
    for scanner.Scan() {
		winning := [100]int{}
		mine := [100]int{}
		curText := scanner.Text()
		split := strings.Split(curText, ":")
		split = strings.Split(split[1], "|")
		for i, v := range strings.Split(strings.Trim(split[0], " "), " ") {
			if v == " " || v == "" {continue}
			val, err := strconv.Atoi(v)
			if err != nil {
				
				fmt.Println("wow and error", v, i, idx)
			}
			winning[val] += 1;
		}
		for _, v := range strings.Split(strings.Trim(split[1], " "), " ") {
			if v == " " || v == "" {continue}
			val, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("wow and error2")
			}
			mine[val]++;
		}

		cur_val := 0
		for i, v := range winning {
			if v != 0 && mine[i] != 0 {
				cur_val++
			}
		}
		cur := total[idx]
		for i:=1; i<=cur_val; i++ {
			total[idx + i] += cur
		}
		idx++
	}
	res := 0
	for i:=0; i<idx; i++ {
		res += total[i]
	}
	fmt.Println(res)
}


// curSum := 0;
// arr := []int{}
// arr = append(arr, curSum)
// i, err := strconv.Atoi(curText)
// slices.Max(arr)
// slices.Sort(arr)
// sort.Sort(sort.Reverse(sort.IntSlice(arr)))