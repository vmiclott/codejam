package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type color struct {
	c int
	m int
	y int
	k int
}

type printer = color

type testCase struct {
	t        int
	printers []printer
}

const mil = 1000000

func solve(testCase testCase) {
	fmt.Printf("Case #%v: ", testCase.t+1)
	var minC, minM, minY, minK = mil, mil, mil, mil
	for _, printer := range testCase.printers {
		if printer.c < minC {
			minC = printer.c
		}
		if printer.m < minM {
			minM = printer.m
		}
		if printer.y < minY {
			minY = printer.y
		}
		if printer.k < minK {
			minK = printer.k
		}
	}
	if minC+minM+minY+minK < mil {
		fmt.Println("IMPOSSIBLE")
		return
	}
	var sum int
	if sum+minC >= mil {
		fmt.Printf("%v 0 0 0\n", minC)
		return
	}
	fmt.Printf("%v ", minC)
	sum += minC
	if sum+minM >= mil {
		fmt.Printf("%v 0 0\n", mil-sum)
		return
	}
	fmt.Printf("%v ", minM)
	sum += minM
	if sum+minY >= mil {
		fmt.Printf("%v 0\n", mil-sum)
		return
	}
	fmt.Printf("%v ", minY)
	sum += minY
	fmt.Printf("%v\n", mil-sum)
}

func lineToPrinter(line string) printer {
	splitLine := strings.Split(line, " ")
	c, _ := strconv.Atoi(splitLine[0])
	m, _ := strconv.Atoi(splitLine[1])
	y, _ := strconv.Atoi(splitLine[2])
	k, _ := strconv.Atoi(splitLine[3])
	return printer{
		c: c, m: m, y: y, k: k,
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
		printer1 := lineToPrinter(line)
		line, _ = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		printer2 := lineToPrinter(line)
		line, _ = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		printer3 := lineToPrinter(line)
		testCases = append(testCases, testCase{t: i, printers: []printer{printer1, printer2, printer3}})
	}
	for _, testCase := range testCases {
		solve(testCase)
	}
}
