package tasks

import (
	"fmt"
	"sort"
)

//Исходный код функции sort.Search взят отсюда https://go.dev/src/sort/search.go
// Под копотом в этой функции и используется бинарный поиск

func Solve17() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	target := 5

	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if index < len(arr) && arr[index] == target {
		fmt.Println("Found:", target)
	} else {
		fmt.Println("Not found:", target)
	}
}
