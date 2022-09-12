/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.
*/

package main

import "fmt"

func main() {
	a, b := sumInt(1, 2, 3, 0, 4, 5)
	fmt.Println(a, b)
}

func sumInt(a ...int) (int, int) {
	total := 0

	for _, elem := range a {
		total += elem
	}
	return len(a), total
}