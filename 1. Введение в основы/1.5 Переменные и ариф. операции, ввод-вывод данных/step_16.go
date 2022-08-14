package main

import "fmt"

func main() {
	var r int
	fmt.Scan(&r)
	d := r * 2
	fmt.Println("It is", d/60, "hours", d%60, "minutes.")
}
