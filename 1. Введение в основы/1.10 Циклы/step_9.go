// Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.

package main

import "fmt"

func main() {
	var (
		x int
		p int
		y int
	)
	step := 0

	fmt.Scan(&x, &p, &y)

	for ; x < y; step++ {
		x += x * p / 100
	}

	fmt.Println(step)
}