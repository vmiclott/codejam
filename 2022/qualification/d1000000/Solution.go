package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type testCase struct {
	t    int
	n    int
	dice []int
}

func solve(testCase testCase) {
	sort.Ints(testCase.dice)
	var count int
	for i := 0; i < testCase.n; i++ {
		if count < testCase.dice[i] {
			count++
		}
	}

	fmt.Printf("Case #%v: %v\n", testCase.t+1, count)
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
		n, _ := strconv.Atoi(line)
		line, _ = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		nums := strings.Split(line, " ")
		var dice []int
		for j := 0; j < n; j++ {
			die, _ := strconv.Atoi(nums[j])
			dice = append(dice, die)
		}
		testCases = append(testCases, testCase{t: i, n: n, dice: dice})
	}
	for _, testCase := range testCases {
		solve(testCase)
	}
}
