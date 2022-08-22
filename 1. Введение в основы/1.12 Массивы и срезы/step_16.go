// Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	lst := make([]int, n)

	for i := range lst {
		fmt.Scan(&lst[i])
	}

	cnt := 0
	for i := 0; i < len(lst); i++ {
		if lst[i] > 0 {
			cnt += 1
		}
	}

	fmt.Print(cnt)
}
