// Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем NN целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

package main
import "fmt"

func main() {
    // здесь должен быть ваш код
	var n int
	fmt.Scan(&n)

	slc := make([]int, n)

	for i := range slc {
		fmt.Scan(&slc[i])
	}

	fmt.Println(slc[3])
}