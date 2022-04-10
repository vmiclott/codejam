package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

var max int = pow(10, 9)

func main() {
	var T int
	fmt.Scanln(&T)
	for t := 0; t < T; t++ {
		var N int
		fmt.Scanln(&N)

		numbers := make(map[int]bool, 2*N)
		i := 0
		for i < N && pow(2, i) <= max {
			numbers[pow(2, i)] = true
			i++
		}
		x := 1
		for i < N {
			for numbers[x] {
				x++
			}
			numbers[x] = true
			i++
		}
		numString := ""
		first := true
		for x := range numbers {
			if first {
				numString += fmt.Sprintf("%v", x)
				first = false
			} else {
				numString += fmt.Sprintf(" %v", x)
			}
		}
		fmt.Println(numString)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		numStrings := strings.Split(scanner.Text(), " ")
		for _, s := range numStrings {
			x, _ = strconv.Atoi(s)
			numbers[x] = true
		}

		var sum int
		var slice []int
		for x := range numbers {
			sum += x
			slice = append(slice, x)
		}

		sort.Ints(slice)
		half := sum / 2
		res := fmt.Sprintf("%v", slice[len(slice)-1])
		sum = slice[len(slice)-1]
		for i := len(slice) - 2; i >= 0; i-- {
			if sum+slice[i] <= half {
				sum += slice[i]
				res += fmt.Sprintf(" %v", slice[i])
			}
			if sum == half {
				break
			}
		}
		fmt.Println(res)
	}
}
