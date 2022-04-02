package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	index   int
	parents []*node
	isRoot  bool
	value   int
}

type testCase struct {
	t     int
	n     int
	nodes []*node
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func reduceParents(node *node) (int, int) {
	parents := node.parents
	if len(parents) == 0 {
		return 0, 0
	}

	sum := parents[0].value
	minVal := parents[0].value
	for i := 1; i < len(parents); i++ {
		val := parents[i].value
		if val < minVal {
			minVal = val
		}
		sum += val
	}
	return minVal, sum - minVal
}

func solve(testCase testCase) {
	var sum int
	for i := testCase.n - 1; i >= 0; i-- {
		minVal, parentsSum := reduceParents(testCase.nodes[i])
		sum += parentsSum
		if testCase.nodes[i].isRoot {
			sum += maxInt(minVal, testCase.nodes[i].value)
		} else {
			testCase.nodes[i].value = maxInt(minVal, testCase.nodes[i].value)
		}
	}
	fmt.Printf("Case #%v: %v\n", testCase.t+1, sum)
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
		vals := strings.Split(line, " ")
		line, _ = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		children := strings.Split(line, " ")
		var nodes []*node
		for j := 0; j < n; j++ {
			val, _ := strconv.Atoi(vals[j])
			child, _ := strconv.Atoi(children[j])
			node := node{
				value:   val,
				parents: []*node{},
				index:   j,
				isRoot:  child == 0,
			}
			nodes = append(nodes, &node)
			if child != 0 {
				nodes[child-1].parents = append(nodes[child-1].parents, &node)
			}
		}
		testCases = append(testCases, testCase{t: i, n: n, nodes: nodes})
	}
	for _, testCase := range testCases {
		solve(testCase)
	}
}
