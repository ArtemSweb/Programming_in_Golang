// Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

func minimumFromFour() int {
	var min, n int

	for i := 0; i < 4; i++ {
		fmt.Scan(&n)
		if i == 0 || n < min {
			min = n
		}
	}
	return min
}