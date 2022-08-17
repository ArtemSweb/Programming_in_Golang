package main

import "fmt"

func main() {
	// put your code here
	var n float64

	fmt.Scan(&n)

	n1 := n * n

	if n <= 0 {
		fmt.Printf("число %2.2f не подходит", n)
	} else if n1 > 10000 {
		fmt.Printf("%e", n)
	} else {
		fmt.Printf("%.4f", n1)
	}
}