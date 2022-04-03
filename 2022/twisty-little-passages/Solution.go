package main

import (
	"fmt"
	"math/rand"
)

type room struct {
	r int
	p int
}

func less(u, v room) bool {
	return u.p < v.p || (u.p == v.p && u.r < v.r)
}

func main() {
	var T int
	fmt.Scanln(&T)
	for t := 0; t < T; t++ {
		var N, K int
		fmt.Scanln(&N, &K)
		if K >= N {
			var R, P int
			fmt.Scanln(&R, &P)
			firstRoom := R
			sum := P
			for i := 0; i < N; i++ {
				if i+1 == firstRoom {
					continue
				}
				fmt.Printf("T %v\n", i+1)
				fmt.Scanln(&R, &P)
				sum += P
			}
			fmt.Printf("E %v\n", sum/2)
		} else {
			var R, P int
			var count, sum int
			var prevRoom, currRoom room
			degrees := make(map[int]int)
			fmt.Scanln(&R, &P)
			prevRoom = room{r: R, p: P}
			degrees[R] = P
			fmt.Print("W\n")
			fmt.Scanln(&R, &P)
			currRoom = room{r: R, p: P}
			degrees[R] = P
			if less(prevRoom, currRoom) {
				sum += prevRoom.p
				count += 1
			}
			for i := 0; i < K-2; i += 2 {
				r := rand.Intn(N) + 1
				fmt.Printf("T %v\n", r)
				fmt.Scanln(&R, &P)
				sum += prevRoom.p
				count += 1
				prevRoom = room{r: R, p: P}
				degrees[R] = P
				fmt.Print("W\n")
				fmt.Scanln(&R, &P)
				currRoom = room{r: R, p: P}
				degrees[R] = P
			}
			eDegree := float64(sum) / float64(count)
			var res float64
			for i := 0; i < N; i++ {
				if d, ok := degrees[i+1]; !ok {
					res += eDegree
				} else {
					res += float64(d)
				}
			}
			e := int(res/float64(2) + 0.5)
			fmt.Printf("E %v\n", e)
		}
	}
}
