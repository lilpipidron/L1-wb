package tasks

import (
	"fmt"
	"sort"
)

// Всего имеет три способа с разной асимптотикой
// 1) O(n) по времени и O(n) по памяти, основывается на мапе, если элементы если в обоих множествах, то там будет стоять 3
// иначе 1 или 2
// 2) O(nlog(n)) по времени и O(1) по памяти, основывается на сортивке, сортируем и аккуратно идем указателями
// 3) O(n^2) по времени и O(1) по памяти, просто за квадрат

func Solve11_1() {
	arr1 := []int{5, 3, 2, 6, 1}
	arr2 := []int{3, 1, 5, 7, 2}
	var ans []int

	mapOfElems := make(map[int]int)

	for _, val := range arr1 {
		mapOfElems[val] = 1
	}
	for _, val := range arr2 {
		if mapOfElems[val] == 1 {
			mapOfElems[val] = 3
		} else {
			mapOfElems[val] = 2
		}
	}

	for key, val := range mapOfElems {
		if val == 3 {
			ans = append(ans, key)
		}
	}

	fmt.Println(ans)
}

func Solve11_2() {
	arr1 := []int{5, 3, 2, 6, 1}
	arr2 := []int{3, 1, 5, 7, 2}
	var ans []int
	var i, j int
	sort.Ints(arr1)
	sort.Ints(arr2)

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			ans = append(ans, arr1[i])
			i++
			j++
			continue
		}
		if arr1[i] > arr2[j] {
			j++
			continue
		}
		if arr1[i] < arr2[j] {
			i++
			continue
		}
	}

	fmt.Println(ans)
}

func Solve11_3() {
	arr1 := []int{5, 3, 2, 6, 1}
	arr2 := []int{3, 1, 5, 7, 2}
	var ans []int

	for _, val1 := range arr1 {
		for _, val2 := range arr2 {
			if val1 == val2 {
				ans = append(ans, val1)
			}
		}
	}

	fmt.Println(ans)
}
