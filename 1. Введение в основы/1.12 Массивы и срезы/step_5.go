package main

import "fmt"

func main() {
	var workArray [10]uint8

	//наполняем массив из ввода
	for i := range workArray {
		fmt.Scan(&workArray[i])
	}

	// по очереди получаем пары индексов для замены и меняем элементы массива
	var i1, i2 uint8
	for i := 0; i < 3; i++ {
		fmt.Scan(&i1, &i2)
		workArray[i1], workArray[i2] = workArray[i2], workArray[i1]
	}

	// выводим элементы массива через пробел
	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
}

