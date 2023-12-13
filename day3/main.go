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
    // file1, err := os.Open("sample.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file1.Close()
    scanner1 := bufio.NewScanner(file1)
	part2(scanner1)
}

// 533421
// func part1(scanner *bufio.Scanner){
// 	sum := 0
// 	mat := [][]string{}
//     for scanner.Scan() {
// 		curText := scanner.Text()
// 		mat = append(mat, strings.Split(curText, ""));
// 	}
// 	n := len(mat)
// 	m := len(mat[0])
// 	for i:=0; i<n; i++ {
// 		start := -1
// 		curVal := ""
// 		for j:=0; j<m; j++ {
// 			num, err := strconv.Atoi(mat[i][j])
// 			if err != nil {
// 				if(start != -1){
// 					if(checkAround(start, j-1, i, mat, n, m)){
// 						num1, _ := strconv.Atoi(curVal)
// 						sum += num1
// 					}
// 				}
// 				start = -1;
// 				curVal = "";
// 				continue;
// 			}
// 			if(start == -1){
// 				start = j
// 				curVal += strconv.Itoa(num);
// 			}else {
// 				curVal += strconv.Itoa(num);
// 			}
// 		}
// 		if(start != -1){
// 			if(checkAround(start, m-1, i, mat, n, m)){
// 				num1, _ := strconv.Atoi(curVal)
// 				sum += num1
// 			}
// 		}
// 	}
// 	fmt.Println(sum)
// }

// 64852549 too low
func part2(scanner *bufio.Scanner){
	sum := 0
	mat := [][]string{}
    for scanner.Scan() {
		curText := scanner.Text()
		mat = append(mat, strings.Split(curText, ""));
	}
	
	n := len(mat)
	m := len(mat[0])

	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if(mat[i][j] == "*"){
				// fmt.Println(i, j)
				var check, curProd = checkAroundPart2(i, j, mat, n, m)
				if(check){
					sum += curProd
				}
				// fmt.Println("finalMainSum", sum)
			}
		}
		// if(start != -1){
		// 	var check, curProd = checkAroundPart2(i, j, mat, n, m)
		// 	if(check){
		// 		sum += curProd
		// 	}
		// }
	}
	fmt.Println(sum)
}

func checkAroundPart2(i, j int, mat [][]string, n int, m int) (bool, int) {
	// start := int(math.Max(float64(j) - 1, 0))
	// end := int(math.Min(float64(j) + 1, float64(m-1)))

	totalNum := 0
	prod := 1
	// fmt.Println("top")
	// top
	{
		if(i != 0){
			// if middle idx is num: there is 1 number
			if(checkIfNum(mat[i-1][j]))	{
				totalNum += 1;
				prod *= getCompleteNum(j, mat[i-1])
			}else {
				var left, right bool = checkIfNum(mat[i-1][max(j-1, 0)]), checkIfNum(mat[i-1][min(j+1, m-1)])
				if left && right{
					totalNum += 2;
					prod *= getCompleteNum(j-1, mat[i-1])
					prod *= getCompleteNum(j+1, mat[i-1])
				}else if left {
					totalNum += 1;
					prod *= getCompleteNum(j-1, mat[i-1])
				}else if right {
					totalNum += 1;
					prod *= getCompleteNum(j+1, mat[i-1])
				}
			} 
		}
	}
	// fmt.Println("L&R", mat[i][j-1], mat[i][j], mat[i][j+1])
	// left and right
	{
		if(j != 0){
			if(checkIfNum(mat[i][j-1]))	{
				fmt.Println("in1", j-1)
				totalNum += 1;
				prod *= getCompleteNum(j-1, mat[i])
			}
		}
		if(j != m-1){
			if(checkIfNum(mat[i][j+1]))	{
				fmt.Println("in2")
				totalNum += 1;
				prod *= getCompleteNum(j+1, mat[i])
			}
		}
	}
	// fmt.Println("Bottom")
	// bottom
	{
		if(i != (m-1)){
			// if middle idx is num: there is 1 number
			if(checkIfNum(mat[i+1][j]))	{
				totalNum += 1;
				prod *= getCompleteNum(j, mat[i+1])
			}else {
				var left, right bool = checkIfNum(mat[i+1][max(j-1, 0)]), checkIfNum(mat[i+1][min(j+1, m-1)])
				if left && right{
					totalNum += 2;
					prod *= getCompleteNum(j-1, mat[i+1])
					prod *= getCompleteNum(j+1, mat[i+1])
				}else if left {
					totalNum += 1;
					prod *= getCompleteNum(j-1, mat[i+1])
				}else if right {
					totalNum += 1;
					prod *= getCompleteNum(j+1, mat[i+1])
				}
			} 
		}
	}
	fmt.Println("totalFinal", totalNum, prod)
	if totalNum == 2 {
		return true, prod
	}
	return false, 0
}

func getCompleteNum(j int, mat []string) int {
	p1, p2 := j, j;
	// fmt.Println("initial", p1, p2)
	len := len(mat)
	for {
		var flag1, flag2 bool
		if p1 == -1 {
			flag1 = false
		} else {
			flag1 = checkIfNum(mat[p1])
		}

		if p2 == len {
			flag2 = false
		} else {
			flag2 = checkIfNum(mat[p2])
		}
		// fmt.Println("for", p1, p2)
		if !flag1 && !flag2 {break}
		if ((p1 == -1 && (p2 == len || !flag2)) || (p2 == len && (p1 == 0 || !flag1))) {break}
		// fmt.Println(j, p1, p2)
		if flag1 {
			p1--;
		}
		if flag2 {
			p2++;
		}
	}


	num := 0
	for i:=p1+1; i<p2; i++ {
		curNum, err := strconv.Atoi(mat[i]);
		if err != nil {
			// fmt.Println("String to num conversion error", i, mat)
			// fmt.Println(p1, p2, mat[i])
			// fmt.Println(mat[p1], mat[p2], mat[j])
			// fmt.Println(mat[p1:p2])
			panic("wtf");
		}
		num = (num*10) + curNum
	}

	fmt.Println(num)
	return num;
}

// func checkAround(start int, end int, i int, mat [][]string, n int, m int) (bool) {
// 	start = int(math.Max(float64(start) - 1, 0))
// 	end = int(math.Min(float64(end) + 1, float64(m-1)))

// 	// top
// 	for k:=start; k<=end; k++ {
// 		if(i != 0){
// 			if !checkIfNum(mat[i-1][k]) {
// 				if(mat[i-1][k] == "*") {
// 					return true
// 				}
// 			}
// 		}
// 		if(i < n-1){
// 			if !checkIfNum(mat[i+1][k]) {
// 				if(mat[i+1][k] != ".") {
// 					return true
// 				}
// 			}
// 		}
// 	}

// 	// left
// 	if !checkIfNum(mat[i][start]) {
// 		if(mat[i][start] != ".") {
// 			return true
// 		}
// 	}

// 	// right
// 	if !checkIfNum(mat[i][end]) {
// 		if(mat[i][end] != ".") {
// 			return true
// 		}
// 	}

// 	return false;
// }

func checkIfNum(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil;
}


// curSum := 0;
// arr := []int{}
// arr = append(arr, curSum)
// i, err := strconv.Atoi(curText)
// slices.Max(arr)
// slices.Sort(arr)
// sort.Sort(sort.Reverse(sort.IntSlice(arr)))