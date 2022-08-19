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