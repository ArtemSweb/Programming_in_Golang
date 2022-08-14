// Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.

package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	fmt.Scan(&a, &b)
	sum := 0
	for ; a <= b; a++ {
		sum += a
	}
	fmt.Println(sum)
}