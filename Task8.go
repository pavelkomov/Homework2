package main

/* 8. Напишите функцию-дженерик IsEqualArrays для comparable типов, которая сравнивает два неотсортированных массива.
Функция выдает булевое значение как результат. true - если массивы равны, false - если нет.
Массивы считаются равными, если в элемент из первого массива существует в другом, и наоборот.
Вне зависимости от расположения.
*/

func isEqualarrays[T comparable](arr1 []T, arr2 []T) bool {

	m1 := make(map[T]struct{})
	m2 := make(map[T]struct{})

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
