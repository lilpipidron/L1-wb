package tasks

import (
	"fmt"
	"sync"
)

// Суть решения: в массив ответ записываем квадрат элемента, который находится на данной позиции в изначальном массиве
// И потом просто циклом обычным проходим по полученному массиву и выводим его
// решение основно на WaitGroup
// мы используем wg для того чтобы дождаться когда произойдет запись во все позиции массива с ответом
// в данном случае мы не используем mutex потому что у нас гарантировано запись идет в разные участки массива

func Solve2() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(len(arr))
	ans := make([]int, len(arr))
	for i, val := range arr {
		go func(value, pos int) {
			defer wg.Done()
			ans[pos] = value * value
		}(val, i)
	}

	wg.Wait()
	for _, val := range ans {
		fmt.Println(val)
	}
}
