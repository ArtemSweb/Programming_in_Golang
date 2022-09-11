// Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.


func vote(x int, y int, z int) int {
	//print your code here
	if x+y+z >= 2 {
		return 1
	} else {
		return 0
	}
}