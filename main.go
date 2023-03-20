package main

import "fmt"

func main() {
	fmt.Println(Max([]int{1, 2, -10, 5, 200, 10, 29, -15}))
}

func Max(numbers []int) int {
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}
