// myMath - пакет с единственной функцией сложения двух чисел

package myMath

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add принимает на вход два числа и возвращает результат их сложения, лол
func Add[T Number](a, b T) T {
	return a + b
}
