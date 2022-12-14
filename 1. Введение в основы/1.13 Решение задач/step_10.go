/*
Цифровой корень
Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр, на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации. Этот процесс повторяется до тех пор, пока не будет получена одна цифра.

Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 . 

По данному числу определите его цифровой корень.
*/

package main

import . "fmt"

func main() {
	// через циклы
	/*  var n int
    Scan(&n)

    for n >= 10 {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		n = sum
    }
    Print(n)
  */
  //через формулу
	var n int
	Scan(&n)

	dt := 1 + ((n - 1) % 9)

	print(dt)
}




