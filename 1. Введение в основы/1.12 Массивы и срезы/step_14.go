// На ввод подаются пять целых чисел, которые записываются в массив. Однако эта часть программы уже написана.
// Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.

package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	ma := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > ma {
			ma = array[i]
		}
	}
	fmt.Println(ma)
}