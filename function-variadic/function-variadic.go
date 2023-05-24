package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	total := sumAll(12, 15, 15, 28, 90)
	fmt.Println(total)

	slice := []int{40, 20, 50}
	total = sumAll(slice...)
	fmt.Println(total)
}
