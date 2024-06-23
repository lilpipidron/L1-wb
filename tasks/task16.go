package tasks

import (
	"fmt"
	"math/rand"
)

// Самая обычная реализация quicksort с случайным пивотом

func Solve16() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quicksort(arr []int, l, r int) {
	if l < r {
		partitionIndex := partition(arr, l, r)
		quicksort(arr, l, partitionIndex)
		quicksort(arr, partitionIndex+1, r)
	}
}
func partition(arr []int, l, r int) int {
	pivot := arr[rand.Int()%len(arr)]
	for l <= r {
		for arr[l] < pivot {
			l++
		}

		for arr[r] > pivot {
			r--
		}

		if l >= r {
			break
		}

		arr[l], arr[r] = arr[r], arr[l]
	}
	return r
}
