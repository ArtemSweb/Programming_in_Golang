// Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток.

package main

import "fmt"

func main() {
	var sec int

	fmt.Scan(&sec)
	h := (sec - (sec % 3600)) / 3600
	m := (sec % 3600) / 60

	fmt.Println("It is", h, "hours", m, "minutes.")
}