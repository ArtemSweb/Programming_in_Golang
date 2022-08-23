//Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.


package main

import . "fmt"

func main() {

	var a, b, c int
	Scan(&a, &b, &c)

	if a+b > c && a+c > b && b+c > a {
		Print("Существует")
	} else {
		Print("Не существует")
	}
}
