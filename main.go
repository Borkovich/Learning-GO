package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMax(1, 2, 3, 2342, 56, -54, 23, 0, 12))

}

func findMax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]

	for i := range numbers {
		if i > max {
			max = i
		}
	}

	return max
}
