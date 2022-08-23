//Найдите количество минимальных элементов в последовательности.


package main

import . "fmt"

func main() {
	var n int
	Scan(&n)

	lst := make([]int, n)

	for i := range lst {
		Scan(&lst[i])
	}

	min_el := lst[0]
	cnt := 0

	for i := 0; i < len(lst); i++ {
		if lst[i] < min_el {
			cnt = 1
			min_el = lst[i]
		} else if min_el == lst[i] {
			cnt += 1
		}
	}

	Print(cnt)
}