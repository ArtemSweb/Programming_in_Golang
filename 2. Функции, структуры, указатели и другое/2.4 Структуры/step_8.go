/*
В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции, как объекты из реальной жизни. Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.

Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно. У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.
*/

package main


type Dread struct {
	On    bool
	Ammo  int
	Power int
}

func (r *Dread) Shoot() bool {
	if r.On == false || r.Ammo == 0 {
		return false
	}
	r.Ammo--
	return true
}

func (r *Dread) RideBike() bool {
	if r.On == false || r.Power == 0 {
		return false
	}
	r.Power--
	return true
}

func main() {
	testStruct := new(Dread)

}