package tasks

import (
	"fmt"
	"sync"
)

// Все аналогично 2 заданию, но тут мы используем еще и mutex
// он нам нужен, чтобы несколько потоков одновременно писали и изменяли одну переменную
// у нас может возникнуть race condition

func Solve3() {
	arr := []int{2, 4, 6, 8, 10}
	var ans int
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	wg.Add(len(arr))
	for _, val := range arr {
		go func(val int) {
			defer wg.Done()
			tmp := val * val
			mx.Lock()
			ans += tmp
			mx.Unlock()
		}(val)
	}

	wg.Wait()

	fmt.Println(ans)
}
