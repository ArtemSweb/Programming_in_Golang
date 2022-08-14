package main

import "fmt"

func main(){
    var a int
    fmt.Scan(&a)
    // здесь ваш код
    res := a * 2
	res += 100
    fmt.Println(res)
}