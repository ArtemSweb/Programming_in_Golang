// Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника. Нужно проверить, является ли треугольник прямоугольным. Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"


package main

import "fmt"

func main() {
	lst := make([]int, 3)

	for i := range lst {
		fmt.Scan(&lst[i])
	}

	a := lst[0]
	b := lst[1]
	c := lst[2]

	if c*c == b*b+a*a {
		fmt.Print("Прямоугольный")
	} else {
		fmt.Print("Непрямоугольный")
	}
}




