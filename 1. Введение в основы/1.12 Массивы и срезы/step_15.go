// Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).


package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	lst := make([]int, n)

	for i := range lst {
		fmt.Scan(&lst[i])
	}
	for i := 0; i < len(lst); i++ {
		if i%2 == 0 {
			fmt.Print(lst[i], " ")
		}
	}
}