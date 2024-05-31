package main

import (
	"golang.org/x/exp/constraints"
)

/* 9. Реализуйте тип-дженерик Numbers, который является слайсом численных типов.
Реализуйте следующие методы для этого типа:
* суммирование всех элементов,
* произведение всех элементов,
* сравнение с другим слайсом на равность,
* проверка аргумента, является ли он элементом массива, если да - вывести индекс первого найденного элемента,
* удаление элемента массива по значению,
* удаление элемента массива по индексу.
*/

//	type Numbers interface {
//		NUm
//	}

type Number interface {
	constraints.Float | constraints.Integer //| constraints.Complex

}

type Vector[N Number] []N

func sum[N Number](num Vector[N]) N {
	var sum N

	for _, v := range num {
		sum += v
	}
	return sum
}

func multiple[N Number](num Vector[N]) N {
	var sum N

	for _, v := range num {
		sum *= v
	}
	return sum

}

func isEqual[N Number](arr1 Vector[N], arr2 Vector[N]) bool {

	m1 := make(map[N]struct{})
	m2 := make(map[N]struct{})

	for _, el := range arr1 {
		m1[el] = struct{}{}
	}

	for _, el := range arr2 {
		m2[el] = struct{}{}
	}

	for k, _ := range m1 {

		_, ok := m2[k]
		if !ok {
			return false
		}

	}
	return true

}
func isElement[N Number](el N, arr Vector[N]) (int, bool) {
	for i, v := range arr {
		if el == v {
			return i, true
		}
	}
	return 0, false
}

func deleteByValue[N Number](el N, arr Vector[N]) (Vector[N], bool) {

	for i, v := range arr {

		if el == v {
			arr = append(arr[:i], arr[i+1:]...)
			return arr, true
		}

	}

	return arr, false
}

func deleteByIndex[N Number](ind int, arr Vector[N]) (Vector[N], bool) {

	if ind > len(arr) {
		return arr, false
	}

	for i, _ := range arr {
		if i == ind {
			arr = append(arr[:i], arr[i+1:]...)
			return arr, true
		}
	}

	return arr, false

}
