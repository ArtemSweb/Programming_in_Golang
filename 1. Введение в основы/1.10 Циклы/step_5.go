// Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.

package main

import "fmt"

func main() {

	var n int
	max_n := 0
	count := 0

	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n > max_n {
			max_n = n
			count = 1
		} else if n == max_n {
			count += 1
		}
	}
	fmt.Println(count)
}