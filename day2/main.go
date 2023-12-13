package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	// "unicode/utf8"
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
		curText := scanner.Text()
		idSpilt := strings.Split(curText, ":")
		gameid, err := strconv.Atoi(strings.Split(idSpilt[0], " ")[1]);
		if err != nil {
			fmt.Println("Invalid Game ID")
		}

		flag := false
		for _, val := range strings.Split(idSpilt[1], ";") {
			for _, val1 := range strings.Split(val, ",") {
				split := strings.Split(val1, " ")
				num, err := strconv.Atoi(split[1])
				if err != nil {
					fmt.Println("Error")
				}
				if (split[2] == "red" && num > 12) || (split[2] == "green" && num > 13) || (split[2] == "blue" && num > 14) {
					flag = true
					break;
				}
			}
			if flag {
				break;
			}
		}
		if !flag {
			sum += gameid
		}
    }
	fmt.Println(sum)
}

func part2(scanner *bufio.Scanner){
	sum := 0
    for scanner.Scan() {
		curText := scanner.Text()
		idSpilt := strings.Split(curText, ":")
		// gameid, err := strconv.Atoi(strings.Split(idSpilt[0], " ")[1]);
		// if err != nil {
		// 	fmt.Println("Invalid Game ID")
		// }

		arr := []int{math.MinInt64, math.MinInt64, math.MinInt64}
		for _, val := range strings.Split(idSpilt[1], ";") {
			for _, val1 := range strings.Split(val, ",") {
				split := strings.Split(val1, " ")
				num, err := strconv.Atoi(split[1])
				if err != nil {
					fmt.Println("Error")
				}
				switch split[2] {
					case "red": 
						arr[0] = int(math.Max(float64(arr[0]), float64(num)))
					case "green": 
						arr[1] = int(math.Max(float64(arr[1]), float64(num)))
					case "blue": 
						arr[2] = int(math.Max(float64(arr[2]), float64(num)))
				}
			}
		}
		prod := 1
		for _, nums := range arr {
			if nums != math.MinInt64 {
				prod *= nums;
			}
		}
		sum += prod
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