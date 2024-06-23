package tasks

import "fmt"

// 1) Решение основано на создании двух срезов, которые не включают элемент и последующего объедениния их
// 2) На место удаляемого элемента ставим самый последний и обрезаем слайс, из проблем имеем, что порядок не сохраняем
// 3) Просто в срез с нужного эелемента копируем срез на эелемент вперед

func Solve23_1() {
	arr := []int{1, 2, 3, 4, 5, 6}
	var n int

	fmt.Scan(&n)

	fmt.Println(append(arr[:n], arr[n+1:]...))
}

func Solve23_2() {
	arr := []int{1, 2, 3, 4, 5, 6}
	var n int

	fmt.Scan(&n)

	arr[n] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	fmt.Println(arr)
}

func Solve23_3() {
	arr := []int{1, 2, 3, 4, 5, 6}
	var n int

	fmt.Scan(&n)

	copy(arr[n:], arr[n+1:])

	arr = arr[:len(arr)-1]

	fmt.Println(arr)
}
