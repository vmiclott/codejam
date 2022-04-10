package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type testCase struct {
	t int
	r int
	c int
}

func draw(testCase testCase) {
	fmt.Printf("Case #%v:\n", testCase.t+1)
	w, h := 2*testCase.c+1, 2*testCase.r+1
	// first line
	fmt.Print("..+")
	for i := 3; i < w; i += 2 {
		fmt.Print("-+")
	}
	fmt.Println()
	// second line
	fmt.Print("..|")
	for i := 3; i < w; i += 2 {
		fmt.Print(".|")
	}
	fmt.Println()
	for i := 2; i < h; i++ {
		if i%2 == 0 {
			fmt.Print("+")
			for j := 2; j < w; j += 2 {
				fmt.Print("-+")
			}
			fmt.Println()
		} else {
			fmt.Print("|")
			for j := 2; j < w; j += 2 {
				fmt.Print(".|")
			}
			fmt.Println()
		}
	}
}

func main() {
	var reader = bufio.NewReader(os.Stdin)

	var testCases []testCase
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	t, _ := strconv.Atoi(line)
	for i := 0; i < t; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		nums := strings.Split(line, " ")
		r, _ := strconv.Atoi(nums[0])
		c, _ := strconv.Atoi(nums[1])
		testCases = append(testCases, testCase{t: i, r: r, c: c})
	}
	for _, testCase := range testCases {
		draw(testCase)
	}
}
