package main

import "fmt"

func fibonacci() func(int) int {
	return func(x int) int {
		if x == 0 {
			return 0
		}
		if x == 1 {
			return 1
		} else {
			s := make([]int, 0, x)
			s = append(s, 0, 1)
			for i := 2; i < x; i++ {
				s = append(s, s[i-1]+s[i-2])
			}
			return s[x-1]
		}
	}
}

func main() {
	pos := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(pos(i))
	}
}
